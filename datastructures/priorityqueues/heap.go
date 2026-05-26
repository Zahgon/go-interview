package priorityqueues

/*
	https://golang.org/pkg/container/heap/
*/

// pqElement is an value in a priority queue.
type pqElement[T any] struct {
	value T

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A heapArray implements heap.Interface and holds Items.
type heapArray[T any] struct {
	compare func(i, j T) bool
	array   []*pqElement[T]
}

// Len length of the queue
func (h heapArray[T]) Len() int {
	_ = "STUB: not implemented"

	// Less compares priority of two elements in the queue
	return 0
}

func (h heapArray[T]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap swaps two elements in the queue
func (h heapArray[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Push pushes to the queue
func (h *heapArray[T]) Push(x any) { _ = "STUB: not implemented"; return }

// Pop pops from the queue
func (h *heapArray[T]) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// for safety

// update modifies the priority and value of an Item in the queue.
// func (h *heapArray) update(element *pqElement, value string, priority int) {
// 	element.Value = value
// 	element.Priority = priority
// 	heap.Fix(h, element.index)
// }
