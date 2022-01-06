package graph

import "fmt"

type DuplicateNodeError struct {
	nodeID NodeID
}

func (e DuplicateNodeError) Error() string {
	return fmt.Sprintf("node with id %d has already been added", e.nodeID)
}

type DuplicateEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e DuplicateEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d has already been added", e.fromID, e.toID)
}

type RedundantEdgeError struct {
	nodeID NodeID
}

func (e RedundantEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d is redundant", e.nodeID, e.nodeID)
}

type NodeNotFoundError struct {
	nodeID NodeID
}

func (e NodeNotFoundError) Error() string {
	return fmt.Sprintf("node with id %d could not be found", e.nodeID)
}

type MultipleValuesForNodeError struct {
	nodeID NodeID
}

func (e MultipleValuesForNodeError) Error() string {
	return fmt.Sprintf("multiple values provided for node with id %d", e.nodeID)
}

type MultipleValuesForEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e MultipleValuesForEdgeError) Error() string {
	return fmt.Sprintf("multiple values provided for edge from %d to %d", e.fromID, e.toID)
}

type EdgeNotFoundError struct {
	fromID NodeID
	toID   NodeID
}

func (e EdgeNotFoundError) Error() string {
	return fmt.Sprintf("edge from %d to %d could not be found", e.fromID, e.toID)
}

type NoValueFoundInNodeError struct {
	nodeID NodeID
}

func (e NoValueFoundInNodeError) Error() string {
	return fmt.Sprintf("no value found in node with id %d", e.nodeID)
}

type NoValueFoundInEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e NoValueFoundInEdgeError) Error() string {
	return fmt.Sprintf("no value found in edge from %d to %d", e.fromID, e.toID)
}

type CannotUseForDirectedGraphError struct {
	methodName string
}

func (e CannotUseForDirectedGraphError) Error() string {
	return fmt.Sprintf("cannot use %s on directed graph", e.methodName)
}

type CannotUseForUndirectedGraphError struct {
	methodName string
}

func (e CannotUseForUndirectedGraphError) Error() string {
	return fmt.Sprintf("cannot use %s on undirected graph", e.methodName)
}
