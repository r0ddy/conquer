package graph

import "sort"

type Graph interface {
	GetNode(id NodeID) (Node, error)
	GetEdge(from NodeID, to NodeID) (Edge, error)
	GetNodes() ([]Node, error)
	GetEdges() ([]Edge, error)
	removeRefs() Graph
}

type rawDirectedGraph struct {
	FromToEdges map[NodeID]map[NodeID]*rawDirectedEdge
	Nodes       map[NodeID]*rawDirectedNode
}

func (rg rawDirectedGraph) GetNode(id NodeID) (Node, error) {
	node, exists := rg.Nodes[id]
	if !exists {
		return nil, nodeNotFoundError{nodeID: id}
	}
	return node, nil
}

func (rg rawDirectedGraph) GetEdge(from NodeID, to NodeID) (Edge, error) {
	if _, fromExists := rg.FromToEdges[from]; fromExists {
		if edge, toExists := rg.FromToEdges[from][to]; toExists {
			return edge, nil
		}
	}
	return nil, edgeNotFoundError{fromID: from, toID: to}
}

func (rg rawDirectedGraph) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	for _, node := range rg.Nodes {
		nodes = append(nodes, node)
	}
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].GetID() < nodes[j].GetID() })
	return nodes, nil
}

func (rg rawDirectedGraph) GetEdges() ([]Edge, error) {
	directedEdges := make([]*rawDirectedEdge, 0)
	for _, subEdges := range rg.FromToEdges {
		for _, edge := range subEdges {
			directedEdges = append(directedEdges, edge)
		}
	}
	sort.Slice(directedEdges, func(i, j int) bool {
		if directedEdges[i].From != directedEdges[j].From {
			return directedEdges[i].From < directedEdges[j].From
		}
		return directedEdges[i].To < directedEdges[j].To
	})
	edges := make([]Edge, 0)
	for _, directedEdge := range directedEdges {
		edges = append(edges, directedEdge)
	}
	return edges, nil
}

func (rg rawDirectedGraph) removeRefs() Graph {
	for _, edges := range rg.FromToEdges {
		for _, edge := range edges {
			edge.RawGraphRef = nil
		}
	}
	for _, node := range rg.Nodes {
		node.RawGraphRef = nil
	}
	return rg
}

type rawUndirectedGraph struct {
	Edges      []*rawUndirectedEdge
	NodesEdges map[NodeID]map[NodeID]*rawUndirectedEdge
	Nodes      map[NodeID]*rawUndirectedNode
}

func (rg rawUndirectedGraph) GetNode(id NodeID) (Node, error) {
	node, exists := rg.Nodes[id]
	if !exists {
		return nil, nodeNotFoundError{nodeID: id}
	}
	return node, nil
}

func (rg rawUndirectedGraph) GetEdge(first NodeID, second NodeID) (Edge, error) {
	if _, firstExists := rg.NodesEdges[first]; firstExists {
		if edge, secondExists := rg.NodesEdges[first][second]; secondExists {
			return edge, nil
		}
	}
	return nil, edgeNotFoundError{fromID: first, toID: second}
}

func (rg rawUndirectedGraph) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	for _, node := range rg.Nodes {
		nodes = append(nodes, node)
	}
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].GetID() < nodes[j].GetID() })
	return nodes, nil
}

func (rg rawUndirectedGraph) GetEdges() ([]Edge, error) {
	edges := make([]Edge, 0)
	for _, edge := range rg.Edges {
		edges = append(edges, *edge)
	}
	return edges, nil
}

func (rg rawUndirectedGraph) removeRefs() Graph {
	for _, edge := range rg.Edges {
		edge.RawGraphRef = nil
	}
	for _, node := range rg.Nodes {
		node.RawGraphRef = nil
	}
	return rg
}

type GraphSerializer struct {
	Edges       []EdgeSerializer                     `json:"edges,omitempty"`
	Nodes       []NodeSerializer                     `json:"nodes,omitempty"`
	NodesToEdge map[NodeID]map[NodeID]EdgeSerializer `json:"nodesToEdge,omitempty"`
}
