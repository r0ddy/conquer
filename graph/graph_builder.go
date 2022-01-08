package graph

import (
	"sort"
)

type GraphBuilder interface {
	AddNode(id NodeID, value ...interface{})
	AddEdge(from NodeID, to NodeID, value ...interface{})
	Build() (Graph, error)
}

type wrappedValue struct {
	HasValue bool
	RawValue interface{}
}

type RawGraphBuilder struct {
	builderOptions BuilderOptions
	nodes          map[NodeID]wrappedValue
	edges          map[NodeID]map[NodeID]wrappedValue
	err            error
}

func (builder *RawGraphBuilder) AddNode(id NodeID, value ...interface{}) {
	// if there is an existing error skip this command
	if builder.err != nil {
		return
	}

	// check if node exists and that duplicate nodes are not allowed
	if _, exists := builder.nodes[id]; exists && !builder.builderOptions.AllowDuplicateNodes {
		builder.err = DuplicateNodeError{nodeID: id}
		return
	}

	// check if multiple values are provided
	if len(value) > 1 {
		builder.err = MultipleValuesForNodeError{nodeID: id}
		return
	}

	wv := wrappedValue{}
	if len(value) == 1 {
		wv.HasValue = true
		wv.RawValue = value[0]
	}
	builder.nodes[id] = wv
}

func (builder *RawGraphBuilder) addEdgeHelper(from NodeID, to NodeID, value ...interface{}) {
	// check if edge exists and that duplicate edges are not allow
	edgeExists := false
	if _, existsFrom := builder.edges[from]; existsFrom {
		if _, existsTo := builder.edges[from][to]; existsTo {
			edgeExists = true
		}
	} else {
		builder.edges[from] = make(map[NodeID]wrappedValue)
	}
	if edgeExists && !builder.builderOptions.AllowDuplicateEdges {
		builder.err = DuplicateEdgeError{fromID: from, toID: to}
		return
	}

	// check if multiple values are provided
	if len(value) > 1 {
		builder.err = MultipleValuesForEdgeError{fromID: from, toID: to}
		return
	}

	// add edge with from as the first enty and to as the second entry
	wv := wrappedValue{}
	if len(value) == 1 {
		wv.HasValue = true
		wv.RawValue = value
	}
	builder.edges[from][to] = wv
}

func (builder *RawGraphBuilder) AddEdge(from NodeID, to NodeID, value ...interface{}) {
	// if there is an existing error skip this command
	if builder.err != nil {
		return
	}

	// check if both nodes exist and if build edges incrementally is enabled
	buildIncrementally := builder.builderOptions.BuildEdgesIncrementally
	if _, existsFrom := builder.nodes[from]; !existsFrom && buildIncrementally {
		builder.err = NodeNotFoundError{nodeID: from}
		return
	}
	if _, existsTo := builder.nodes[to]; !existsTo && buildIncrementally {
		builder.err = NodeNotFoundError{nodeID: to}
		return
	}

	// check that edge is not redundant
	if from == to && !builder.builderOptions.AllowRedundantEdges {
		builder.err = RedundantEdgeError{nodeID: from}
		return
	}

	builder.addEdgeHelper(from, to, value...)
}

