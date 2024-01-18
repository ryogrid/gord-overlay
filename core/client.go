package core

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/model"
	"github.com/ryogrid/gord-overlay/server"
	"github.com/ryogrid/gord-overlay/serverconnect"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/ryogrid/gossip-overlay/util"
	"github.com/weaveworks/mesh"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"net/http"
	"sync"
	"time"
)

type ApiClient struct {
	olPeer     *overlay.OverlayPeer
	hostNode   *chord.LocalNode
	timeout    time.Duration
	clientPool map[string]*http.Client
	poolLock   sync.Mutex
	opts       grpc.CallOption
}

func NewChordApiClient(hostNode *chord.LocalNode, olPeer *overlay.OverlayPeer, timeout time.Duration) chord.Transport {
	return &ApiClient{
		olPeer:     olPeer,
		hostNode:   hostNode,
		timeout:    timeout,
		clientPool: make(map[string]*http.Client),
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
	//	DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
	//		fmt.Println("DialTLSContext", network, addr)
	//		//return c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint64(addr))), nil
	//		return tls.Client(c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint16(addr))), &tls.Config{InsecureSkipVerify: true}), nil
	//	},
	//	//DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
	//	//	fmt.Println("DialContext", network, addr)
	//	//	//return c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint64(addr))), nil
	//	//	return c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint16(addr))), nil
	//	//},
	//	ForceAttemptHTTP2:     true,              //false,             //true,
	//	MaxIdleConns:          100,               //0,
	//	IdleConnTimeout:       180 * time.Second, //0,//1 * time.Second,
	//	TLSHandshakeTimeout:   10 * time.Second,
	//	ExpectContinueTimeout: 90 * time.Second,
	//	MaxIdleConnsPerHost:   3, //100, //0,
	//	//DisableKeepAlives:     true,
	//}
	cli := &http.Client{
		Transport: &http2.Transport{
			//AllowHTTP: true,
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				fmt.Println("DialContext", network, addr)
				//return c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint64(addr))), nil
				return c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint16(addr))), nil
				//return tls.Client(c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint16(addr))), &tls.Config{InsecureSkipVerify: true}), nil
				//return net.Dial(network, addr)
			},
		},
	}
	//cli := tls.Client(c.olPeer.OpenStreamToTargetPeer(mesh.PeerName(util.NewHashIDUint16(address))), &tls.Config{InsecureSkipVerify: true})
	//cli.Transport = overlayTransport

	//return serverconnect.NewInternalServiceClient(cli, "http://"+address), nil
	return serverconnect.NewInternalServiceClient(cli, "http://"+address), nil

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
	_, err = client.Ping(ctx, connect.NewRequest(&emptypb.Empty{}))
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
	nodes, err := client.Successors(ctx, connect.NewRequest(&emptypb.Empty{}))
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
	node, err := client.Predecessor(ctx, connect.NewRequest(&emptypb.Empty{}))
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
	node, err := client.FindSuccessorByTable(ctx, connect.NewRequest(&server.FindRequest{Id: id}))
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
	node, err := client.FindSuccessorByList(ctx, connect.NewRequest(&server.FindRequest{Id: id}))
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
	node, err := client.FindClosestPrecedingNode(ctx, connect.NewRequest(&server.FindRequest{Id: id}))
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
	_, err = client.Notify(ctx, connect.NewRequest(&server.Node{
		Host: node.Host,
	}),
	)
	if err != nil {
		return server.HandleError(err)
	}
	return nil
}

func (c *ApiClient) Shutdown() {
	//c.poolLock.Lock()
	//defer c.poolLock.Unlock()
	//for _, client := range c.clientPool {
	//	client.Close()
	//}
}

func (c *ApiClient) PutValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string, value *string) (bool, error) {
	client, err := c.getGrpcConn(to.Host)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := client.PutValueInner(ctx, connect.NewRequest(&server.PutValueInnerRequest{Key: *key, Value: *value}))
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
	resp, err := client.GetValueInner(ctx, connect.NewRequest(&server.GetValueInnerRequest{Key: *key}))
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
	resp, err := client.DeleteValueInner(ctx, connect.NewRequest(&server.DeleteValueInnerRequest{Key: *key}))
	if err != nil {
		return false, server.HandleError(err)
	}
	return resp.Msg.Success, nil
}
