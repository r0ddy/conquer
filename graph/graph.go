package graph

import "sort"

type Graph interface {
	GetNode(id NodeID) (Node, error)
	GetEdge(from NodeID, to NodeID) (Edge, error)
	GetNodes() ([]Node, error)
	GetEdges() ([]Edge, error)
	Serialize() GraphSerializer
}

type rawDirectedGraph struct {
	FromToEdges map[NodeID]map[NodeID]*rawDirectedEdge
	Nodes       map[NodeID]*rawDirectedNode
}

func (rg rawDirectedGraph) GetNode(id NodeID) (Node, error) {
	node, exists := rg.Nodes[id]
	if !exists {
		return nil, NodeNotFoundError{nodeID: id}
	}
	return node, nil
}

func (rg rawDirectedGraph) GetEdge(from NodeID, to NodeID) (Edge, error) {
	if _, fromExists := rg.FromToEdges[from]; fromExists {
		if edge, toExists := rg.FromToEdges[from][to]; toExists {
			return edge, nil
		}
	}
	return nil, EdgeNotFoundError{fromID: from, toID: to}
}

func (rg rawDirectedGraph) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	for _, node := range rg.Nodes {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (rg rawDirectedGraph) GetEdges() ([]Edge, error) {
	edges := make([]Edge, 0)
	for _, subEdges := range rg.FromToEdges {
		for _, edge := range subEdges {
			edges = append(edges, edge)
		}
	}
	return edges, nil
}

func (rg rawDirectedGraph) Serialize() GraphSerializer {
	gs := GraphSerializer{
		Edges:       make([]EdgeSerializer, 0),
		Nodes:       make([]NodeSerializer, 0),
		NodesToEdge: make(map[NodeID]map[NodeID]EdgeSerializer),
	}
	for _, edges := range rg.FromToEdges {
		for _, edge := range edges {
			gs.Edges = append(gs.Edges, edge.Serialize())
		}
	}
	for _, node := range rg.Nodes {
		gs.Nodes = append(gs.Nodes, node.Serialize())
	}
	sort.Slice(gs.Edges, func(i, j int) bool {
		if gs.Edges[i].From != gs.Edges[j].From {
			return gs.Edges[i].From < gs.Edges[j].To
		}
		return gs.Edges[i].To < gs.Edges[j].To
	})

	for from, toEdges := range rg.FromToEdges {
		for to, edge := range toEdges {
			if _, exists := gs.NodesToEdge[from]; !exists {
				gs.NodesToEdge[to] = make(map[NodeID]EdgeSerializer)
			}
			gs.NodesToEdge[from][to] = edge.Serialize()
		}
	}

	return gs
}

type rawUndirectedGraph struct {
	Edges      []*rawUndirectedEdge
	NodesEdges map[NodeID]map[NodeID]*rawUndirectedEdge
	Nodes      map[NodeID]*rawUndirectedNode
}

func (rg rawUndirectedGraph) GetNode(id NodeID) (Node, error) {
	node, exists := rg.Nodes[id]
	if !exists {
		return nil, NodeNotFoundError{nodeID: id}
	}
	return node, nil
}

func (rg rawUndirectedGraph) GetEdge(first NodeID, second NodeID) (Edge, error) {
	if _, firstExists := rg.NodesEdges[first]; firstExists {
		if edge, secondExists := rg.NodesEdges[first][second]; secondExists {
			return edge, nil
		}
	}
	return nil, EdgeNotFoundError{fromID: first, toID: second}
}

func (rg rawUndirectedGraph) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	for _, node := range rg.Nodes {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (rg rawUndirectedGraph) GetEdges() ([]Edge, error) {
	edges := make([]Edge, 0)
	for _, edge := range rg.Edges {
		edges = append(edges, *edge)
	}
	return edges, nil
}

func (rg rawUndirectedGraph) Serialize() GraphSerializer {
	gs := GraphSerializer{
		Edges:       make([]EdgeSerializer, 0),
		Nodes:       make([]NodeSerializer, 0),
		NodesToEdge: make(map[NodeID]map[NodeID]EdgeSerializer),
	}
	for _, edge := range rg.Edges {
		gs.Edges = append(gs.Edges, edge.Serialize())
	}
	for _, node := range rg.Nodes {
		gs.Nodes = append(gs.Nodes, node.Serialize())
	}
	sort.Slice(gs.Nodes, func(i, j int) bool { return gs.Nodes[i].ID < gs.Nodes[j].ID })

	for first, secondEdges := range rg.NodesEdges {
		for second, edge := range secondEdges {
			if _, exists := gs.NodesToEdge[first]; !exists {
				gs.NodesToEdge[first] = make(map[NodeID]EdgeSerializer)
			}
			gs.NodesToEdge[first][second] = edge.Serialize()
		}
	}

	return gs
}

type GraphSerializer struct {
	Edges       []EdgeSerializer                     `json:"edges,omitempty"`
	Nodes       []NodeSerializer                     `json:"nodes,omitempty"`
	NodesToEdge map[NodeID]map[NodeID]EdgeSerializer `json:"nodesToEdge,omitempty"`
}
