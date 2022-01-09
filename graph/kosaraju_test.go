package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func removeRef(sccs []StronglyConnectedComponent) []StronglyConnectedComponent {
	reflessSccs := make([]StronglyConnectedComponent, 0)
	for _, scc := range sccs {
		reflessScc := make(StronglyConnectedComponent, 0)
		for _, node := range scc {
			reflessScc = append(reflessScc, node.removeRef())
		}
		reflessSccs = append(reflessSccs, reflessScc)
	}
	return reflessSccs
}

func AssertSCCsEquals(t *testing.T, expected, actual []StronglyConnectedComponent) {
	reflessExpected := removeRef(expected)
	reflessActual := removeRef(actual)
	assert.True(t, cmp.Equal(reflessExpected, reflessActual), cmp.Diff(reflessExpected, reflessActual))
}

func Test_Kosaraju_Directed(t *testing.T) {
	gb := NewGraphBuilder(BuilderOptions{IsDirected: true})
	for i := 1; i < 9; i++ {
		gb.AddNode(NodeID(i))
	}
	// cycle of 1, 2, 3, 4
	gb.AddEdge(1, 2)
	gb.AddEdge(2, 3)
	gb.AddEdge(3, 4)
	gb.AddEdge(4, 1)

	// this edge connects two cycles
	gb.AddEdge(3, 5)

	// cycle of 5, 6, 7
	gb.AddEdge(5, 6)
	gb.AddEdge(6, 7)
	gb.AddEdge(7, 5)

	// this edge connects the prev cycle to a single node
	gb.AddEdge(7, 8)

	graph, err := gb.Build()
	assert.NoError(t, err)

	actual_sccs := Kosaraju(graph)
	expected_sccs := []StronglyConnectedComponent{
		{rawDirectedNode{ID: 8, Incoming: []NodeID{7}, Outgoing: []NodeID{}}},
		{
			rawDirectedNode{ID: 5, Incoming: []NodeID{3, 7}, Outgoing: []NodeID{6}},
			rawDirectedNode{ID: 6, Incoming: []NodeID{5}, Outgoing: []NodeID{7}},
			rawDirectedNode{ID: 7, Incoming: []NodeID{6}, Outgoing: []NodeID{5, 8}},
		},
		{
			rawDirectedNode{ID: 1, Incoming: []NodeID{4}, Outgoing: []NodeID{2}},
			rawDirectedNode{ID: 2, Incoming: []NodeID{1}, Outgoing: []NodeID{3}},
			rawDirectedNode{ID: 3, Incoming: []NodeID{2}, Outgoing: []NodeID{4, 5}},
			rawDirectedNode{ID: 4, Incoming: []NodeID{3}, Outgoing: []NodeID{1}},
		},
	}
	AssertSCCsEquals(t, expected_sccs, actual_sccs)
}

func Test_Kosaraju_Undirected(t *testing.T) {
	gb := NewGraphBuilder()
	for i := 1; i < 5; i++ {
		gb.AddNode(NodeID(i))
	}
	gb.AddEdge(1, 2)
	gb.AddEdge(3, 4)
	graph, err := gb.Build()
	assert.NoError(t, err)
	actual_sccs := Kosaraju(graph)
	expected_sccs := []StronglyConnectedComponent{
		{
			rawUndirectedNode{ID: 1, Neighbors: []NodeID{2}},
			rawUndirectedNode{ID: 2, Neighbors: []NodeID{1}},
		},
		{
			rawUndirectedNode{ID: 3, Neighbors: []NodeID{4}},
			rawUndirectedNode{ID: 4, Neighbors: []NodeID{3}},
		},
	}
	AssertSCCsEquals(t, expected_sccs, actual_sccs)
}
