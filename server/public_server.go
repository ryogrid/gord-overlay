package server

import (
	"context"
	"fmt"
	"github.com/ryogird/gord-overlay/api_external"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/pkg/model"
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

func (g *ExternalServer) newRpcServer() *api_external.Server {
	// TODO: need to implement ExternalServer::newRpcServer method
	//s := grpc.NewServer()
	//reflection.Register(s)
	//RegisterExternalServiceServer(s, g)
	s, err := api_external.NewServer(g)
	if err != nil {
		panic(err)
	}
	return s
}

// Run runs chord server.
func (g *ExternalServer) Run() {
	go func() {
		//lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
		//if err != nil {
		//	log.Fatalf("failed to run server. reason: %#v", err)
		//}
		rpcServer := g.newRpcServer()
		err := http.ListenAndServe("0.0.0.0:"+g.port, rpcServer)
		if err != nil {
			panic(err)
		}
		//if err := grpcServer.Serve(lis); err != nil {
		//	log.Fatalf("failed to run server. reason: %#v", err)
		//}
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
func (g *ExternalServer) ExternalServiceFindHostForKey(ctx context.Context, params api_external.ExternalServiceFindHostForKeyParams) (*api_external.ServerNode, error) {
	//func (g *ExternalServer) FindHostForKey(ctx context.Context, req *FindHostRequest) (*Node, error) {
	id := model.NewHashID(params.Key.Value)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	return &api_external.ServerNode{
		Host: api_external.NewOptString(s.Reference().Host),
	}, nil
}

func (g *ExternalServer) ExternalServicePutValue(ctx context.Context, params api_external.ExternalServicePutValueParams) (*api_external.ServerPutValueResponse, error) {
	//func (g *ExternalServer) PutValue(ctx context.Context, req *PutValueRequest) (*PutValueResponse, error) {
	id := model.NewHashID(params.Key.Value)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::PutValue)
	success, err2 := s.PutValue(ctx, &params.Key.Value, &params.Value.Value)
	if err2 != nil {
		log.Errorf("External PutValue failed. reason: %#v", err)
		return nil, err2
	}
	return &api_external.ServerPutValueResponse{
		Success: api_external.NewOptBool(success),
	}, nil
}

func (g *ExternalServer) ExternalServiceGetValue(ctx context.Context, params api_external.ExternalServiceGetValueParams) (*api_external.ServerGetValueResponse, error) {
	//func (g *ExternalServer) GetValue(ctx context.Context, req *GetValueRequest) (*GetValueResponse, error) {
	id := model.NewHashID(params.Key.Value)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::GetValue)
	val, success, err2 := s.GetValue(ctx, &params.Key.Value)
	if err2 != nil {
		log.Errorf("External GetValue failed. reason: %#v", err)
		return nil, err2
	}
	return &api_external.ServerGetValueResponse{
		Value:   api_external.NewOptString(*val),
		Success: api_external.NewOptBool(success),
	}, nil
}

func (g *ExternalServer) ExternalServiceDeleteValue(ctx context.Context, params api_external.ExternalServiceDeleteValueParams) (*api_external.ServerDeleteValueResponse, error) {
	//func (g *ExternalServer) DeleteValue(ctx context.Context, req *api_external.DeleteValueRequest) (*api_external.ServerDeleteValueResponse, error) {
	id := model.NewHashID(params.Key.Value)
	s, err := g.process.FindSuccessorByTable(ctx, id)
	if err != nil {
		log.Errorf("FindHostForKey failed. reason: %#v", err)
		return nil, err
	}
	// TODO: need to consider repllication (ExternalServer::DeleteValue)
	success, err2 := s.DeleteValue(ctx, &params.Key.Value)
	if err2 != nil {
		log.Errorf("External DeleteValue failed. reason: %#v", err)
		return nil, err2
	}
	return &api_external.ServerDeleteValueResponse{
		Success: api_external.NewOptBool(success),
	}, nil
}

func (g *ExternalServer) NewError(ctx context.Context, err error) *api_external.ErrorStatusCode {
	return &api_external.ErrorStatusCode{Response: api_external.Error{Message: fmt.Sprintf("%v", err), Code: -1}}
}
