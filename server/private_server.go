package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ryogrid/gord-overlay/api_internal"
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

func (is *InternalServer) newRpcServer() *api_internal.Server {
	//s := grpc.NewServer()
	//reflection.Register(s)
	//RegisterInternalServiceServer(s, is)

	s, err := api_internal.NewServer(is)
	if err != nil {
		panic(err)
	}
	return s
}

// Run runs chord server.
func (is *InternalServer) Run(ctx context.Context) {
	go func() {
		//lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
		//if err != nil {
		//	log.Fatalf("failed to run server. reason: %#v", err)
		//}
		rpcServer := is.newRpcServer()
		err := http.ListenAndServe("0.0.0.0:"+is.port, rpcServer)
		if err != nil {
			panic(err)
		}
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

func (is *InternalServer) InternalServicePing(ctx context.Context) (*api_internal.ServerSuccessResponse, error) {
	//func (is *InternalServer) Ping(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if is.process.IsShutdown {
		return &api_internal.ServerSuccessResponse{Success: api_internal.NewOptBool(false)}, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	return &api_internal.ServerSuccessResponse{Success: api_internal.NewOptBool(true)}, nil
}

func (is *InternalServer) InternalServiceSuccessors(ctx context.Context) (*api_internal.ServerNodes, error) {
	//func (is *InternalServer) Successors(ctx context.Context, req *empty.Empty) (*Nodes, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successors, err := is.process.GetSuccessors(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. successor is not set.")
	}
	var nodes []api_internal.ServerNode
	for _, suc := range successors {
		if suc == nil {
			continue
		}
		nodes = append(nodes, api_internal.ServerNode{
			Host: api_internal.NewOptString(suc.Reference().Host),
		})
	}
	return &api_internal.ServerNodes{
		Nodes: nodes,
	}, nil
}

func (is *InternalServer) InternalServicePredecessor(ctx context.Context) (*api_internal.ServerNode, error) {
	//func (is *InternalServer) Predecessor(ctx context.Context, _ *empty.Empty) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	pred, err := is.process.GetPredecessor(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. predecessor is not set.")
	}
	if pred != nil {
		return &api_internal.ServerNode{
			Host: api_internal.NewOptString(pred.Reference().Host),
		}, nil
	}
	return nil, status.Errorf(codes.NotFound, "server: predecessor is not set.")
}

func (is *InternalServer) InternalServiceFindSuccessorByTable(ctx context.Context, params api_internal.InternalServiceFindSuccessorByTableReq) (*api_internal.ServerNode, error) {
	//func (is *InternalServer) FindSuccessorByTable(ctx context.Context, req *api_internal.FindRequest) (*api_internal.ServerNode, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByTable(ctx, params.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor failed. reason = %#v", err)
	}
	return &api_internal.ServerNode{
		Host: api_internal.NewOptString(successor.Reference().Host),
	}, nil
}

func (is *InternalServer) InternalServiceFindSuccessorByList(ctx context.Context, params api_internal.InternalServiceFindSuccessorByListReq) (*api_internal.ServerNode, error) {
	//func (is *InternalServer) FindSuccessorByList(ctx context.Context, req *FindRequest) (*Node, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByList(ctx, params.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor fallback failed. reason = %#v", err)
	}
	return &api_internal.ServerNode{
		Host: api_internal.NewOptString(successor.Reference().Host),
	}, nil
}

func (is *InternalServer) InternalServiceFindClosestPrecedingNode(ctx context.Context, params api_internal.InternalServiceFindClosestPrecedingNodeReq) (*api_internal.ServerNode, error) {
	//func (is *InternalServer) FindClosestPrecedingNode(ctx context.Context, req *api_internal.FindRequest) (*api_internal.ServerNode, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	node, err := is.process.FindClosestPrecedingNode(ctx, params.ID)
	if err == chord.ErrStabilizeNotCompleted {
		return nil, status.Error(codes.NotFound, "Stabilize not completed.")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find closest preceding node failed. reason = %#v", err)
	}
	return &api_internal.ServerNode{
		Host: api_internal.NewOptString(node.Reference().Host),
	}, nil
}

func (is *InternalServer) InternalServiceNotify(ctx context.Context, params api_internal.InternalServiceNotifyReq) (*api_internal.ServerSuccessResponse, error) {
	//func (is *InternalServer) Notify(ctx context.Context, req *Node) (*empty.Empty, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	err := is.process.Notify(ctx, chord.NewRemoteNode(params.Host.Value, is.process.Transport))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: notify failed. reason = %#v", err)
	}
	return &api_internal.ServerSuccessResponse{Success: api_internal.NewOptBool(true)}, nil
}

func (is *InternalServer) InternalServicePutValueInner(ctx context.Context, params api_internal.InternalServicePutValueInnerReq) (*api_internal.ServerPutValueInnerResponse, error) {
	//func (is *InternalServer) PutValueInner(ctx context.Context, req *PutValueInnerRequest) (*PutValueInnerResponse, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	success, err := is.process.PutValue(ctx, &params.Key.Value, &params.Value.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: put value failed. reason = %#v", err)
	}
	return &api_internal.ServerPutValueInnerResponse{
		Success: api_internal.NewOptBool(success),
	}, nil

}

func (is *InternalServer) InternalServiceGetValueInner(ctx context.Context, params api_internal.InternalServiceGetValueInnerReq) (*api_internal.ServerGetValueInnerResponse, error) {
	//func (is *InternalServer) GetValueInner(ctx context.Context, req *GetValueInnerRequest) (*GetValueInnerResponse, error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	val, success, err := is.process.GetValue(ctx, &params.Key.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: get value failed. reason = %#v", err)
	}

	return &api_internal.ServerGetValueInnerResponse{
		Value:   api_internal.NewOptString(*val),
		Success: api_internal.NewOptBool(success),
	}, nil
}

func (is *InternalServer) InternalServiceDeleteValueInner(ctx context.Context, params api_internal.InternalServiceDeleteValueInnerReq) error {
	//func (is *InternalServer) DeleteValueInner(ctx context.Context, req *DeleteValueInnerRequest) (*DeleteValueInnerResponse, error) {
	if is.process.IsShutdown {
		//return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
		return errors.New("server has started shutdown")
	}
	_, err := is.process.DeleteValue(ctx, &params.Key.Value)
	if err != nil {
		//return nil, status.Errorf(codes.Internal, "server: delete value failed. reason = %#v", err)
		panic(err)
	}
	//return &api_internal.ServerDeleteValueInnerResponse{
	//	Success: success,
	//}, nil
	return nil
}

func (is *InternalServer) NewError(ctx context.Context, err error) *api_internal.ErrorStatusCode {
	return &api_internal.ErrorStatusCode{Response: api_internal.Error{Message: fmt.Sprintf("%v", err), Code: -1}}
}
