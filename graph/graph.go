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

type EdgeID int

type Edge interface {
	GetID() EdgeID
	GetTo() Node
	GetFrom() Node
	GetNodes() [2]Node
	GetValue() (interface{}, error)
}

type RawEdge struct {
	ID          EdgeID
	To          NodeID
	From        NodeID
	RawGraphRef *RawGraph
	Value       interface{}
}

func (re RawEdge) GetID() EdgeID {
	return re.ID
}

func (re RawEdge) GetTo() Node {
	return RawNode{}
}

func (re RawEdge) GetFrom() Node {
	return RawNode{}
}

func (re RawEdge) GetNodes() Node {
	return RawNode{}
}

func (re RawEdge) GetValue() (interface{}, error) {
	return nil, nil
}

type Graph interface {
	GetNode(id NodeID) (Node, error)
	GetNodes() ([]Node, error)
	GetEdges() ([]Edge, error)
	GetRaw() RawGraph
}

type RawGraph struct {
}

func (rg RawGraph) GetNode(id NodeID) (Node, error) {
	return nil, nil
}

func (rg RawGraph) GetNodes() ([]Node, error) {
	return []Node{}, nil
}

func (rg RawGraph) GetEdges() ([]Edge, error) {
	return []Edge{}, nil
}

func (rg RawGraph) GetRaw() RawGraph {
	return rg
}

type GraphBuilder interface {
	AddNode(id NodeID, value ...interface{})
	AddEdge(from NodeID, to NodeID, value ...interface{})
	Build() (Graph, error)
}

type RawGraphBuilder struct {
}

func (rgb RawGraphBuilder) AddNode(id NodeID, value ...interface{}) {

}

func (rgb RawGraphBuilder) AddEdge(from NodeID, to NodeID, value ...interface{}) {

}

func (rgb RawGraphBuilder) Build() (Graph, error) {
	return RawGraph{}, nil
}

type BuilderOptions struct {
	AllowDuplicateEdges bool
	AllowDuplicateNodes bool
	AllowRedundantEdges bool
	IsDirected          bool
}

func NewGraphBuilder(bo ...BuilderOptions) GraphBuilder {
	return RawGraphBuilder{}
}
