package core

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/server"
	"github.com/ryogrid/gord-overlay/serverconnect"
	"github.com/ryogrid/gossip-overlay/overlay"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

// InternalServer represents gRPC server to expose for internal chord processes
type InternalServer struct {
	olPeer     *overlay.OverlayPeer
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
		host:            "0.0.0.0",
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
func NewChordServer(process *chord.Process, olPeer *overlay.OverlayPeer, port string, opts ...InternalServerOptionFunc) *InternalServer {
	opt := newDefaultServerOption()
	for _, o := range opts {
		o(opt)
	}
	return &InternalServer{
		olPeer:     olPeer,
		port:       port,
		process:    process,
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
		path, handler := serverconnect.NewInternalServiceHandler(is)
		mux.Handle(path, handler)
		//http.ListenAndServe(
		//	"0.0.0.0"+":"+is.port,
		//	//mux,
		//	//// Use h2c so we can serve HTTP/2 without TLS.
		//	h2c.NewHandler(mux, &http2.Server{}),
		//)

		//serv := &http2.Server{}
		////http.Serve(overlay.NewOverlayListener("0.0.0.0"+":"+is.port), mux)
		http.Serve(is.olPeer.GetOverlayListener(), mux)
		////serv.MaxReadFrameSize = 1 << 31
		//http.Serve(is.olPeer.GetOverlayListener(), h2c.NewHandler(mux, serv))
		//oserv, err := overlay.NewOverlayServer(is.olPeer.Peer, is.olPeer.Peer.GossipMM)
		//if err != nil {
		//	panic(err)
		//}
		//
		//serv := &http2.Server{}
		//for {
		//	channel, _, _, err := oserv.Accept()
		//	if err != nil {
		//		fmt.Println("InternalServer::Run", fmt.Sprintf("%v", err))
		//		continue
		//	}
		//	fmt.Println("InternalServer::Run", fmt.Sprintf("%v", channel))
		//
		//	tlsChannel := tls.Server(channel, &tls.Config{MinVersion: tls.VersionTLS10, InsecureSkipVerify: true})
		//	serv.ServeConn(tlsChannel, &http2.ServeConnOpts{
		//		Handler: mux,
		//	})
		//}
	}()
	if err := is.process.Start(ctx, is.opt.processOpts...); err != nil {
		log.Fatalf("failed to run chord server. reason: %v", err)
	}
	log.Info("Running Chord server...")
	log.Infof("Chord listening on %s:%s", is.process.Host, is.port)

	// TODO: for debugging
	log.Infof("Chord listening on %d", is.olPeer.Peer.GossipDataMan.Self)
	<-is.shutdownCh
	is.process.Shutdown()
}

func (is *InternalServer) Shutdown() {
	is.shutdownCh <- struct{}{}
}

func (is *InternalServer) Ping(_ context.Context, _ *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	return &connect.Response[emptypb.Empty]{
		Msg: &emptypb.Empty{},
	}, nil
}

func (is *InternalServer) Successors(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[server.Nodes], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successors, err := is.process.GetSuccessors(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. successor is not set.")
	}
	var nodes []*server.Node
	for _, suc := range successors {
		if suc == nil {
			continue
		}
		nodes = append(nodes, &server.Node{
			Host: suc.Reference().Host,
		})
	}
	return &connect.Response[server.Nodes]{
		Msg: &server.Nodes{
			Nodes: nodes,
		},
	}, nil
}

func (is *InternalServer) Predecessor(ctx context.Context, _ *connect.Request[emptypb.Empty]) (*connect.Response[server.Node], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	pred, err := is.process.GetPredecessor(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: internal error occured. predecessor is not set.")
	}
	if pred != nil {
		return &connect.Response[server.Node]{
			Msg: &server.Node{
				Host: pred.Reference().Host,
			},
		}, nil
	}
	return nil, status.Errorf(codes.NotFound, "server: predecessor is not set.")
}

func (is *InternalServer) FindSuccessorByTable(ctx context.Context, req *connect.Request[server.FindRequest]) (*connect.Response[server.Node], error) {
	fmt.Println("InternalServer::FindSuccessorByTable", req.Msg.Id)
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByTable(ctx, req.Msg.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor failed. reason = %#v", err)
	}
	return &connect.Response[server.Node]{
		Msg: &server.Node{
			Host: successor.Reference().Host,
		},
	}, nil
}

func (is *InternalServer) FindSuccessorByList(ctx context.Context, req *connect.Request[server.FindRequest]) (*connect.Response[server.Node], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	successor, err := is.process.FindSuccessorByList(ctx, req.Msg.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find successor fallback failed. reason = %#v", err)
	}
	return &connect.Response[server.Node]{
		Msg: &server.Node{
			Host: successor.Reference().Host,
		},
	}, nil
}

func (is *InternalServer) FindClosestPrecedingNode(ctx context.Context, req *connect.Request[server.FindRequest]) (*connect.Response[server.Node], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	node, err := is.process.FindClosestPrecedingNode(ctx, req.Msg.Id)
	if err == chord.ErrStabilizeNotCompleted {
		return nil, status.Error(codes.NotFound, "Stabilize not completed.")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: find closest preceding node failed. reason = %#v", err)
	}
	return &connect.Response[server.Node]{
		Msg: &server.Node{
			Host: node.Reference().Host,
		},
	}, nil
}

func (is *InternalServer) Notify(ctx context.Context, req *connect.Request[server.Node]) (*connect.Response[emptypb.Empty], error) {
	fmt.Println("InternalServer::Notify", req.Msg.Host)
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	err := is.process.Notify(ctx, chord.NewRemoteNode(req.Msg.Host, is.process.Transport))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: notify failed. reason = %#v", err)
	}
	return &connect.Response[emptypb.Empty]{
		Msg: &emptypb.Empty{},
	}, nil
}

func (is *InternalServer) PutValueInner(ctx context.Context, req *connect.Request[server.PutValueInnerRequest]) (*connect.Response[server.PutValueInnerResponse], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	success, err := is.process.PutValue(ctx, &req.Msg.Key, &req.Msg.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: put value failed. reason = %#v", err)
	}
	return &connect.Response[server.PutValueInnerResponse]{
		Msg: &server.PutValueInnerResponse{
			Success: success,
		},
	}, nil
}

func (is *InternalServer) GetValueInner(ctx context.Context, req *connect.Request[server.GetValueInnerRequest]) (*connect.Response[server.GetValueInnerResponse], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	val, success, err := is.process.GetValue(ctx, &req.Msg.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: get value failed. reason = %#v", err)
	}

	return &connect.Response[server.GetValueInnerResponse]{
		Msg: &server.GetValueInnerResponse{
			Value:   *val,
			Success: success,
		},
	}, nil
}

func (is *InternalServer) DeleteValueInner(ctx context.Context, req *connect.Request[server.DeleteValueInnerRequest]) (*connect.Response[server.DeleteValueInnerResponse], error) {
	if is.process.IsShutdown {
		return nil, status.Errorf(codes.Unavailable, "server has started shutdown")
	}
	success, err := is.process.DeleteValue(ctx, &req.Msg.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "server: delete value failed. reason = %#v", err)
	}
	return &connect.Response[server.DeleteValueInnerResponse]{
		Msg: &server.DeleteValueInnerResponse{
			Success: success,
		},
	}, nil
}
