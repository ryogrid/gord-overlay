package chord

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Process represents chord process.
// Process manages a local node and some stabilizers.
type Process struct {
	*LocalNode
	AliveStabilizer       Stabilizer
	SuccessorStabilizer   Stabilizer
	FingerTableStabilizer Stabilizer
	Transport             Transport
	IsShutdown            bool

	opt *processOption
}

type processOption struct {
	stabilizerInterval time.Duration
	timeoutConnNode    time.Duration
	existNode          RingNode
}

// ProcessOptionFunc is function to apply options to a process
type ProcessOptionFunc func(option *processOption)

func newDefaultProcessOption() *processOption {
	return &processOption{
		stabilizerInterval: 200 * time.Millisecond, //50 * time.Millisecond,
		timeoutConnNode:    1 * time.Second,
	}
}

func WithStabilizeInterval(duration time.Duration) ProcessOptionFunc {
	return func(option *processOption) {
		option.stabilizerInterval = duration
	}
}

func WithExistNode(node RingNode) ProcessOptionFunc {
	return func(option *processOption) {
		option.existNode = node
	}
}

// NewProcess creates a process.
func NewProcess(localNode *LocalNode, transport Transport) *Process {
	process := &Process{
		LocalNode: localNode,
		Transport: transport,
	}
	process.AliveStabilizer = NewAliveStabilizer(localNode)
	process.SuccessorStabilizer = NewSuccessorStabilizer(localNode)
	process.FingerTableStabilizer = NewFingerTableStabilizer(localNode)
	return process
}

// Start starts a process.
// Creates or joins in chord ring and starts some stabilizers of a process.
func (p *Process) Start(ctx context.Context, opts ...ProcessOptionFunc) error {
	p.opt = newDefaultProcessOption()
	for _, opt := range opts {
		opt(p.opt)
	}
	if p.opt.existNode != nil && p.opt.existNode.Reference().Host == p.Host {
		log.Fatalf("exist node must be different from local node.")
	}
	if err := p.activate(ctx, p.opt.existNode); err != nil {
		return err
	}
	p.scheduleStabilizers(ctx, p.opt.stabilizerInterval, p.SuccessorStabilizer, p.FingerTableStabilizer, p.AliveStabilizer)
	return nil
}

func (p *Process) activate(ctx context.Context, existNode RingNode) error {
	if existNode == nil {
		p.LocalNode.CreateRing()
		return nil
	}
	p.LocalNode.predecessor = existNode
	if err := p.LocalNode.JoinRing(ctx, existNode); err != nil {
		return err
	}
	return nil
}

// Shutdown stops process
func (p *Process) Shutdown() {
	p.IsShutdown = true
	p.LocalNode.Shutdown()
	p.Transport.Shutdown()
}

func (p *Process) scheduleStabilizers(ctx context.Context, interval time.Duration, stabilizers ...Stabilizer) {
	if p.IsShutdown {
		return
	}
	go func() {
		for _, s := range stabilizers {
			s.Stabilize(ctx)
		}
		time.AfterFunc(interval, func() {
			p.scheduleStabilizers(ctx, interval, stabilizers...)
		})
	}()
}
