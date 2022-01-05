package graph

type GraphBuilder interface {
	AddNode(id NodeID, value ...interface{})
	AddEdge(from NodeID, to NodeID, value ...interface{})
	Build() (Graph, error)
}

type RawGraphBuilder struct {
	builderOptions BuilderOptions
	nodes          map[NodeID]interface{}
	edges          map[NodeID]map[NodeID]interface{}
	err            error
}

func (rgb RawGraphBuilder) AddNode(id NodeID, value ...interface{}) {
	if rgb.err != nil {
		return
	}
}

func (rgb RawGraphBuilder) AddEdge(from NodeID, to NodeID, value ...interface{}) {
	if rgb.err != nil {

	}
}

func (rgb RawGraphBuilder) Build() (Graph, error) {
	return RawGraph{}, rgb.err
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
