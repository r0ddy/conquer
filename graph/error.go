package graph

import "fmt"

type duplicateNodeError struct {
	nodeID NodeID
}

func (e duplicateNodeError) Error() string {
	return fmt.Sprintf("node with id %d has already been added", e.nodeID)
}

type duplicateEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e duplicateEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d has already been added", e.fromID, e.toID)
}

type redundantEdgeError struct {
	nodeID NodeID
}

func (e redundantEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d is redundant", e.nodeID, e.nodeID)
}

type nodeNotFoundError struct {
	nodeID NodeID
}

func (e nodeNotFoundError) Error() string {
	return fmt.Sprintf("node with id %d could not be found", e.nodeID)
}

type multipleValuesForNodeError struct {
	nodeID NodeID
}

func (e multipleValuesForNodeError) Error() string {
	return fmt.Sprintf("multiple values provided for node with id %d", e.nodeID)
}

type multipleValuesForEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e multipleValuesForEdgeError) Error() string {
	return fmt.Sprintf("multiple values provided for edge from %d to %d", e.fromID, e.toID)
}

type edgeNotFoundError struct {
	fromID NodeID
	toID   NodeID
}

func (e edgeNotFoundError) Error() string {
	return fmt.Sprintf("edge from %d to %d could not be found", e.fromID, e.toID)
}

type noValueFoundInNodeError struct {
	nodeID NodeID
}

func (e noValueFoundInNodeError) Error() string {
	return fmt.Sprintf("no value found in node with id %d", e.nodeID)
}

type noValueFoundInEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (e noValueFoundInEdgeError) Error() string {
	return fmt.Sprintf("no value found in edge from %d to %d", e.fromID, e.toID)
}

type cannotUseForDirectedGraphError struct {
	methodName string
}

func (e cannotUseForDirectedGraphError) Error() string {
	return fmt.Sprintf("cannot use %s on directed graph", e.methodName)
}

type cannotUseForUndirectedGraphError struct {
	methodName string
}

func (e cannotUseForUndirectedGraphError) Error() string {
	return fmt.Sprintf("cannot use %s on undirected graph", e.methodName)
}
