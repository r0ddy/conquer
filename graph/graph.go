package graph

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
