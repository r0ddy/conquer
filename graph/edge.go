package graph

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
