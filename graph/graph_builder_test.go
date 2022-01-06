package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertGraphEquals(t *testing.T, expected, actual Graph) bool {
	serialize := cmp.Transformer("serialize_graph", func(g Graph) GraphSerializer {
		return g.Serialize()
	})
	return assert.True(t, cmp.Equal(expected, actual, serialize), cmp.Diff(expected, actual, serialize))
}

func Test_UndirectedTree(t *testing.T) {
	// build undirected graph
	graphBuilder := NewGraphBuilder()
	graphBuilder.AddNode(1)
	graphBuilder.AddNode(2)
	graphBuilder.AddNode(3)
	graphBuilder.AddNode(4)
	graphBuilder.AddEdge(1, 2)
	graphBuilder.AddEdge(3, 1)
	graphBuilder.AddEdge(4, 1)
	graph, err := graphBuilder.Build()

	// assert graph is correct
	assert.NoError(t, err)
	expected_graph := rawUndirectedGraph{
		Nodes: map[NodeID]*rawUndirectedNode{
			1: {ID: 1, Neighbors: []NodeID{2, 3, 4}},
			2: {ID: 2, Neighbors: []NodeID{1}},
			3: {ID: 3, Neighbors: []NodeID{1}},
			4: {ID: 4, Neighbors: []NodeID{1}},
		},
		Edges: []*rawUndirectedEdge{
			{Nodes: [2]NodeID{1, 2}},
			{Nodes: [2]NodeID{1, 3}},
			{Nodes: [2]NodeID{1, 4}},
		},
		NodesEdges: map[NodeID]map[NodeID]*rawUndirectedEdge{
			1: {
				2: {Nodes: [2]NodeID{1, 2}},
				3: {Nodes: [2]NodeID{1, 3}},
				4: {Nodes: [2]NodeID{1, 4}},
			},
			2: {1: {Nodes: [2]NodeID{1, 2}}},
			3: {1: {Nodes: [2]NodeID{1, 3}}},
			4: {1: {Nodes: [2]NodeID{1, 4}}},
		},
	}
	AssertGraphEquals(t, expected_graph, graph)
}

// func Test_UndirectedCycle(t *testing.T) {
// 	gb := NewGraphBuilder()
// 	gb.AddNode(1)
// 	gb.AddNode(2)
// 	gb.AddNode(3)
// 	gb.AddEdge(1, 2)
// 	gb.AddEdge(2, 3)
// 	gb.AddEdge(3, 1)
// 	actual_graph, err := gb.Build()
// 	assert.NoError(t, err)

// 	expected_graph := rawGraph{
// 		Nodes: map[NodeID]*rawNode{
// 			1: {ID: 1, Outgoing: []NodeID{2, 3}, Incoming: []NodeID{2, 3}},
// 			2: {ID: 2, Outgoing: []NodeID{1, 3}, Incoming: []NodeID{1, 3}},
// 			3: {ID: 3, Outgoing: []NodeID{1, 2}, Incoming: []NodeID{1, 2}},
// 		},
// 		FromToEdges: map[NodeID]map[NodeID]*rawEdge{
// 			1: {
// 				2: {From: 1, To: 2},
// 				3: {From: 1, To: 3},
// 			},
// 			2: {
// 				1: {From: 2, To: 1},
// 				3: {From: 2, To: 3},
// 			},
// 			3: {
// 				1: {From: 3, To: 1},
// 				2: {From: 3, To: 2},
// 			},
// 		},
// 	}
// 	AssertGraphEquals(t, expected_graph, actual_graph)
// }

func Test_UndirectedForest(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddNode(4)
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 4)
	// actual_graph, err := gb.Build()
	// assert.NoError(t, err)
}
