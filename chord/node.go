package chord

import (
	"context"
	"github.com/ryogrid/gord-overlay/pkg/model"
)

// RingNode represents a node of Chord Ring
type RingNode interface {
	Ping(ctx context.Context) error
	Reference() *model.NodeRef
	GetSuccessors(ctx context.Context) ([]RingNode, error)
	GetPredecessor(ctx context.Context) (RingNode, error)
	FindSuccessorByTable(ctx context.Context, id model.HashID) (RingNode, error)
	FindSuccessorByList(ctx context.Context, id model.HashID) (RingNode, error)
	FindClosestPrecedingNode(ctx context.Context, id model.HashID) (RingNode, error)
	Notify(ctx context.Context, node RingNode) error
	PutValue(ctx context.Context, key *string, value *string) (bool, error)
	GetValue(ctx context.Context, key *string) (*string, bool, error)
	DeleteValue(ctx context.Context, key *string) (bool, error)
}

// Transport represents rpc to remote node
type Transport interface {
	PingRPC(ctx context.Context, to *model.NodeRef) error
	SuccessorsRPC(ctx context.Context, to *model.NodeRef) ([]RingNode, error)
	PredecessorRPC(ctx context.Context, to *model.NodeRef) (RingNode, error)
	FindSuccessorByTableRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error)
	FindSuccessorByListRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error)
	FindClosestPrecedingNodeRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error)
	NotifyRPC(ctx context.Context, to *model.NodeRef, node *model.NodeRef) error
	Shutdown()
	PutValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string, value *string) (bool, error)
	GetValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (*string, bool, error)
	DeleteValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (bool, error)
}
