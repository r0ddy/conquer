package graph

import "fmt"

func Example() {
	// create a directed graph
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1, "node-val1")
	gb.AddNode(2, "node-val2")
	gb.AddEdge(1, 2, "edge-val")
	graph, err := gb.Build()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(graph)
	}
}
