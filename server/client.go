package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/pkg/model"
	"google.golang.org/grpc"
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

func (c *ApiClient) createRingNodeFrom(node *Node) chord.RingNode {
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
	_, err = client.Ping(ctx, &empty.Empty{})
	if err != nil {
		return handleError(err)
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
	nodes, err := client.Successors(ctx, &empty.Empty{})
	if err != nil {
		return nil, handleError(err)
	}
	ringNodes := make([]chord.RingNode, len(nodes.Nodes))
	for i, node := range nodes.Nodes {
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
	node, err := client.Predecessor(ctx, &empty.Empty{})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindSuccessorByTableRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindSuccessorByTable(ctx, &FindRequest{Id: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindSuccessorByListRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindSuccessorByList(ctx, &FindRequest{Id: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindClosestPrecedingNodeRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.FindClosestPrecedingNode(ctx, &FindRequest{Id: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) NotifyRPC(ctx context.Context, to *model.NodeRef, node *model.NodeRef) error {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err = client.Notify(ctx, &Node{
		Host: node.Host,
	})
	if err != nil {
		return handleError(err)
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

func (c *ApiClient) PutValueInnerRPC(ctx context.Context, key *string, value *string) (bool, error) {
	// TODO: need to implement ApiClient::PutValueInnerRPC
	panic("not implemented")
}

func (c *ApiClient) GetValueInnerRPC(ctx context.Context, key *string) (*string, bool, error) {
	// TODO: need to implement ApiClient::GetValueInnerRPC
	panic("not implemented")
}

func (c *ApiClient) DeleteValueInnerRPC(ctx context.Context, key *string) (bool, error) {
	// TODO: need to implement ApiClient::DeleteValueInnerRPC
	panic("not implemented")
}
