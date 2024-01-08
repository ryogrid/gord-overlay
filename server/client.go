package server

import (
	"context"
	"github.com/ryogird/gord-overlay/api_internal"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/pkg/model"
	"google.golang.org/grpc"
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

// TODO: Enable mTLS
// TODO: Add conn pool capacity limit for file descriptors.
func (c *ApiClient) getRpcClient(address string) (*api_internal.Client, error) {
	//c.poolLock.Lock()
	//defer c.poolLock.Unlock()
	//conn, ok := c.connPool[address]
	//if ok {
	//	return NewInternalServiceClient(conn), nil
	//}
	//
	//conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, c.serverPort), grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	return nil, err
	//}
	//c.connPool[address] = conn
	//return NewInternalServiceClient(conn), nil
	customClient := http.DefaultClient
	// customClient.Transport = nil // TODO: need to implement transport (ApiClient::getRpcClient method)
	ret, err := api_internal.NewClient("http://"+address+":"+c.serverPort, api_internal.WithClient(customClient))
	return ret, err
}

func (c *ApiClient) createRingNodeFrom(node *api_internal.ServerNode) chord.RingNode {
	if c.hostNode.Host == node.Host.Value {
		return c.hostNode
	}
	return chord.NewRemoteNode(node.Host.Value, c)
}

func (c *ApiClient) PingRPC(ctx context.Context, to *model.NodeRef) error {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err = client.InternalServicePing(ctx)
	if err != nil {
		return handleError(err)
	}
	return nil
}

func (c *ApiClient) SuccessorsRPC(ctx context.Context, to *model.NodeRef) ([]chord.RingNode, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	nodes, err := client.InternalServiceSuccessors(ctx)
	if err != nil {
		return nil, handleError(err)
	}
	ringNodes := make([]chord.RingNode, len(nodes.Nodes))
	for i, node := range nodes.Nodes {
		ringNodes[i] = c.createRingNodeFrom(&node)
	}
	return ringNodes, nil
}

func (c *ApiClient) PredecessorRPC(ctx context.Context, to *model.NodeRef) (chord.RingNode, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.InternalServicePredecessor(ctx)
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindSuccessorByTableRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.InternalServiceFindSuccessorByTable(ctx, api_internal.InternalServiceFindSuccessorByTableParams{ID: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindSuccessorByListRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.InternalServiceFindSuccessorByList(ctx, api_internal.InternalServiceFindSuccessorByListParams{ID: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) FindClosestPrecedingNodeRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (chord.RingNode, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	node, err := client.InternalServiceFindClosestPrecedingNode(ctx, api_internal.InternalServiceFindClosestPrecedingNodeParams{ID: id})
	if err != nil {
		return nil, handleError(err)
	}
	return c.createRingNodeFrom(node), nil
}

func (c *ApiClient) NotifyRPC(ctx context.Context, to *model.NodeRef, node *model.NodeRef) error {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err = client.InternalServiceNotify(ctx, api_internal.InternalServiceNotifyParams{Host: api_internal.NewOptString(node.Host)})
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

func (c *ApiClient) PutValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string, value *string) (bool, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.InternalServicePutValueInner(ctx, api_internal.InternalServicePutValueInnerParams{Key: api_internal.OptString{Value: *key}, Value: api_internal.NewOptString(*value)})
	if err != nil {
		return false, handleError(err)
	}
	return resp.Success.Value, nil
}

func (c *ApiClient) GetValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (*string, bool, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return nil, false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.InternalServiceGetValueInner(ctx, api_internal.InternalServiceGetValueInnerParams{Key: api_internal.NewOptString(*key)})
	if err != nil {
		return nil, false, handleError(err)
	}
	return &resp.Value.Value, resp.Success.Value, nil
}

func (c *ApiClient) DeleteValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (bool, error) {
	client, err := c.getRpcClient(to.Host)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	err = client.InternalServiceDeleteValueInner(ctx, api_internal.InternalServiceDeleteValueInnerParams{Key: api_internal.NewOptString(*key)})
	if err != nil {
		return false, handleError(err)
	}
	return true, nil
}
