package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DuplicateNodeError(t *testing.T) {
	actual_error := duplicateNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "node with id 1 has already been added")
}

func Test_DuplicateEdgeError(t *testing.T) {
	actual_error := duplicateEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "edge from 1 to 2 has already been added")
}

func Test_RedundantEdgeError(t *testing.T) {
	actual_error := redundantEdgeError{nodeID: 1}
	assert.EqualError(t, actual_error, "edge from 1 to 1 is redundant")
}

func Test_NodeNotFoundError(t *testing.T) {
	actual_error := nodeNotFoundError{nodeID: 1}
	assert.EqualError(t, actual_error, "node with id 1 could not be found")
}

func Test_MultipleValuesForNodeError(t *testing.T) {
	actual_error := multipleValuesForNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "multiple values provided for node with id 1")
}

func Test_MultipleValuesForEdgeError(t *testing.T) {
	actual_error := multipleValuesForEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "multiple values provided for edge from 1 to 2")
}

func Test_EdgeNotFoundError(t *testing.T) {
	actual_error := edgeNotFoundError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "edge from 1 to 2 could not be found")
}

func Test_NoValueFoundInNodeError(t *testing.T) {
	actual_error := noValueFoundInNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "no value found in node with id 1")
}

func Test_NoValueFoundInEdgeError(t *testing.T) {
	actual_error := noValueFoundInEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "no value found in edge from 1 to 2")
}

func Test_CannotUseForDirectedGraphError(t *testing.T) {
	actual_error := cannotUseForDirectedGraphError{methodName: "Node.GetID"}
	assert.EqualError(t, actual_error, "cannot use Node.GetID on directed graph")
}

func Test_CannotUseForUndirectedGraphError(t *testing.T) {
	actual_error := cannotUseForUndirectedGraphError{methodName: "Node.GetID"}
	assert.EqualError(t, actual_error, "cannot use Node.GetID on undirected graph")
}
