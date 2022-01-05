package graph

import "fmt"

type NodeID int

type Node interface {
	GetID() NodeID
	GetIncomingEdges() []Edge
	GetOutgoingEdges() []Edge
	GetValue() (interface{}, error)
}

type RawNode struct {
	ID          NodeID
	Neighbors   []EdgeID
	RawGraphRef *RawGraph
	Value       interface{}
}

func (rn RawNode) GetID() NodeID {
	return rn.ID
}

func (rn RawNode) GetIncomingEdges() []Edge {
	return []Edge{}
}

func (rn RawNode) GetOutgoingEdges() []Edge {
	return []Edge{}
}

func (rn RawNode) GetValue() (interface{}, error) {
	return nil, fmt.Errorf("unimplemented")
}