func (builder *RawGraphBuilder) buildUndirectedGraph() (Graph, error) {
	graph := rawUndirectedGraph{
		Edges:      make([]*rawUndirectedEdge, 0),
		NodesEdges: make(map[NodeID]map[NodeID]*rawUndirectedEdge),
		Nodes:      make(map[NodeID]*rawUndirectedNode),
	}

	// map nodeIDs to nodes
	for id, val := range builder.nodes {
		graph.Nodes[id] = &rawUndirectedNode{
			ID:          id,
			Neighbors:   make([]NodeID, 0),
			RawGraphRef: &graph,
			Value:       val,
		}
	}

	for first, toVals := range builder.edges {
		for second, val := range toVals {
			// construct edge
			edge := rawUndirectedEdge{
				Nodes:       [2]NodeID{first, second},
				RawGraphRef: &graph,
				Value:       val,
			}
			// add each node to the other node's neighbor list, throw error if it does not exist
			if firstNode, firstNodeExists := graph.Nodes[first]; firstNodeExists {
				firstNode.Neighbors = append(firstNode.Neighbors, second)
			} else {
				return nil, &NodeNotFoundError{nodeID: first}
			}

			// if first and second are equal to eah other (i.e. self loop) then only add once
			if secondNode, secondNodeExists := graph.Nodes[second]; secondNodeExists {
				if first != second {
					secondNode.Neighbors = append(secondNode.Neighbors, first)
				}
			} else {
				return nil, &NodeNotFoundError{nodeID: second}
			}

			// map first, second and second, first to pointer to edge
			if _, firstExists := graph.NodesEdges[first]; !firstExists {
				graph.NodesEdges[first] = make(map[NodeID]*rawUndirectedEdge)
			}
			if _, secondExists := graph.NodesEdges[second]; !secondExists {
				graph.NodesEdges[second] = make(map[NodeID]*rawUndirectedEdge)
			}
			graph.NodesEdges[first][second] = &edge
			graph.NodesEdges[second][first] = &edge

			// add edge to list of edges to maintain list of unique edges
			graph.Edges = append(graph.Edges, &edge)
		}
	}

	// sort node.neighbors asc
	for _, node := range graph.Nodes {
		sort.Slice(node.Neighbors, func(i, j int) bool { return node.Neighbors[i] < node.Neighbors[j] })
	}
	// sort edge.nodes asc
	for _, edge := range graph.Edges {
		if edge.Nodes[0] > edge.Nodes[1] {
			edge.Nodes[0], edge.Nodes[1] = edge.Nodes[1], edge.Nodes[0]
		}
	}
	// sort graph.edges by edge.nodes
	sort.Slice(graph.Edges, func(i, j int) bool {
		if graph.Edges[i].Nodes[0] != graph.Edges[j].Nodes[0] {
			return graph.Edges[i].Nodes[0] < graph.Edges[j].Nodes[0]
		}
		return graph.Edges[i].Nodes[1] < graph.Edges[j].Nodes[1]
	})
	return graph, nil
}

func (builder *RawGraphBuilder) buildDirectedGraph() (Graph, error) {
	graph := rawDirectedGraph{
		FromToEdges: make(map[NodeID]map[NodeID]*rawDirectedEdge),
		Nodes:       make(map[NodeID]*rawDirectedNode),
	}

	// map nodeIDs to nodes
	for id, val := range builder.nodes {
		graph.Nodes[id] = &rawDirectedNode{
			ID:          id,
			Outgoing:    make([]NodeID, 0),
			Incoming:    make([]NodeID, 0),
			RawGraphRef: &graph,
			Value:       val,
		}
	}

	// map from-to to edges
	for from, toVals := range builder.edges {
		for to, val := range toVals {
			// add from to the outgoing list of to and vice versa, throw err if they don't exist
			if fromNode, fromNodeExists := graph.Nodes[from]; fromNodeExists {
				fromNode.Outgoing = append(fromNode.Outgoing, to)
			} else {
				return nil, &NodeNotFoundError{nodeID: from}
			}
			if toNode, toNodeExists := graph.Nodes[to]; toNodeExists {
				toNode.Incoming = append(toNode.Incoming, from)
			} else {
				return nil, &NodeNotFoundError{nodeID: to}
			}

			// map from-to to edge
			if _, fromExists := graph.FromToEdges[from]; !fromExists {
				graph.FromToEdges[from] = make(map[NodeID]*rawDirectedEdge)
			}
			graph.FromToEdges[from][to] = &rawDirectedEdge{
				From:        from,
				To:          to,
				RawGraphRef: &graph,
				Value:       val,
			}
		}
	}

	// sort node.outgoing and node.incoming
	for _, node := range graph.Nodes {
		sort.Slice(node.Incoming, func(i, j int) bool { return node.Incoming[i] < node.Incoming[j] })
		sort.Slice(node.Outgoing, func(i, j int) bool { return node.Outgoing[i] < node.Outgoing[j] })
	}
	return graph, nil
}

func (builder *RawGraphBuilder) Build() (Graph, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	if builder.builderOptions.IsDirected {
		return builder.buildDirectedGraph()
	}
	return builder.buildUndirectedGraph()
}

type BuilderOptions struct {
	AllowDuplicateEdges     bool
	AllowDuplicateNodes     bool
	AllowRedundantEdges     bool
	BuildEdgesIncrementally bool // if false will throw NodeNotFound when node has not been added yet
	IsDirected              bool
}

func NewGraphBuilder(bo ...BuilderOptions) GraphBuilder {
	builderOptions := BuilderOptions{}
	if len(bo) == 1 {
		builderOptions = bo[0]
	}
	return &RawGraphBuilder{
		builderOptions: builderOptions,
		nodes:          make(map[NodeID]wrappedValue),
		edges:          make(map[NodeID]map[NodeID]wrappedValue),
		err:            nil,
	}
}
