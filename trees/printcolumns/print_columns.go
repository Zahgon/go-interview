package printcolumns

import (
	"github.com/shomali11/go-interview/datastructures/trees"

	"golang.org/x/exp/constraints"
)

// PrintColumns prints a tree column by column
func PrintColumns[T constraints.Ordered](node *trees.BinaryNode[T]) {
	_ = "STUB: not implemented"
	return
}

func minMax[T constraints.Ordered](values ...T) (T, T) {
	_ = "STUB: not implemented"
	return *new(T), *new(T)
}
