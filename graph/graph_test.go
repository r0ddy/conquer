package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UndirectedGraph_Node(t *testing.T) {
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
	expected_graph := RawGraph{}
	assert.ObjectsAreEqual(expected_graph, graph)

	// assert each node has correct id, edges and value
	node, err := graph.GetNode(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, node.GetID())
	expected_incoming_edges := []Edge{}
	expected_outgoing_edges := []Edge{}
	assert.ObjectsAreEqual(expected_incoming_edges, node.GetIncomingEdges())
	assert.ObjectsAreEqual(expected_outgoing_edges, node.GetOutgoingEdges())
	expected_value := 0
	value, err := node.GetValue()
	assert.NoError(t, err)
	assert.ObjectsAreEqual(expected_value, value)
}
