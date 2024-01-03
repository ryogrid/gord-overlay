package chord

import (
	"context"
	"github.com/ryogrid/gord-overlay/pkg/model"
)

// MockTransport does nothing
type MockTransport struct{}

// PingRPC does nothing
func (m *MockTransport) PingRPC(ctx context.Context, to *model.NodeRef) error {
	return nil
}

// SuccessorsRPC does nothing
func (m *MockTransport) SuccessorsRPC(ctx context.Context, to *model.NodeRef) ([]RingNode, error) {
	return nil, nil
}

// PredecessorRPC does nothing
func (m *MockTransport) PredecessorRPC(ctx context.Context, to *model.NodeRef) (RingNode, error) {
	return nil, nil
}

// FindSuccessorByTableRPC does nothing
func (m *MockTransport) FindSuccessorByTableRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error) {
	return nil, nil
}

// FindSuccessorByListRPC does nothing
func (m *MockTransport) FindSuccessorByListRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error) {
	return nil, nil
}

// FindClosestPrecedingNodeRPC does nothing
func (m *MockTransport) FindClosestPrecedingNodeRPC(ctx context.Context, to *model.NodeRef, id model.HashID) (RingNode, error) {
	return nil, nil
}

// NotifyRPC does nothing
func (m *MockTransport) NotifyRPC(ctx context.Context, to *model.NodeRef, node *model.NodeRef) error {
	return nil
}

// Shutdown does nothing
func (m *MockTransport) Shutdown() {
}

func (m *MockTransport) PutValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string, value *string) (bool, error) {
	return false, nil
}
func (m *MockTransport) GetValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (*string, bool, error) {
	return nil, false, nil
}
func (m *MockTransport) DeleteValueInnerRPC(ctx context.Context, to *model.NodeRef, key *string) (bool, error) {
	return false, nil
}
