package core

import (
	"connectrpc.com/connect"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ryogird/gord-overlay/server"
	"github.com/ryogird/gord-overlay/serverconnect"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"sync"
	"time"
)

type ApiClient struct {
	hostNode   *chord.LocalNode
	serverPort string
	timeout    time.Duration
	connPool   map[string]*grpc.ClientConn
	poolLock   sync.Mutex
	opts       grpc.CallOption
}

func NewChordApiClient(hostNode *chord.LocalNode, port string, timeout time.Duration) chord.Transport {
	return &ApiClient{
		hostNode:   hostNode,
		serverPort: port,
		timeout:    timeout,
		connPool:   map[string]*grpc.ClientConn{},
	}
}

/*
// TODO: Enable mTLS
// TODO: Add conn pool capacity limit for file descriptors.
func (c *ApiClient) getGrpcConn(address string) (InternalServiceClient, error) {
	c.poolLock.Lock()
	defer c.poolLock.Unlock()
	conn, ok := c.connPool[address]
	if ok {
		return NewInternalServiceClient(conn), nil
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, c.serverPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.connPool[address] = conn
	return NewInternalServiceClient(conn), nil
}
*/

func (c *ApiClient) getGrpcConn(address string) (serverconnect.InternalServiceClient, error) {
	//cli := http.DefaultClient
	//overlayTransport := &http.Transport{
	//	Proxy: http.ProxyFromEnvironment,
	//	//DialContext: defaultTransportDialContext(&net.Dialer{
	//	//	Timeout:   30 * time.Second,
	//	//	KeepAlive: 30 * time.Second,
	//	//}),
	//	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
	//		return nil, nil
	//	},
	//	ForceAttemptHTTP2:     false,
	//	MaxIdleConns:          100,
	//	IdleConnTimeout:       90 * time.Second,
	//	TLSHandshakeTimeout:   10 * time.Second,
	//	ExpectContinueTimeout: 1 * time.Second,
	//}
	//cli.Transport = overlayTransport

	return serverconnect.NewInternalServiceClient(http.DefaultClient, address), nil
}

func (c *ApiClient) createRingNodeFrom(node *server.Node) chord.RingNode {
	if c.hostNode.Host == node.Host {
		return c.hostNode
	}
	return chord.NewRemoteNode(node.Host, c)
}

func (c *ApiClient) PingRPC(ctx context.Context, to *model.NodeRef) error {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err = client.Ping(ctx, &connect.Request[emptypb.Empty]{Msg: &empty.Empty{}})
	if err != nil {
		return server.HandleError(err)
	}
	return nil
}

func (c *ApiClient) SuccessorsRPC(ctx context.Context, to *model.NodeRef) ([]chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	nodes, err := client.Successors(ctx, &connect.Request[emptypb.Empty]{Msg: &emptypb.Empty{}})
	if err != nil {
		return nil, server.HandleError(err)
	}
	ringNodes := make([]chord.RingNode, len(nodes.Msg.Nodes))
	for i, node := range nodes.Msg.Nodes {
		ringNodes[i] = c.createRingNodeFrom(node)
	}
	return ringNodes, nil
}

func (c *ApiClient) PredecessorRPC(ctx context.Context, to *model.NodeRef) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.Predecessor(ctx, &connect.Request[emptypb.Empty]{Msg: &emptypb.Empty{}})
	if err != nil {
		return nil, server.HandleError(err)
	}
	return c.createRingNodeFrom(node.Msg), nil
}

func (c *ApiClient) FindSuccessorByTableRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindSuccessorByTable(ctx, &connect.Request[server.FindRequest]{Msg: &server.FindRequest{Id: id}})
	if err != nil {
		return nil, server.HandleError(err)
	}
	return c.createRingNodeFrom(node.Msg), nil
}

func (c *ApiClient) FindSuccessorByListRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindSuccessorByList(ctx, &connect.Request[server.FindRequest]{Msg: &server.FindRequest{Id: id}})
	if err != nil {
		return nil, server.HandleError(err)
	}
	return c.createRingNodeFrom(node.Msg), nil
}

func (c *ApiClient) FindClosestPrecedingNodeRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindClosestPrecedingNode(ctx, &connect.Request[server.FindRequest]{Msg: &server.FindRequest{Id: id}})
	if err != nil {
		return nil, server.HandleError(err)
	}
	return c.createRingNodeFrom(node.Msg), nil
}

func (c *ApiClient) NotifyRPC(ctx context.Context, to *model.NodeRef, node *model.NodeRef) error {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err = client.Notify(ctx, &connect.Request[server.Node]{Msg: &server.Node{
		Host: node.Host,
	}},
	)
	if err != nil {
		return server.HandleError(err)
	}
	return nil
}

func (c *ApiClient) Shutdown() {
	c.poolLock.Lock()
	defer c.poolLock.Unlock()
	for _, conn := range c.connPool {
		conn.Close()
	}
}

func (c *ApiClient) PutValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string, value *string) (bool, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.PutValueInner(ctx, &connect.Request[server.PutValueInnerRequest]{Msg: &server.PutValueInnerRequest{Key: *key, Value: *value}})
	if err != nil {
		return false, server.HandleError(err)
	}
	return resp.Msg.Success, nil
}

func (c *ApiClient) GetValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (*string, bool, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.GetValueInner(ctx, &connect.Request[server.GetValueInnerRequest]{Msg: &server.GetValueInnerRequest{Key: *key}})
	if err != nil {
		return nil, false, server.HandleError(err)
	}
	return &resp.Msg.Value, resp.Msg.Success, nil
}

func (c *ApiClient) DeleteValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (bool, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.DeleteValueInner(ctx, &connect.Request[server.DeleteValueInnerRequest]{Msg: &server.DeleteValueInnerRequest{Key: *key}})
	if err != nil {
		return false, server.HandleError(err)
	}
	return resp.Msg.Success, nil
}
