package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UndirectedEdge_GetTo(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	_, err = edge.GetTo()
	assert.ErrorIs(t, err, CannotUseForUndirectedGraphError{"Edge.GetTo"})
}

func Test_DirectedEdge_GetTo(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_node, err := edge.GetTo()
	assert.NoError(t, err)
	expected_node := rawDirectedNode{ID: 2, Incoming: []NodeID{1}, Outgoing: []NodeID{}}
	AssertNodeEquals(t, expected_node, actual_node)
}

func Test_UndirectedEdge_GetFrom(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	_, err = edge.GetFrom()
	assert.ErrorIs(t, err, CannotUseForUndirectedGraphError{"Edge.GetFrom"})
}

func Test_DirectedEdge_GetFrom(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_node, err := edge.GetFrom()
	assert.NoError(t, err)
	expected_node := rawDirectedNode{ID: 1, Incoming: []NodeID{}, Outgoing: []NodeID{2}}
	AssertNodeEquals(t, expected_node, actual_node)
}

func Test_UndirectedEdge_GetNodes(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_nodes, err := edge.GetNodes()
	assert.NoError(t, err)
	expected_nodes := []Node{
		rawUndirectedNode{ID: 1, Neighbors: []NodeID{2}},
		rawUndirectedNode{ID: 2, Neighbors: []NodeID{1}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)

	// test redundant edge
	gb = NewGraphBuilder(BuilderOptions{AllowRedundantEdges: true})
	gb.AddNode(1)
	gb.AddEdge(1, 1)
	graph, err = gb.Build()
	assert.NoError(t, err)
	edge, err = graph.GetEdge(1, 1)
	assert.NoError(t, err)
	actual_nodes, err = edge.GetNodes()
	assert.NoError(t, err)
	expected_nodes = []Node{
		rawUndirectedNode{ID: 1, Neighbors: []NodeID{1}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)
}

func Test_DirectedEdge_GetNodes(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddEdge(1, 2)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_nodes, err := edge.GetNodes()
	assert.NoError(t, err)
	expected_nodes := []Node{
		rawDirectedNode{ID: 1, Outgoing: []NodeID{2}, Incoming: []NodeID{}},
		rawDirectedNode{ID: 2, Outgoing: []NodeID{}, Incoming: []NodeID{1}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)

	// test redundant edge
	gb = NewGraphBuilder(BuilderOptions{IsDirected: true, AllowRedundantEdges: true})
	gb.AddNode(1)
	gb.AddEdge(1, 1)
	graph, err = gb.Build()
	assert.NoError(t, err)
	edge, err = graph.GetEdge(1, 1)
	assert.NoError(t, err)
	actual_nodes, err = edge.GetNodes()
	assert.NoError(t, err)
	expected_nodes = []Node{
		rawDirectedNode{ID: 1, Outgoing: []NodeID{1}, Incoming: []NodeID{1}},
	}
	AssertNodesEquals(t, expected_nodes, actual_nodes)
}

func Test_UndirectedEdge_GetValue(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2, "val")
	gb.AddEdge(1, 3)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_value, err := edge.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "val", actual_value)

	edge, err = graph.GetEdge(1, 3)
	assert.NoError(t, err)
	_, err = edge.GetValue()
	assert.ErrorIs(t, err, NoValueFoundInEdgeError{fromID: 1, toID: 3})
}

func Test_DirectedEdge_GetValue(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2, "val")
	gb.AddEdge(1, 3)
	graph, err := gb.Build()
	assert.NoError(t, err)
	edge, err := graph.GetEdge(1, 2)
	assert.NoError(t, err)
	actual_value, err := edge.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "val", actual_value)

	edge, err = graph.GetEdge(1, 3)
	assert.NoError(t, err)
	_, err = edge.GetValue()
	assert.ErrorIs(t, err, NoValueFoundInEdgeError{fromID: 1, toID: 3})
}
