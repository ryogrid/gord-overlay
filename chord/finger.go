package chord

import (
	"github.com/ryogrid/gord-overlay/model"
	"math/big"
)

// Finger represents an element of routing table
type Finger struct {
	Index int
	ID    model.HashID
	Node  RingNode
}

// NewFingerTable creates a finger table.
func NewFingerTable(id model.HashID) []*Finger {
	table := make([]*Finger, model.BitSize)
	for i := range table {
		table[i] = NewFinger(id, i, nil)
	}
	return table
}

// NewFinger creates a finger.
// index is an order of finger table.
// node is this finger table's owner.
func NewFinger(id model.HashID, index int, successor RingNode) *Finger {
	nodeID := big.NewInt(0).SetBytes(id)
	base := big.NewInt(2)
	offset := big.NewInt(0).Exp(base, big.NewInt(int64(index)), nil)       // 2^i
	sum := big.NewInt(0).Add(nodeID, offset)                               // n + 2^i
	ring := big.NewInt(0).Exp(base, big.NewInt(int64(model.BitSize)), nil) //2^m
	fingerIDBytes := big.NewInt(0).Mod(sum, ring).Bytes()
	return &Finger{
		Index: index,
		ID:    model.BytesToHashID(fingerIDBytes),
		Node:  successor,
	}
}
