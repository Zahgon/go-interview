package priorityqueues

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/maps/hashmultimaps"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new priority queues
func New[T comparable](compare func(i, j T) bool, values ...T) *PriorityQueue[T] {
	_ = "STUB: not implemented"
	return nil
}

// PriorityQueue Priority Queue structure
type PriorityQueue[T comparable] struct {
	pq       *heapArray[T]
	multiMap *hashmultimaps.HashMultiMap[T, *pqElement[T]]
}

// Push pushes to the Priority Queue
func (s *PriorityQueue[T]) Push(values ...T) { _ = "STUB: not implemented"; return }

// Contains checks if the value exists in the Priority Queue
func (s *PriorityQueue[T]) Contains(value T) bool { _ = "STUB: not implemented"; return false }

// Remove removes from the Priority Queue
func (s *PriorityQueue[T]) Remove(values ...T) { _ = "STUB: not implemented"; return }

// IsEmpty checks if the Priority Queue is empty
func (s *PriorityQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the Priority Queue
func (s *PriorityQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Clear clears the Priority Queue
func (s *PriorityQueue[T]) Clear() {
	_ = "STUB: not implemented"

	// Pop removes from the Priority Queue
	return
}

func (s *PriorityQueue[T]) Pop() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// Peek returns top of the Priority Queue
func (s *PriorityQueue[T]) Peek() (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
