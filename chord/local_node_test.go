package chord

import (
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/pkg/model"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func createNodes(n int) []*LocalNode {
	nodes := make([]*LocalNode, n)
	for i := 0; i < n; i++ {
		node := NewLocalNode(fmt.Sprintf("gord%d", i))
		node.ID = big.NewInt(int64(i) + 1).Bytes()
		nodes[i] = node
	}
	return nodes
}

func TestLocalNode_CreateRing(t *testing.T) {
	node := NewLocalNode("gord")
	node.CreateRing()
	assert.NotNil(t, node.predecessor, nil)
	assert.Equal(t, node.ID, node.predecessor.Reference().ID)
	assert.Equal(t, model.BitSize/2, cap(node.successors.nodes))
	assert.Equal(t, 1, len(node.successors.nodes))
	assert.Equal(t, node.successors.nodes[0].Reference().ID, node.ID)
	assert.Equal(t, len(node.fingerTable), model.BitSize)
	for _, finger := range node.fingerTable {
		assert.Equal(t, finger.Node.Reference().ID, node.ID)
	}
}

func TestLocalNode_JoinRing(t *testing.T) {
	ctx := context.Background()
	nodes := createNodes(3)
	node1, node2, node3 := nodes[0], nodes[1], nodes[2]
	node1.CreateRing()
	// node2 joins in chord rinmg!
	node2.JoinRing(ctx, node1)
	assert.Equal(t, node2.successors.nodes[0].Reference().ID, node1.ID)
	assert.Equal(t, node1.predecessor.Reference().ID, node2.ID)

	// node3 joins in chord ring!
	node3.JoinRing(ctx, node2)
	assert.Equal(t, node3.successors.nodes[0].Reference().ID, node1.ID)
	assert.Equal(t, node1.predecessor.Reference().ID, node3.ID)
}

func TestLocalNode_Notify(t *testing.T) {
	ctx := context.Background()
	nodes := createNodes(3)
	node1, node2, node3 := nodes[0], nodes[1], nodes[2]
	assert.NoError(t, node3.Notify(ctx, node1))
	assert.Equal(t, node1.ID, node3.predecessor.Reference().ID)
	assert.NoError(t, node2.Notify(ctx, node1))
	assert.Equal(t, node1.ID, node2.predecessor.Reference().ID)
}

func TestLocalNode_JoinSuccessors(t *testing.T) {
	nodes := createNodes(3)
	node1, node2, node3 := nodes[0], nodes[1], nodes[2]
	node1.CreateRing()

	node1.JoinSuccessors(0, []RingNode{node2})
	assert.Equal(t, model.BitSize/2, cap(node1.successors.nodes))
	assert.Equal(t, []RingNode{node2}, node1.successors.nodes)

	node1.JoinSuccessors(1, []RingNode{node1})
	assert.Equal(t, []RingNode{node2, node1}, node1.successors.nodes)
	assert.Equal(t, model.BitSize/2, cap(node1.successors.nodes))

	node1.JoinSuccessors(1, []RingNode{})
	assert.Equal(t, []RingNode{node2, node1}, node1.successors.nodes)
	assert.Equal(t, model.BitSize/2, cap(node1.successors.nodes))

	node1.JoinSuccessors(cap(node1.successors.nodes), []RingNode{node3})
	assert.Equal(t, []RingNode{node2, node1}, node1.successors.nodes)
	assert.Equal(t, model.BitSize/2, cap(node1.successors.nodes))

	node1.JoinSuccessors(2, []RingNode{node1, node3, node2})
	assert.Equal(t, []RingNode{node2, node1, node3}, node1.successors.nodes)
	assert.Equal(t, model.BitSize/2, cap(node1.successors.nodes))
}

func TestLocalNode_PutSuccessor(t *testing.T) {
	nodes := createNodes(3)
	node1, node2, node3 := nodes[0], nodes[1], nodes[2]
	node1.CreateRing()

	node1.PutSuccessor(node2)
	assert.Equal(t, 2, len(node1.successors.nodes))
	assert.Equal(t, node2.ID, node1.successors.nodes[0].Reference().ID)
	assert.Equal(t, node1.ID, node1.successors.nodes[1].Reference().ID)

	node1.PutSuccessor(node3)
	assert.Equal(t, 3, len(node1.successors.nodes))
	assert.Equal(t, node3.ID, node1.successors.nodes[0].Reference().ID)
	assert.Equal(t, node2.ID, node1.successors.nodes[1].Reference().ID)
	assert.Equal(t, node1.ID, node1.successors.nodes[2].Reference().ID)
}
