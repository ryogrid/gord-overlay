package chord

import (
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/model"
)

// RemoteNode represents remote nodes.
// To communicate with remote nodes, need to call rpc.
type RemoteNode struct {
	*model.NodeRef
	Transport
}

func NewRemoteNode(hostAndPort string, transport Transport) RingNode {
	return &RemoteNode{
		NodeRef:   model.NewNodeRef(hostAndPort),
		Transport: transport,
	}
}

func (r *RemoteNode) Ping(ctx context.Context) error {
	return r.PingRPC(ctx, r.NodeRef)
}

func (r *RemoteNode) Reference() *model.NodeRef {
	return r.NodeRef
}

func (r *RemoteNode) GetSuccessors(ctx context.Context) ([]RingNode, error) {
	return r.SuccessorsRPC(ctx, r.NodeRef)
}

func (r *RemoteNode) GetPredecessor(ctx context.Context) (RingNode, error) {
	return r.PredecessorRPC(ctx, r.NodeRef)
}

func (r *RemoteNode) FindSuccessorByList(ctx context.Context, id model.HashID) (RingNode, error) {
	return r.FindSuccessorByListRPC(ctx, r.NodeRef, id)
}

func (r *RemoteNode) FindSuccessorByTable(ctx context.Context, id model.HashID) (RingNode, error) {
	fmt.Println("RemoteNode::FindSuccessorByTable", r.NodeRef.Host, id)
	return r.FindSuccessorByTableRPC(ctx, r.NodeRef, id)
}

func (r *RemoteNode) FindClosestPrecedingNode(ctx context.Context, id model.HashID) (RingNode, error) {
	return r.FindClosestPrecedingNodeRPC(ctx, r.NodeRef, id)
}

func (r *RemoteNode) Notify(ctx context.Context, node RingNode) error {
	return r.NotifyRPC(ctx, r.NodeRef, node.Reference())
}

func (r *RemoteNode) PutValue(ctx context.Context, key *string, value *string) (bool, error) {
	return r.PutValueInnerRPC(ctx, r.NodeRef, key, value)
}

func (r *RemoteNode) GetValue(ctx context.Context, key *string) (*string, bool, error) {
	return r.GetValueInnerRPC(ctx, r.NodeRef, key)
}

func (r *RemoteNode) DeleteValue(ctx context.Context, key *string) (bool, error) {
	return r.DeleteValueInnerRPC(ctx, r.NodeRef, key)
}
