package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UndirectedNode_GetID(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_id := node.GetID()
	assert.Equal(t, NodeID(1), actual_id)
}

func Test_DirectedNode_GetID(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_id := node.GetID()
	assert.Equal(t, NodeID(1), actual_id)
}

func Test_UndirectedNode_GetIncomingEdges(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(3, 1)
	gb.AddEdge(2, 1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	_, err = node.GetIncomingEdges()
	assert.ErrorIs(t, err, cannotUseForUndirectedGraphError{methodName: "Node.GetIncomingEdges"})
}

func Test_DirectedNode_GetIncomingEdges(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(3, 1)
	gb.AddEdge(2, 1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_incoming, err := node.GetIncomingEdges()
	expected_incoming := []Edge{
		rawDirectedEdge{From: 2, To: 1},
		rawDirectedEdge{From: 3, To: 1},
	}
	AssertEdgesEquals(t, expected_incoming, actual_incoming)
}

func Test_UndirectedNode_GetOutgoingEdges(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 3)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	_, err = node.GetOutgoingEdges()
	assert.ErrorIs(t, err, cannotUseForUndirectedGraphError{methodName: "Node.GetOutgoingEdges"})
}

func Test_DirectedNode_GetOutgoingEdges(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(1, 3)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_outgoing, err := node.GetOutgoingEdges()
	assert.NoError(t, err)
	expected_outgoing := []Edge{
		rawDirectedEdge{From: 1, To: 2},
		rawDirectedEdge{From: 1, To: 3},
	}
	AssertEdgesEquals(t, expected_outgoing, actual_outgoing)
}

func Test_UndirectedNode_GetIncidentEdges(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_incident, err := node.GetIncidentEdges()
	assert.NoError(t, err)
	expected_incident := []Edge{
		rawUndirectedEdge{Nodes: [2]NodeID{1, 2}},
		rawUndirectedEdge{Nodes: [2]NodeID{1, 3}},
	}
	AssertEdgesEquals(t, expected_incident, actual_incident)
}

func Test_DirectedNode_GetIncidentEdges(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1)
	gb.AddNode(2)
	gb.AddNode(3)
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 1)
	graph, err := gb.Build()
	assert.NoError(t, err)
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_incident, err := node.GetIncidentEdges()
	assert.NoError(t, err)
	expected_incident := []Edge{
		rawDirectedEdge{From: 3, To: 1},
		rawDirectedEdge{From: 1, To: 2},
	}
	AssertEdgesEquals(t, expected_incident, actual_incident)
}

func Test_UndirectedNode_GetValue(t *testing.T) {
	gb := NewGraphBuilder()
	gb.AddNode(1, "val")
	gb.AddNode(2)
	graph, err := gb.Build()
	assert.NoError(t, err)

	// assert value in node 1
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_value, err := node.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "val", actual_value)

	// assert error in node 2
	node, err = graph.GetNode(2)
	assert.NoError(t, err)
	_, err = node.GetValue()
	assert.ErrorIs(t, err, noValueFoundInNodeError{2})
}

func Test_DirectedNode_GetValue(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	gb.AddNode(1, "val")
	gb.AddNode(2)
	graph, err := gb.Build()
	assert.NoError(t, err)

	// assert value in node 1
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	actual_value, err := node.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "val", actual_value)

	// assert error in node 2
	node, err = graph.GetNode(2)
	assert.NoError(t, err)
	_, err = node.GetValue()
	assert.ErrorIs(t, err, noValueFoundInNodeError{2})
}
