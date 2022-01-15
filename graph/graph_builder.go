package graph

import (
	"sort"
)

// GraphBuilder uses the builder pattern to create directed or undirected graphs.
type GraphBuilder interface {
	// AddNode adds a node to the current graph current being built.
	// The id parameter is the unique id used to identify this node.
	// The value parameter can optionally be used to store a value in this node.
	// Additionally, AddEdge will connect nodes added via their ids.
	AddNode(id NodeID, value ...interface{})

	// AddEdge adds an edge connecting two nodes.
	// In a directed graph, it uses the fromID parameter and toID parameter to connect an edge from the former to the latter.
	// In an undirected graph, it will create a undirected edge between the two.
	// The value parameter can optionally be used to store a value in this edge.
	AddEdge(from NodeID, to NodeID, value ...interface{})

	// Build creates a directed/indirect graph using the ndoes and edges created above.
	// Returns an error if any of the aforementioned errors is detected.
	Build() (Graph, error)
}

type wrappedValue struct {
	HasValue bool
	RawValue interface{}
}

type rawGraphBuilder struct {
	builderOptions BuilderOptions
	nodes          map[NodeID]wrappedValue
	edges          map[NodeID]map[NodeID]wrappedValue
	err            error
}

func (builder *rawGraphBuilder) AddNode(id NodeID, value ...interface{}) {
	// if there is an existing error skip this command
	if builder.err != nil {
		return
	}

	// check if node exists and that duplicate nodes are not allowed
	if _, exists := builder.nodes[id]; exists && !builder.builderOptions.AllowDuplicateNodes {
		builder.err = duplicateNodeError{nodeID: id}
		return
	}

	// check if multiple values are provided
	if len(value) > 1 {
		builder.err = multipleValuesForNodeError{nodeID: id}
		return
	}

	wv := wrappedValue{}
	if len(value) == 1 {
		wv.HasValue = true
		wv.RawValue = value[0]
	}
	builder.nodes[id] = wv
}

func (builder *rawGraphBuilder) addEdgeHelper(from NodeID, to NodeID, value ...interface{}) {
	// ensures that addEdgeHelper(8, 9) and addEdgeHelper(9, 8) only add one edge
	if !builder.builderOptions.IsDirected && from > to {
		from, to = to, from
	}
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
		builder.err = duplicateEdgeError{fromID: from, toID: to}
		return
	}

	// check if multiple values are provided
	if len(value) > 1 {
		builder.err = multipleValuesForEdgeError{fromID: from, toID: to}
		return
	}

	// add edge with from as the first enty and to as the second entry
	wv := wrappedValue{}
	if len(value) == 1 {
		wv.HasValue = true
		wv.RawValue = value[0]
	}
	builder.edges[from][to] = wv
}

func (builder *rawGraphBuilder) AddEdge(fromID NodeID, toID NodeID, value ...interface{}) {
	// if there is an existing error skip this command
	if builder.err != nil {
		return
	}

	// check if both nodes exist and if build edges incrementally is enabled
	buildIncrementally := builder.builderOptions.BuildEdgesIncrementally
	if _, existsFrom := builder.nodes[fromID]; !existsFrom && buildIncrementally {
		builder.err = nodeNotFoundError{nodeID: fromID}
		return
	}
	if _, existsTo := builder.nodes[toID]; !existsTo && buildIncrementally {
		builder.err = nodeNotFoundError{nodeID: toID}
		return
	}

	// check that edge is not redundant
	if fromID == toID && !builder.builderOptions.AllowRedundantEdges {
		builder.err = redundantEdgeError{nodeID: fromID}
		return
	}

	builder.addEdgeHelper(fromID, toID, value...)
}

func (builder *rawGraphBuilder) buildUndirectedGraph() (Graph, error) {
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
				return nil, &nodeNotFoundError{nodeID: first}
			}

			// if first and second are equal to eah other (i.e. self loop) then only add once
			if secondNode, secondNodeExists := graph.Nodes[second]; secondNodeExists {
				if first != second {
					secondNode.Neighbors = append(secondNode.Neighbors, first)
				}
			} else {
				return nil, &nodeNotFoundError{nodeID: second}
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
		sortNodeIDs(node.Neighbors)
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

func (builder *rawGraphBuilder) buildDirectedGraph() (Graph, error) {
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
				return nil, &nodeNotFoundError{nodeID: from}
			}
			if toNode, toNodeExists := graph.Nodes[to]; toNodeExists {
				toNode.Incoming = append(toNode.Incoming, from)
			} else {
				return nil, &nodeNotFoundError{nodeID: to}
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
		sortNodeIDs(node.Incoming)
		sortNodeIDs(node.Outgoing)
	}
	return graph, nil
}

func (builder *rawGraphBuilder) Build() (Graph, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	if builder.builderOptions.IsDirected {
		return builder.buildDirectedGraph()
	}
	return builder.buildUndirectedGraph()
}

// BuilderOptions determine what is or isn't allowed in during the construction of a graph.
// It also specifies whether the graph is directed or undirected.
type BuilderOptions struct {
	// AllowDuplicateEdges will allow AddEdge(a, b, val) to be executed multiple times if set to true.
	// The value of edge a-b can change with each call, but the graph will use the last one.
	// If false, Build will return a duplicate edge a-b error.
	AllowDuplicateEdges bool
	// AllowDuplicateNodes will allow AddNode(id, val) to be executed multiple times if set to true.
	// The value of the node id can change with each call, but the graph will use the last one.
	// If false, Build will return a duplicate node id error.
	AllowDuplicateNodes bool
	// AllowRedundantEdges will allow AddEdge(a, a, val) to create a self-loop on node a if set to true.
	// If false, Build will return a redundant edge a-a error.
	AllowRedundantEdges bool
	// BuildEdgesIncrementally will allow creation of a graph where AddEdge(a, b, val) occurs before AddNode(a) and/or AddNode(b) if set to true.
	// If false, Build will return a node a/b not found error
	BuildEdgesIncrementally bool
	// IsDirected tells the builder to create a directed graph if set to true.
	// If false, builder will assume undirected graph.
	IsDirected bool
}

func NewGraphBuilder(bo ...BuilderOptions) GraphBuilder {
	builderOptions := BuilderOptions{}
	if len(bo) == 1 {
		builderOptions = bo[0]
	}
	return &rawGraphBuilder{
		builderOptions: builderOptions,
		nodes:          make(map[NodeID]wrappedValue),
		edges:          make(map[NodeID]map[NodeID]wrappedValue),
		err:            nil,
	}
}
