package chord

import (
	"context"
	"github.com/ryogrid/gord-overlay/pkg/model"
)

// RemoteNode represents remote nodes.
// To communicate with remote nodes, need to call rpc.
type RemoteNode struct {
	*model.NodeRef
	Transport
}

func NewRemoteNode(host string, transport Transport) RingNode {
	return &RemoteNode{
		NodeRef:   model.NewNodeRef(host),
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
	return r.FindSuccessorByTableRPC(ctx, r.NodeRef, id)
}

func (r *RemoteNode) FindClosestPrecedingNode(ctx context.Context, id model.HashID) (RingNode, error) {
	return r.FindClosestPrecedingNodeRPC(ctx, r.NodeRef, id)
}

func (r *RemoteNode) Notify(ctx context.Context, node RingNode) error {
	return r.NotifyRPC(ctx, r.NodeRef, node.Reference())
}

func (r *RemoteNode) PutValueInner(ctx context.Context, key *string, value *string) (bool, error) {
	// TODO: need to implement RemoteNode::PutValueInner
	return r.PutValueInnerRPC(ctx, key, value)
}

func (r *RemoteNode) GetValueInnerRPC(ctx context.Context, key *string) (*string, bool, error) {
	// TODO: need to implement RemoteNode::GetValueInner
	return r.GetValueInnerRPC(ctx, key)
}

func (r *RemoteNode) DeleteValueInnerRPC(ctx context.Context, key *string) (bool, error) {
	// TODO: need to implement RemoteNode::DeleteValueInner
	return r.DeleteValueInnerRPC(ctx, key)
}
