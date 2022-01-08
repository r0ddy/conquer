package graph

type NodeID int

type Node interface {
	GetID() NodeID
	GetIncomingEdges() ([]Edge, error)
	GetOutgoingEdges() ([]Edge, error)
	GetIncidentEdges() ([]Edge, error)
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

type NodeSerializer struct {
	ID        NodeID      `json:"id"`
	Outgoing  []NodeID    `json:"outgoing,omitempty"`
	Incoming  []NodeID    `json:"incoming,omitempty"`
	Neighbors []NodeID    `json:"neighbors,omitempty"`
	Value     interface{} `json:"value,omitempty"`
}
