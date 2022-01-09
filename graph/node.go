package graph

type NodeID int

// Node represents a node in a graph with access to its
// incident edges. If its in a directed graph, it also has access
// to its incoming/outgoing edges. It can also store a value.
type Node interface {
	// GetID returns the node's unique identifier.
	GetID() NodeID

	// GetIncomingEdges returns the edges that are pointing to this node in a directed graph.
	// The edges are sorted by NodeID on the other side of the incoming edge (ascending).
	// In a undirected graph, this returns a "cannot use this method" error.
	GetIncomingEdges() ([]Edge, error)

	// GetOutgoingEdges returns the edges that are stemming from this node in a directed graph.
	// The edges are sorted by NodeID on the other side of the outgoing edge (ascending).
	// In a undirected graph, this returns a "cannot use this method" error
	GetOutgoingEdges() ([]Edge, error)

	// GetIncidentEdges returns all the edges that this node is an endpoint of (directed or undirected).
	// If the edges are from an undirected graph, the nodes in each edge will be sorted by id (ascending).
	// Then the edges are sorted by the first entry and the second entry in this NodeID slice (ascending).
	// If the edges are from a directed graph, the incoming edges are first then the outgoing edges.
	GetIncidentEdges() ([]Edge, error)

	// GetValue return the value stored in this node.
	// If there is no value then this returns a "no value" error.
	GetValue() (interface{}, error)
	removeRef() Node
}

type rawDirectedNode struct {
	ID          NodeID
	Incoming    []NodeID
	Outgoing    []NodeID
	RawGraphRef *rawDirectedGraph
	Value       wrappedValue
}

func (rn rawDirectedNode) GetID() NodeID {
	return rn.ID
}

func (rn rawDirectedNode) GetIncomingEdges() ([]Edge, error) {
	incoming := make([]Edge, 0)
	for _, fromID := range rn.Incoming {
		edge, err := rn.RawGraphRef.GetEdge(fromID, rn.ID)
		if err != nil {
			return nil, err
		}
		incoming = append(incoming, edge)
	}
	return incoming, nil
}

func (rn rawDirectedNode) GetOutgoingEdges() ([]Edge, error) {
	outgoing := make([]Edge, 0)
	for _, toID := range rn.Outgoing {
		edge, err := rn.RawGraphRef.GetEdge(rn.ID, toID)
		if err != nil {
			return nil, err
		}
		outgoing = append(outgoing, edge)
	}
	return outgoing, nil
}

func (rn rawDirectedNode) GetIncidentEdges() ([]Edge, error) {
	incoming, err := rn.GetIncomingEdges()
	if err != nil {
		return nil, err
	}
	outgoing, err := rn.GetOutgoingEdges()
	if err != nil {
		return nil, err
	}
	incident := append(incoming, outgoing...)
	return incident, nil
}

func (rn rawDirectedNode) GetValue() (interface{}, error) {
	if !rn.Value.HasValue {
		return nil, noValueFoundInNodeError{rn.ID}
	}
	return rn.Value.RawValue, nil
}

func (rn rawDirectedNode) removeRef() Node {
	rn.RawGraphRef = nil
	return rn
}

type rawUndirectedNode struct {
	ID          NodeID
	Neighbors   []NodeID
	RawGraphRef *rawUndirectedGraph
	Value       wrappedValue
}

func (rn rawUndirectedNode) GetID() NodeID {
	return rn.ID
}

func (rn rawUndirectedNode) GetIncomingEdges() ([]Edge, error) {
	return nil, cannotUseForUndirectedGraphError{"Node.GetIncomingEdges"}
}

func (rn rawUndirectedNode) GetOutgoingEdges() ([]Edge, error) {
	return nil, cannotUseForUndirectedGraphError{"Node.GetOutgoingEdges"}
}

func (rn rawUndirectedNode) GetIncidentEdges() ([]Edge, error) {
	edges := make([]Edge, 0)
	for _, nodeID := range rn.Neighbors {
		edge, err := rn.RawGraphRef.GetEdge(rn.ID, nodeID)
		if err != nil {
			return nil, err
		}
		edges = append(edges, edge)
	}
	return edges, nil
}

func (rn rawUndirectedNode) GetValue() (interface{}, error) {
	if !rn.Value.HasValue {
		return nil, noValueFoundInNodeError{rn.ID}
	}
	return rn.Value.RawValue, nil
}

func (rn rawUndirectedNode) removeRef() Node {
	rn.RawGraphRef = nil
	return rn
}
