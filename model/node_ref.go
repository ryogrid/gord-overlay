package model

type NodeRef struct {
	ID   HashID
	Host string
}

func NewNodeRef(host string) *NodeRef {
	return &NodeRef{
		ID:   NewHashID(host),
		Host: host,
	}
}
