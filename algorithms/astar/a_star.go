package astar

import (
	"math"

	"github.com/shomali11/go-interview/datastructures/priorityqueues"
	"github.com/shomali11/go-interview/datastructures/sets/hashsets"
)

var (
	positiveInfinity = math.Inf(1)
)

// Node a node interface
type Node interface {
	EstimateDistanceToGoal(goal Node) float64
	ActualDistanceToNeighbor(neighbor Node) float64
	GetNeighbors() []Node
}

type wrapper struct {
	h      float64
	g      float64
	node   Node
	parent *wrapper
}

// New new factory
func New() *AStar { _ = "STUB: not implemented"; return nil }

// Tie Breaker ... Maximum G Value

// Maximum G Value

// Minimum F Value

// AStar a star
type AStar struct {
	openPQ         *priorityqueues.PriorityQueue[*wrapper]
	closedSet      *hashsets.HashSet[*wrapper]
	nodeWrapperMap map[Node]*wrapper
}

// Search search
func (s *AStar) Search(start Node, goal Node) ([]Node, float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// While not Empty ...

// Extract Top of the Heap ...

// Did we reach Goal ?

// Shortest Path ...

// Set Node to Visited ...

// Get List of neighbors ..

// Traverse the neighbors ...

// We only care about the ones we did not visit ...

// Is not in the Heap ..

func (s *AStar) updateVertex(currentWrapper *wrapper, neighborWrapper *wrapper, goalWrapper *wrapper) {
	_ = "STUB: not implemented"
	// Check for Consistency ..
	return
}

// Update the G Function

// If Min Heap Has the Node ..

// Remove it ..

// Add the state with the new function value ..

func (s *AStar) clear() { _ = "STUB: not implemented"; return }

func (s *AStar) getWrapper(node Node) *wrapper { _ = "STUB: not implemented"; return nil }

func getPath(goalWrapper *wrapper) []Node { _ = "STUB: not implemented"; return nil }
