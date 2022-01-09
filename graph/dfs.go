package graph

type NodeFunction func(Node)

type DepthFirstSearchOptions struct {
	ReverseGraph    bool
	CustomOrder     []Node
	BeforeRecursion NodeFunction
	AfterRecursion  NodeFunction
	BeforeSearch    NodeFunction
	AfterSearch     NodeFunction
	IsUndirected    bool
}

func depthFirstSearchHelper(node Node, visited map[NodeID]bool, options DepthFirstSearchOptions) {
	if !visited[node.GetID()] {
		visited[node.GetID()] = true
		edges, _ := node.GetOutgoingEdges()
		if options.ReverseGraph {
			edges, _ = node.GetIncomingEdges()
		}
		if options.IsUndirected {
			edges, _ = node.GetIncidentEdges()
		}
		if options.BeforeRecursion != nil {
			options.BeforeRecursion(node)
		}
		for _, edge := range edges {
			other_node, _ := edge.GetTo()
			if options.ReverseGraph {
				other_node, _ = edge.GetFrom()
			}
			if options.IsUndirected {
				nodes, _ := edge.GetNodes()
				if len(nodes) == 1 {
					other_node = nodes[0]
				} else if nodes[0].GetID() != node.GetID() {
					other_node = nodes[0]
				} else {
					other_node = nodes[1]
				}
			}
			depthFirstSearchHelper(other_node, visited, options)
		}
		if options.AfterRecursion != nil {
			options.AfterRecursion(node)
		}
	}
}

func DepthFirstSearch(g Graph, options DepthFirstSearchOptions) {
	if !g.IsDirected() {
		options.IsUndirected = true
	}
	visited := make(map[NodeID]bool)
	allNodes, _ := g.GetNodes()
	for _, node := range allNodes {
		visited[node.GetID()] = false
	}
	if options.CustomOrder != nil {
		allNodes = options.CustomOrder
	}
	for _, node := range allNodes {
		if options.BeforeSearch != nil {
			options.BeforeSearch(node)
		}
		depthFirstSearchHelper(node, visited, options)
		if options.AfterSearch != nil {
			options.AfterSearch(node)
		}
	}
}
