package grids

import (
	"github.com/shomali11/go-interview/algorithms/astar"
)

// New creates a 2D grid of nodes
func New(rows, columns int) *GridWorld { _ = "STUB: not implemented"; return nil }

// GridWorld grid world
type GridWorld struct {
	grid [][]*GridNode
}

// Get retrieves a Grid Node
func (n *GridWorld) Get(i int, j int) *GridNode { _ = "STUB: not implemented"; return nil }

// IsValid checks whether coordinates are within the grid
func (n *GridWorld) IsValid(i int, j int) bool { _ = "STUB: not implemented"; return false }

// PrintGrid prints a Grid
func (n *GridWorld) PrintGrid(nodes []astar.Node) { _ = "STUB: not implemented"; return }
