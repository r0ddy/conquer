package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertGraphEquals(t *testing.T, expected, actual Graph) bool {
	refLessExpected := expected.removeRefs()
	refLessActual := actual.removeRefs()
	return assert.True(t,
		cmp.Equal(refLessExpected, refLessActual),
		cmp.Diff(refLessExpected, refLessActual),
	)
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

func Test_UndirectedCycle(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(2, 3)
	gb.AddEdge(3, 1)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)

	expected_graph := rawUndirectedGraph{
		Edges: []*rawUndirectedEdge{
			{Nodes: [2]NodeID{1, 2}},
			{Nodes: [2]NodeID{1, 3}},
			{Nodes: [2]NodeID{2, 3}},
		},
		Nodes: map[NodeID]*rawUndirectedNode{
			1: {ID: 1, Neighbors: []NodeID{2, 3}},
			2: {ID: 2, Neighbors: []NodeID{1, 3}},
			3: {ID: 3, Neighbors: []NodeID{1, 2}},
		},
		NodesEdges: map[NodeID]map[NodeID]*rawUndirectedEdge{
			1: {
				2: {Nodes: [2]NodeID{1, 2}},
				3: {Nodes: [2]NodeID{1, 3}},
			},
			2: {
				1: {Nodes: [2]NodeID{1, 2}},
				3: {Nodes: [2]NodeID{2, 3}},
			},
			3: {
				1: {Nodes: [2]NodeID{1, 3}},
				2: {Nodes: [2]NodeID{2, 3}},
			},
		},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}

func Test_UndirectedForest(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddNode(4)
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 4)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)
	expected_graph := rawUndirectedGraph{
		Edges: []*rawUndirectedEdge{
			{Nodes: [2]NodeID{1, 2}},
			{Nodes: [2]NodeID{3, 4}},
		},
		Nodes: map[NodeID]*rawUndirectedNode{
			1: {ID: 1, Neighbors: []NodeID{2}},
			2: {ID: 2, Neighbors: []NodeID{1}},
			3: {ID: 3, Neighbors: []NodeID{4}},
			4: {ID: 4, Neighbors: []NodeID{3}},
		},
		NodesEdges: map[NodeID]map[NodeID]*rawUndirectedEdge{
			1: {
				2: {Nodes: [2]NodeID{1, 2}},
			},
			2: {
				1: {Nodes: [2]NodeID{1, 2}},
			},
			3: {
				4: {Nodes: [2]NodeID{3, 4}},
			},
			4: {
				3: {Nodes: [2]NodeID{3, 4}},
			},
		},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}

func Test_DirectedTree(t *testing.T) {
	graphBuilder := NewGraphBuilder(BuilderOptions{IsDirected: true})
	graphBuilder.AddNode(1)
	graphBuilder.AddNode(2)
	graphBuilder.AddNode(3)
	graphBuilder.AddNode(4)
	graphBuilder.AddEdge(1, 2)
	graphBuilder.AddEdge(1, 3)
	graphBuilder.AddEdge(1, 4)
	graph, err := graphBuilder.Build()

	assert.NoError(t, err)
	expected_graph := rawDirectedGraph{
		Nodes: map[NodeID]*rawDirectedNode{
			1: {ID: 1, Outgoing: []NodeID{2, 3, 4}, Incoming: []NodeID{}},
			2: {ID: 2, Outgoing: []NodeID{}, Incoming: []NodeID{1}},
			3: {ID: 3, Outgoing: []NodeID{}, Incoming: []NodeID{1}},
			4: {ID: 4, Outgoing: []NodeID{}, Incoming: []NodeID{1}},
		},
		FromToEdges: map[NodeID]map[NodeID]*rawDirectedEdge{
			1: {
				2: {From: 1, To: 2},
				3: {From: 1, To: 3},
				4: {From: 1, To: 4},
			},
		},
	}
	AssertGraphEquals(t, expected_graph, graph)
}

func Test_DirectedCycle(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(2, 3)
	gb.AddEdge(3, 1)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)

	expected_graph := rawDirectedGraph{
		Nodes: map[NodeID]*rawDirectedNode{
			1: {ID: 1, Outgoing: []NodeID{2}, Incoming: []NodeID{3}},
			2: {ID: 2, Outgoing: []NodeID{3}, Incoming: []NodeID{1}},
			3: {ID: 3, Outgoing: []NodeID{1}, Incoming: []NodeID{2}},
		},
		FromToEdges: map[NodeID]map[NodeID]*rawDirectedEdge{
			1: {
				2: {From: 1, To: 2},
			},
			2: {
				3: {From: 2, To: 3},
			},
			3: {
				1: {From: 3, To: 1},
			},
		},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}

func Test_DirectedForest(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddNode(4)
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 4)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)

	expected_graph := rawDirectedGraph{
		Nodes: map[NodeID]*rawDirectedNode{
			1: {ID: 1, Outgoing: []NodeID{2}, Incoming: []NodeID{}},
			2: {ID: 2, Outgoing: []NodeID{}, Incoming: []NodeID{1}},
			3: {ID: 3, Outgoing: []NodeID{4}, Incoming: []NodeID{}},
			4: {ID: 4, Outgoing: []NodeID{}, Incoming: []NodeID{3}},
		},
		FromToEdges: map[NodeID]map[NodeID]*rawDirectedEdge{
			1: {
				2: {From: 1, To: 2},
			},
			3: {
				4: {From: 3, To: 4},
			},
		},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}

func Test_DuplicateNode(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{AllowDuplicateNodes: false})
	gb.AddNode(1)
	gb.AddNode(1)
	_, err := gb.Build()
	assert.ErrorIs(t, err, DuplicateNodeError{nodeID: 1})

	gb = NewGraphBuilder(BuilderOptions{AllowDuplicateNodes: true})
	gb.AddNode(1)
	gb.AddNode(1)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)
	expected_graph := rawUndirectedGraph{
		Nodes: map[NodeID]*rawUndirectedNode{
			1: {ID: 1, Neighbors: []NodeID{}},
		},
		Edges:      []*rawUndirectedEdge{},
		NodesEdges: map[NodeID]map[NodeID]*rawUndirectedEdge{},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}

func Test_DuplicateEdge(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{AllowDuplicateEdges: false})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 2)
	_, err := gb.Build()
	assert.ErrorIs(t, err, DuplicateEdgeError{fromID: 1, toID: 2})

	gb = NewGraphBuilder(BuilderOptions{AllowDuplicateEdges: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 2)
	actual_graph, err := gb.Build()
	assert.NoError(t, err)
	expected_graph := rawUndirectedGraph{
		Nodes: map[NodeID]*rawUndirectedNode{
			1: {
				ID: 1, Neighbors: []NodeID{2},
			},
			2: {
				ID: 2, Neighbors: []NodeID{1},
			},
		},
		NodesEdges: map[NodeID]map[NodeID]*rawUndirectedEdge{
			1: {
				2: &rawUndirectedEdge{Nodes: [2]NodeID{1, 2}},
			},
			2: {
				1: &rawUndirectedEdge{Nodes: [2]NodeID{1, 2}},
			},
		},
		Edges: []*rawUndirectedEdge{
			{Nodes: [2]NodeID{1, 2}},
		},
	}
	AssertGraphEquals(t, expected_graph, actual_graph)
}
