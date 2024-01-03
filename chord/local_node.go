package chord

import (
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/pkg/model"
	"sync"
)

// exclusiveNodeList represents node list.
// It restricts no overlapped host nodes.
type exclusiveNodeList struct {
	nodes   []RingNode
	hostMap map[string]struct{}
}

func newNodeList(cap int) *exclusiveNodeList {
	return &exclusiveNodeList{
		nodes:   emptyNodes(cap),
		hostMap: map[string]struct{}{},
	}
}

func (q *exclusiveNodeList) refreshHostMap() {
	hostMap := map[string]struct{}{}
	for _, node := range q.nodes {
		hostMap[node.Reference().Host] = struct{}{}
	}
	q.hostMap = hostMap
}

func (q *exclusiveNodeList) hasHostKey(host string) bool {
	_, ok := q.hostMap[host]
	return ok
}

func (q *exclusiveNodeList) head() (RingNode, error) {
	if len(q.nodes) == 0 {
		return nil, ErrNoSuccessorAlive
	}
	return q.nodes[0], nil
}

func emptyNodes(cap int) []RingNode {
	return make([]RingNode, 0, cap)
}

func (q *exclusiveNodeList) appendHead(node RingNode) {
	if node == nil {
		return
	}
	if q.hasHostKey(node.Reference().Host) {
		return
	}

	newNodes := append(emptyNodes(cap(q.nodes)), node)
	if len(q.nodes) >= cap(q.nodes) {
		q.nodes = append(newNodes, q.nodes[:len(q.nodes)-1]...)
		return
	}
	q.nodes = append(newNodes, q.nodes[:]...)
	q.hostMap[node.Reference().Host] = struct{}{}
}

func (q *exclusiveNodeList) join(offset int, nodes []RingNode) {
	if len(nodes) == 0 {
		return
	}
	if cap(q.nodes) <= offset {
		return
	}
	if len(nodes) > (cap(q.nodes) - offset) {
		nodes = nodes[:cap(q.nodes)-offset]
	}

	q.nodes = q.nodes[0:offset]
	q.refreshHostMap()
	for _, node := range nodes {
		if q.hasHostKey(node.Reference().Host) {
			continue
		}
		q.nodes = append(q.nodes, node)
	}
	q.refreshHostMap()
}

// LocalNode represents local host node.
type LocalNode struct {
	*model.NodeRef

	fingerTable []*Finger
	successors  *exclusiveNodeList
	predecessor RingNode
	isShutdown  bool
	// string => *string
	storedValues sync.Map
	lock         sync.Mutex
}

// NewLocalNode creates a local node.
func NewLocalNode(host string) *LocalNode {
	id := model.NewHashID(host)
	return &LocalNode{
		NodeRef:     model.NewNodeRef(host),
		fingerTable: NewFingerTable(id),
	}
}

func (l *LocalNode) initSuccessors(suc RingNode) {
	l.successors = newNodeList(model.BitSize / 2)
	l.PutSuccessor(suc)
}

func (l *LocalNode) Shutdown() {
	l.isShutdown = true
}

func (l *LocalNode) CreateRing() {
	l.initSuccessors(l)
	l.predecessor = l
	for _, finger := range l.fingerTable {
		finger.Node = l
	}
}

func (l *LocalNode) JoinRing(ctx context.Context, existNode RingNode) error {
	successor, err := existNode.FindSuccessorByTable(ctx, l.ID)
	if err != nil {
		return fmt.Errorf("find successor failed. err = %#v", err)
	}
	l.initSuccessors(successor)

	firstSuc, err := l.successors.head()
	if err != nil {
		return err
	}

	err = firstSuc.Notify(ctx, l)
	if err != nil {
		return fmt.Errorf("notify failed. err = %#v", err)
	}

	successors, err := firstSuc.GetSuccessors(ctx)
	if err != nil {
		return fmt.Errorf("get successors failed. err = %#v", err)
	}

	l.JoinSuccessors(1, successors)
	return nil
}

func (l *LocalNode) JoinSuccessors(offset int, successors []RingNode) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.successors.join(offset, successors)
}

func (l *LocalNode) PutSuccessor(suc RingNode) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.successors.appendHead(suc)
	l.fingerTable[0].Node = suc
}

func (l *LocalNode) Ping(_ context.Context) error {
	if l.isShutdown {
		return ErrNodeUnavailable
	}
	return nil
}

func (l *LocalNode) Reference() *model.NodeRef {
	return l.NodeRef
}

func (l *LocalNode) GetSuccessors(_ context.Context) ([]RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	return l.successors.nodes, nil
}

func (l *LocalNode) GetPredecessor(_ context.Context) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	return l.predecessor, nil
}

func (l *LocalNode) FindSuccessorByList(ctx context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	for _, successor := range l.successors.nodes {
		if id.Between(l.ID, successor.Reference().ID) {
			return successor, nil
		}
	}
	return nil, ErrNotFound
}

func (l *LocalNode) FindSuccessorByTable(ctx context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	node, err := l.findPredecessor(ctx, id)
	if err != nil {
		return l.FindSuccessorByList(ctx, id)
	}
	successors, err := node.GetSuccessors(ctx)
	if err != nil {
		return nil, err
	}
	for _, successor := range successors {
		if err := successor.Ping(ctx); err == nil {
			return successor, nil
		}
	}
	return nil, ErrNoSuccessorAlive
}

func (l *LocalNode) findPredecessor(ctx context.Context, id model.HashID) (RingNode, error) {
	var (
		targetNode RingNode = l
	)
	for {
		successors, err := targetNode.GetSuccessors(ctx)
		if err != nil {
			return nil, err
		}
		if successors == nil || len(successors) <= 0 {
			return nil, ErrNotFound
		}
		suc := successors[0]
		if targetNode.Reference().ID.Equals(suc.Reference().ID) {
			return targetNode, nil
		}
		if id.Between(targetNode.Reference().ID, suc.Reference().ID.Add(1)) {
			break
		}
		node, err := targetNode.FindClosestPrecedingNode(ctx, id)
		if err != nil {
			return nil, ErrNotFound
		}
		targetNode = node
	}
	return targetNode, nil
}

func (l *LocalNode) FindClosestPrecedingNode(_ context.Context, id model.HashID) (RingNode, error) {
	if l.isShutdown {
		return nil, ErrNodeUnavailable
	}
	for i := range l.fingerTable {
		finger := l.fingerTable[len(l.fingerTable)-(i+1)]
		// If the fingerTable has not been updated
		if finger.Node == nil {
			return nil, ErrStabilizeNotCompleted
		}
		if finger.Node.Reference().ID.Between(l.ID, id) {
			return finger.Node, nil
		}
	}
	return l, nil
}

func (l *LocalNode) Notify(_ context.Context, node RingNode) error {
	if l.isShutdown {
		return ErrNodeUnavailable
	}
	if l.predecessor == nil || node.Reference().ID.Between(l.predecessor.Reference().ID, l.ID) {
		l.predecessor = node
	}
	return nil
}

func (l *LocalNode) PutValueInner(ctx context.Context, key *string, value *string) (bool, error) {
	// TODO: need to implement LocalNode::PutValueInner
	panic("not implemented")
}

func (l *LocalNode) GetValueInner(ctx context.Context, key *string) (*string, bool, error) {
	// TODO: need to implement LocalNode::GetValueInner
	panic("not implemented")
}

func (l *LocalNode) DeleteValueInner(ctx context.Context, key *string) (bool, error) {
	// TODO: need to implement LocalNode::DeleteValueInner
	panic("not implemented")
}
