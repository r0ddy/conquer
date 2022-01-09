package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertNodeEquals(t *testing.T, expected, actual Node) bool {
	refLessExpected := expected.removeRef()
	refLessActual := actual.removeRef()
	return assert.True(t,
		cmp.Equal(refLessExpected, refLessActual),
		cmp.Diff(refLessExpected, refLessActual),
	)
}

func Test_UndirectedGraphGetNode(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddNode(4)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 3)
	gb.AddEdge(1, 4)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_node, err := graph.GetNode(1)
	assert.NoError(t, err)
	expected_node := rawUndirectedNode{ID: 1, Neighbors: []NodeID{2, 3, 4}}
	AssertNodeEquals(t, expected_node, actual_node)

	_, err = graph.GetNode(5)
	assert.ErrorIs(t, err, nodeNotFoundError{nodeID: 5})
}

func Test_DirectedGraphGetNode(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddNode(4)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 3)
	gb.AddEdge(1, 4)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_node, err := graph.GetNode(1)
	assert.NoError(t, err)
	expected_node := rawDirectedNode{ID: 1, Outgoing: []NodeID{2, 3, 4}, Incoming: []NodeID{}}
	AssertNodeEquals(t, expected_node, actual_node)

	_, err = graph.GetNode(5)
	assert.ErrorIs(t, err, nodeNotFoundError{nodeID: 5})
}

func AssertEdgeEquals(t *testing.T, expected, actual Edge) bool {
	refLessExpected := expected.removeRef()
	refLessActual := actual.removeRef()
	return assert.True(t,
		cmp.Equal(refLessExpected, refLessActual),
		cmp.Diff(refLessExpected, refLessActual),
	)
}

func Test_UndirectedGraphGetEdge(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	expected_edge := rawUndirectedEdge{Nodes: [2]NodeID{1, 2}}
	AssertEdgeEquals(t, expected_edge, actual_edge)

	actual_edge, err = graph.GetEdge(2, 1)
	assert.NoError(t, err)
	AssertEdgeEquals(t, expected_edge, actual_edge)

	_, err = graph.GetEdge(1, 3)
	assert.ErrorIs(t, err, edgeNotFoundError{fromID: 1, toID: 3})
}

func Test_DirectedGraphGetEdge(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	expected_edge := rawDirectedEdge{From: 1, To: 2}
	AssertEdgeEquals(t, expected_edge, actual_edge)

	_, err = graph.GetEdge(2, 1)
	assert.ErrorIs(t, err, edgeNotFoundError{fromID: 2, toID: 1})
}

func AssertNodesEquals(t *testing.T, expected, actual []Node) bool {
	refLessExpected := make([]Node, 0)
	for _, node := range expected {
		refLessExpected = append(refLessExpected, node.removeRef())
	}
	refLessActual := make([]Node, 0)
	for _, node := range actual {
		refLessActual = append(refLessActual, node.removeRef())
	}
	return assert.True(t,
		cmp.Equal(refLessExpected, refLessActual),
		cmp.Diff(refLessExpected, refLessActual),
	)
}

func Test_UndirectedGraphGetNodes(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1, "val1")
	gb.AddNode(2, "val2")
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_nodes, err := graph.GetNodes()
	assert.NoError(t, err)
	expected_nodes := []Node{
		rawUndirectedNode{ID: 1, Neighbors: []NodeID{2}, Value: wrappedValue{HasValue: true, RawValue: "val1"}},
		rawUndirectedNode{ID: 2, Neighbors: []NodeID{1}, Value: wrappedValue{HasValue: true, RawValue: "val2"}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)
}

func Test_DirectedGraphGetNodes(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1, "val1")
	gb.AddNode(2, "val2")
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_nodes, err := graph.GetNodes()
	assert.NoError(t, err)
	expected_nodes := []Node{
		rawDirectedNode{ID: 1, Outgoing: []NodeID{2}, Incoming: []NodeID{}, Value: wrappedValue{HasValue: true, RawValue: "val1"}},
		rawDirectedNode{ID: 2, Outgoing: []NodeID{}, Incoming: []NodeID{1}, Value: wrappedValue{HasValue: true, RawValue: "val2"}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)
}
func AssertEdgesEquals(t *testing.T, expected, actual []Edge) bool {
	refLessExpected := make([]Edge, 0)
	for _, edge := range expected {
		refLessExpected = append(refLessExpected, edge.removeRef())
	}
	refLessActual := make([]Edge, 0)
	for _, edge := range actual {
		refLessActual = append(refLessActual, edge.removeRef())
	}
	return assert.True(t,
		cmp.Equal(refLessExpected, refLessActual),
		cmp.Diff(refLessExpected, refLessActual),
	)
}

func Test_UndirectedGraphGetEdges(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2, "val1")
	gb.AddEdge(2, 3, "val2")
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_edges, err := graph.GetEdges()
	expected_edges := []Edge{
		rawUndirectedEdge{Nodes: [2]NodeID{1, 2}, Value: wrappedValue{HasValue: true, RawValue: "val1"}},
		rawUndirectedEdge{Nodes: [2]NodeID{2, 3}, Value: wrappedValue{HasValue: true, RawValue: "val2"}},
	}
	AssertEdgesEquals(t, expected_edges, actual_edges)
}

func Test_DirectedGraphGetEdges(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2, "val1")
	gb.AddEdge(2, 3, "val2")
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_edges, err := graph.GetEdges()
	expected_edges := []Edge{
		rawDirectedEdge{From: 1, To: 2, Value: wrappedValue{HasValue: true, RawValue: "val1"}},
		rawDirectedEdge{From: 2, To: 3, Value: wrappedValue{HasValue: true, RawValue: "val2"}},
	}
	AssertEdgesEquals(t, expected_edges, actual_edges)
}

func Test_UndirectedGraph_IsDirected(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	assert.Equal(t, false, graph.IsDirected())
}

func Test_DirectedGraph_IsDirected(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	assert.Equal(t, true, graph.IsDirected())
}
