package core

import (
	"connectrpc.com/connect"
	"context"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/model"
	server2 "github.com/ryogrid/gord-overlay/server"
	"github.com/ryogrid/gord-overlay/serverconnect"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ExternalServer represents gRPC server to expose for gord users
type ExternalServer struct {
	port       string
	process    *chord.Process
	shutdownCh chan struct{}
}

// NewExternalServer creates an gRPC server to expose
func NewExternalServer(process *chord.Process, port string) *ExternalServer {
	return &ExternalServer{
		port:       port,
		process:    process,
		shutdownCh: make(chan struct{}, 1),
	}
}

//func (g *ExternalServer) newGrpcServer() *grpc.Server {
//	s := grpc.NewServer()
//	reflection.Register(s)
//	RegisterExternalServiceServer(s, g)
//	return s
//}

// Run runs chord server.
func (g *ExternalServer) Run() {
	go func() {
		//lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
		//if err != nil {
		//	log.Fatalf("failed to run server. reason: %#v", err)
		//}
		//grpcServer := g.newGrpcServer()
		//if err := grpcServer.Serve(lis); err != nil {
		//	log.Fatalf("failed to run server. reason: %#v", err)
		//}

		//mux := http.NewServeMux()
		//mux.Handle()
		//http.Serve(nil, mux)

		mux := http.NewServeMux()
		path, handler := serverconnect.NewExternalServiceHandler(g)
		mux.Handle(path, handler)
		http.ListenAndServe(
			"127.0.0.1",
			mux,
			//// Use h2c so we can serve HTTP/2 without TLS.
			//h2c.NewHandler(mux, &http2.Server{}),
		)
	}()
	log.Info("Running Gord server...")
	log.Infof("Gord is listening on %s:%s", g.process.Host, g.port)
	<-g.shutdownCh
}

// Shutdown shutdowns gRPC server.
func (g *ExternalServer) Shutdown() {
	g.shutdownCh <- struct{}{}
}

// FindHostForKey search for a given key's node.
// It is implemented for PublicService.
func (g *ExternalServer) FindHostForKey(ctx context.Context, req *connect.Request[server2.FindHostRequest]) (*connect.Response[server2.Node], error) {
	id := model.NewHashID(req.Msg.Key)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	return &connect.Response[server2.Node]{
		Msg: &server2.Node{
			Host: s.Reference().Host,
		},
	}, nil
}

func (g *ExternalServer) PutValue(ctx context.Context, req *connect.Request[server2.PutValueRequest]) (*connect.Response[server2.PutValueResponse], error) {
	id := model.NewHashID(req.Msg.Key)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::PutValue)
	success, err2 := s.PutValue(ctx, &req.Msg.Key, &req.Msg.Value)
	if err2 != nil {
		log.Errorf("External PutValue failed. reason: %#v", err)
		return nil, err2
	}
	return &connect.Response[server2.PutValueResponse]{
		Msg: &server2.PutValueResponse{
			Success: success,
		},
	}, nil
}

func (g *ExternalServer) GetValue(ctx context.Context, req *connect.Request[server2.GetValueRequest]) (*connect.Response[server2.GetValueResponse], error) {
	id := model.NewHashID(req.Msg.Key)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::GetValue)
	val, success, err2 := s.GetValue(ctx, &req.Msg.Key)
	if err2 != nil {
		log.Errorf("External GetValue failed. reason: %#v", err)
		return nil, err2
	}
	return &connect.Response[server2.GetValueResponse]{
		Msg: &server2.GetValueResponse{
			Value:   *val,
			Success: success,
		},
	}, nil
}

func (g *ExternalServer) DeleteValue(ctx context.Context, req *connect.Request[server2.DeleteValueRequest]) (*connect.Response[server2.DeleteValueResponse], error) {
	id := model.NewHashID(req.Msg.Key)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::DeleteValue)
	success, err2 := s.DeleteValue(ctx, &req.Msg.Key)
	if err2 != nil {
		log.Errorf("External DeleteValue failed. reason: %#v", err)
		return nil, err2
	}
	return &connect.Response[server2.DeleteValueResponse]{
		Msg: &server2.DeleteValueResponse{
			Success: success,
		},
	}, nil
}
