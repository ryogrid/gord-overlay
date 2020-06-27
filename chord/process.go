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

type Process struct {
	*LocalNode
	SuccessorStabilizer   Stabilizer
	FingerTableStabilizer Stabilizer
	Transport             Transport

	opt        *processOption
	isShutdown bool
}

type processOption struct {
	successorStabilizerInterval   time.Duration
	fingerTableStabilizerInterval time.Duration
	timeoutConnNode               time.Duration
	existNode                     RingNode
}

type ProcessOptionFunc func(option *processOption)

func newDefaultProcessOption() *processOption {
	return &processOption{
		successorStabilizerInterval:   1 * time.Second,
		fingerTableStabilizerInterval: 100 * time.Millisecond,
		timeoutConnNode:               1 * time.Second,
	}
}

func WithSuccessorStabilizeInterval(duration time.Duration) ProcessOptionFunc {
	return func(option *processOption) {
		option.successorStabilizerInterval = duration
	}
}

func WithFingerTableStabilizeInterval(duration time.Duration) ProcessOptionFunc {
	return func(option *processOption) {
		option.fingerTableStabilizerInterval = duration
	}
}

func WithExistNode(node RingNode) ProcessOptionFunc {
	return func(option *processOption) {
		option.existNode = node
	}
}

func NewProcess(localNode *LocalNode, transport Transport) *Process {
	process := &Process{
		LocalNode: localNode,
		Transport: transport,
	}
	process.SuccessorStabilizer = NewSuccessorStabilizer(localNode)
	process.FingerTableStabilizer = NewFingerTableStabilizer(localNode)
	return process
}

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
	p.scheduleStabilizer(ctx, p.opt.successorStabilizerInterval, p.SuccessorStabilizer)
	p.scheduleStabilizer(ctx, p.opt.fingerTableStabilizerInterval, p.FingerTableStabilizer)
	return nil
}

func (p *Process) activate(ctx context.Context, existNode RingNode) error {
	if existNode == nil {
		p.LocalNode.CreateRing()
		return nil
	}
	if err := p.LocalNode.JoinRing(ctx, existNode); err != nil {
		return err
	}
	return nil
}

func (p *Process) Shutdown() {
	p.isShutdown = true
	p.Transport.Shutdown()
}

func (p *Process) scheduleStabilizer(ctx context.Context, interval time.Duration, stabilizer Stabilizer) {
	if p.isShutdown {
		return
	}
	go func() {
		stabilizer.Stabilize(ctx)
		time.AfterFunc(interval, func() {
			p.scheduleStabilizer(ctx, interval, stabilizer)
		})
	}()
}
