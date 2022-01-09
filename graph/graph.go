package graph

import "sort"

// Graph represents a directed/undirected graph.
type Graph interface {
	// GetNode fetches a node by its id.
	// If the id does not exist in the graph, it returns a node not found error.
	GetNode(id NodeID) (Node, error)

	// GetEdge fetches an edge by its two endpoint.
	// In a directed graph, the from parameter takes the id of where the edge starts
	// while the to parameter takes the id of where the edge ends.
	// In an undirected graph, GetEdge(a, b) is equivalent to GetEdge(b, a).
	// If edge from-to does not exist in the graph, it returns an edge not found error.
	GetEdge(from NodeID, to NodeID) (Edge, error)

	// GetNodes fetches all the unique nodes of this graph sorted by id (asecnding).
	GetNodes() ([]Node, error)

	// GetEdges fetches all the unique edges of this graph.
	// In a directed graph, the edges are sorted by from NodeID then to NodeID (ascending).
	// In a undirected graph, the nodes in an edge are sorted by id (ascending) and
	// then the edges are sorted by the first entry in that node slice (also ascending).
	// If those are equal, then they're sorted by the next entry.
	GetEdges() ([]Edge, error)

	// IsDirected returns true if the graph is directed and false if its undirected.
	IsDirected() bool
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
	sortNodes(nodes)
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

func (g rawDirectedGraph) IsDirected() bool {
	return true
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
	sortNodes(nodes)
	return nodes, nil
}

func (rg rawUndirectedGraph) GetEdges() ([]Edge, error) {
	edges := make([]Edge, 0)
	for _, edge := range rg.Edges {
		edges = append(edges, *edge)
	}
	return edges, nil
}

func (g rawUndirectedGraph) IsDirected() bool {
	return false
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
