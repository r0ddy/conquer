package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DuplicateNodeError(t *testing.T) {
	actual_error := DuplicateNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "node with id 1 has already been added")
}

func Test_DuplicateEdgeError(t *testing.T) {
	actual_error := DuplicateEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "edge from 1 to 2 has already been added")
}

func Test_RedundantEdgeError(t *testing.T) {
	actual_error := RedundantEdgeError{nodeID: 1}
	assert.EqualError(t, actual_error, "edge from 1 to 1 is redundant")
}

func Test_NodeNotFoundError(t *testing.T) {
	actual_error := NodeNotFoundError{nodeID: 1}
	assert.EqualError(t, actual_error, "node with id 1 could not be found")
}

func Test_MultipleValuesForNodeError(t *testing.T) {
	actual_error := MultipleValuesForNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "multiple values provided for node with id 1")
}

func Test_MultipleValuesForEdgeError(t *testing.T) {
	actual_error := MultipleValuesForEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "multiple values provided for edge from 1 to 2")
}

func Test_EdgeNotFoundError(t *testing.T) {
	actual_error := EdgeNotFoundError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "edge from 1 to 2 could not be found")
}

func Test_NoValueFoundInNodeError(t *testing.T) {
	actual_error := NoValueFoundInNodeError{nodeID: 1}
	assert.EqualError(t, actual_error, "no value found in node with id 1")
}

func Test_NoValueFoundInEdgeError(t *testing.T) {
	actual_error := NoValueFoundInEdgeError{fromID: 1, toID: 2}
	assert.EqualError(t, actual_error, "no value found in edge from 1 to 2")
}

func Test_CannotUseForDirectedGraphError(t *testing.T) {
	actual_error := CannotUseForDirectedGraphError{methodName: "Node.GetID"}
	assert.EqualError(t, actual_error, "cannot use Node.GetID on directed graph")
}

func Test_CannotUseForUndirectedGraphError(t *testing.T) {
	actual_error := CannotUseForUndirectedGraphError{methodName: "Node.GetID"}
	assert.EqualError(t, actual_error, "cannot use Node.GetID on undirected graph")
}
