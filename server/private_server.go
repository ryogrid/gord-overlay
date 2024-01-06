package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ryogird/gord-overlay/serverconnect"
	"github.com/ryogrid/gord-overlay/chord"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

// InternalServer represents gRPC server to expose for internal chord processes
type InternalServer struct {
	port       string
	process    *chord.Process
	opt        *chordOption
	shutdownCh chan struct{}
}

type chordOption struct {
	host            string
	timeoutConnNode time.Duration
	processOpts     []chord.ProcessOptionFunc
}

// InternalServerOptionFunc represents server options for internal
type InternalServerOptionFunc func(option *chordOption)

func newDefaultServerOption() *chordOption {
	return &chordOption{
		host:            "127.0.0.1",
		timeoutConnNode: time.Second * 5,
	}
}

func WithNodeOption(host string) InternalServerOptionFunc {
	return func(option *chordOption) {
		option.host = host
	}
}

func WithProcessOptions(opts ...chord.ProcessOptionFunc) InternalServerOptionFunc {
	return func(option *chordOption) {
		option.processOpts = append(option.processOpts, opts...)
	}
}

func WithTimeoutConnNode(duration time.Duration) InternalServerOptionFunc {
	return func(option *chordOption) {
		option.timeoutConnNode = duration
	}
}

// NewChordServer creates a chord server
func NewChordServer(process *chord.Process, port string, opts ...InternalServerOptionFunc) *InternalServer {
	opt := newDefaultServerOption()
	for _, o := range opts {
		o(opt)
	}
	return &InternalServer{
		process:    process,
		port:       port,
		opt:        opt,
		shutdownCh: make(chan struct{}, 1),
	}
}

/*
func (is *InternalServer) newGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	RegisterInternalServiceServer(s, is)
	return s
}
*/

// Run runs chord server.
func (is *InternalServer) Run(ctx context.Context) {
	go func() {
		//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", is.opt.host, is.port))
		//if err != nil {
		//	log.Fatalf("failed to run chord server. reason: %#v", err)
		//}
		//grpcServer := is.newGrpcServer()
		//if err := grpcServer.Serve(lis); err != nil {
		//	log.Fatalf("failed to run chord server. reason: %#v", err)
		//}

		mux := http.NewServeMux()
		path, handler := serverconnect.NewExternalServiceHandler(is)
		mux.Handle(path, handler)
		http.ListenAndServe(
			"127.0.0.1:26041",
			mux,
			//// Use h2c so we can serve HTTP/2 without TLS.
			//h2c.NewHandler(mux, &http2.Server{}),
		)
	}()
	if err := is.process.Start(ctx, is.opt.processOpts...); err != nil {
		log.Fatalf("failed to run chord server. reason: %#v", err)
	}
	log.Info("Running Chord server...")
	log.Infof("Chord listening on %s:%s", is.process.Host, is.port)
	<-is.shutdownCh
	is.process.Shutdown()
}

func (is *InternalServer) Shutdown() {
	is.shutdownCh <- struct{}{}
}

func (is *InternalServer) Ping(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	return &empty.Empty{}, nil
}

func (is *InternalServer) Successors(ctx context.Context, req *empty.Empty) (*Nodes, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successors, err := is.process.GetSuccessors(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. successor is not set.")
	}
	var nodes []*Node
	for _, suc := range successors {
		if suc == nil {
			continue
		}
		nodes = append(nodes, &Node{
			Host: suc.Reference().Host,
		})
	}
	return &Nodes{
		Nodes: nodes,
	}, nil
}

func (is *InternalServer) Predecessor(ctx context.Context, _ *empty.Empty) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	pred, err := is.process.GetPredecessor(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. predecessor is not set.")
	}
	if pred != nil {
		return &Node{
			Host: pred.Reference().Host,
		}, nil
	}
	return nil, status.Errorf(codes.NotFound, "server: predecessor is not set.")
}

func (is *InternalServer) FindSuccessorByTable(ctx context.Context, req *FindRequest) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByTable(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor failed. reason = %#v", err)
	}
	return &Node{
		Host: successor.Reference().Host,
	}, nil
}

func (is *InternalServer) FindSuccessorByList(ctx context.Context, req *FindRequest) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByList(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor fallback failed. reason = %#v", err)
	}
	return &Node{
		Host: successor.Reference().Host,
	}, nil
}

func (is *InternalServer) FindClosestPrecedingNode(ctx context.Context, req *FindRequest) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	node, err := is.process.FindClosestPrecedingNode(ctx, req.Id)
	if err == chord.ErrStabilizeNotCompleted {
		return nil, status.Error(codes.NotFound, "Stabilize not completed.")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find closest preceding node failed. reason = %#v", err)
	}
	return &Node{
		Host: node.Reference().Host,
	}, nil
}

func (is *InternalServer) Notify(ctx context.Context, req *Node) (*empty.Empty, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	err := is.process.Notify(ctx, chord.NewRemoteNode(req.Host, is.process.Transport))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: notify failed. reason = %#v", err)
	}
	return &empty.Empty{}, nil
}

func (is *InternalServer) PutValueInner(ctx context.Context, req *PutValueInnerRequest) (*PutValueInnerResponse, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	success, err := is.process.PutValue(ctx, &req.Key, &req.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: put value failed. reason = %#v", err)
	}
	return &PutValueInnerResponse{
		Success: success,
	}, nil

}

func (is *InternalServer) GetValueInner(ctx context.Context, req *GetValueInnerRequest) (*GetValueInnerResponse, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	val, success, err := is.process.GetValue(ctx, &req.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: get value failed. reason = %#v", err)
	}

	return &GetValueInnerResponse{
		Value:   *val,
		Success: success,
	}, nil
}

func (is *InternalServer) DeleteValueInner(ctx context.Context, req *DeleteValueInnerRequest) (*DeleteValueInnerResponse, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	success, err := is.process.DeleteValue(ctx, &req.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: delete value failed. reason = %#v", err)
	}
	return &DeleteValueInnerResponse{
		Success: success,
	}, nil
}
