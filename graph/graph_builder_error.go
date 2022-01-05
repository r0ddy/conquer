package graph

import "fmt"

type DuplicateNodeError struct {
	nodeID NodeID
}

func (dne DuplicateNodeError) Error() string {
	return fmt.Sprintf("node with id %d has already been added", dne.nodeID)
}

type DuplicateEdgeError struct {
	fromID NodeID
	toID   NodeID
}

func (dee DuplicateEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d has already been added", dee.fromID, dee.toID)
}

type RedundantEdgeError struct {
	nodeID NodeID
}

func (ree RedundantEdgeError) Error() string {
	return fmt.Sprintf("edge from %d to %d is redundant", ree.nodeID, ree.nodeID)
}

type NodeNotFoundError struct {
	nodeID NodeID
}

func (nnfe NodeNotFoundError) Error() string {
	return fmt.Sprintf("node with id %d could not be found", nnfe.nodeID)
}
