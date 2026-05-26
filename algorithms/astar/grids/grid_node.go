package grids

import (
	"github.com/shomali11/go-interview/algorithms/astar"
)

// GridNode grid node
type GridNode struct {
	X          int
	Y          int
	Cost       float64
	IsObstacle bool
	world      *GridWorld
}

// EstimateDistanceToGoal get distance to the goal
func (n *GridNode) EstimateDistanceToGoal(goal astar.Node) float64 {
	_ = "STUB: not implemented"
	return 0
}

// ActualDistanceToNeighbor get distance to a neighbor
func (n *GridNode) ActualDistanceToNeighbor(node astar.Node) float64 {
	_ = "STUB: not implemented"
	return 0
}

// GetNeighbors get neighbors
func (n *GridNode) GetNeighbors() []astar.Node { _ = "STUB: not implemented"; return nil }

// If Same Object ...

// Skip ...

func (n *GridNode) isBehindObstacles(node *GridNode) bool { _ = "STUB: not implemented"; return false }

// Lower Right

// Upper Left

// Upper Right

// Lower Left

func isDiagonal(n *GridNode, node *GridNode) bool { _ = "STUB: not implemented"; return false }

// Lower Right

// Upper Left

// Upper Right

// Lower Left

func distance(current *GridNode, other *GridNode) float64 { _ = "STUB: not implemented"; return 0 }
