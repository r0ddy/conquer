package graph

// Edge represents an edge in a graph with access to its node endpoints.
// If its in a directed graph, it also has access to the node the edge comes from
// and the node the edge goes into. It can also store a value.
type Edge interface {
	// GetTo returns the node the edge goes into in a directed graph.
	// In a undirected graph, this returns a "cannot use this method" error.
	GetTo() (Node, error)

	// GetFrom returns the node the edge comes from in a directed graph.
	// In a undirected graph, this returns a "cannot use this method" error.
	GetFrom() (Node, error)

	// GetNodes returns the endpoint nodes of this edge.
	// Nodes are sorted by id (ascending).
	GetNodes() ([]Node, error)

	// GetValue returns the value stored in this edge.
	// If there is no value then this returns a "no value" error.
	GetValue() (interface{}, error)
	removeRef() Edge
}

type rawDirectedEdge struct {
	From        NodeID
	To          NodeID
	RawGraphRef *rawDirectedGraph
	Value       wrappedValue
}

func (re rawDirectedEdge) GetTo() (Node, error) {
	return re.RawGraphRef.GetNode(re.To)
}

func (re rawDirectedEdge) GetFrom() (Node, error) {
	return re.RawGraphRef.GetNode(re.From)
}

func (re rawDirectedEdge) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	from, err := re.GetFrom()
	if err != nil {
		return nodes, err
	}
	to, err := re.GetTo()
	if err != nil {
		return nodes, err
	}
	nodes = append(nodes, from)
	if re.From != re.To {
		nodes = append(nodes, to)
	}
	return nodes, nil
}

func (re rawDirectedEdge) GetValue() (interface{}, error) {
	if !re.Value.HasValue {
		return nil, noValueFoundInEdgeError{fromID: re.From, toID: re.To}
	}
	return re.Value.RawValue, nil
}

func (re rawDirectedEdge) removeRef() Edge {
	re.RawGraphRef = nil
	return re
}

type rawUndirectedEdge struct {
	Nodes       [2]NodeID
	RawGraphRef *rawUndirectedGraph
	Value       wrappedValue
}

func (re rawUndirectedEdge) GetTo() (Node, error) {
	return nil, cannotUseForUndirectedGraphError{"Edge.GetTo"}
}

func (re rawUndirectedEdge) GetFrom() (Node, error) {
	return nil, cannotUseForUndirectedGraphError{"Edge.GetFrom"}
}

func (re rawUndirectedEdge) GetNodes() ([]Node, error) {
	nodes := make([]Node, 0)
	node, err := re.RawGraphRef.GetNode(re.Nodes[0])
	if err != nil {
		return nodes, err
	}
	nodes = append(nodes, node)
	node, err = re.RawGraphRef.GetNode(re.Nodes[1])
	if err != nil {
		return nodes, err
	}
	if re.Nodes[0] != re.Nodes[1] {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (re rawUndirectedEdge) GetValue() (interface{}, error) {
	if !re.Value.HasValue {
		return nil, noValueFoundInEdgeError{fromID: re.Nodes[0], toID: re.Nodes[1]}
	}
	return re.Value.RawValue, nil
}

func (re rawUndirectedEdge) removeRef() Edge {
	re.RawGraphRef = nil
	return re
}
