package queues

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new Queues
func New[T any](values ...T) *Queue[T] { _ = "STUB: not implemented"; return nil }

// Queue Queue structure
type Queue[T any] struct {
	array []T
}

// Enqueue add to the Queue
func (q *Queue[T]) Enqueue(values ...T) { _ = "STUB: not implemented"; return }

// IsEmpty checks if the Queue is empty
func (q *Queue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the Queue
func (q *Queue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Clear clears Queue
func (q *Queue[T]) Clear() {
	_ = "STUB: not implemented"

	// Dequeue remove from the Queue
	return
}

func (q *Queue[T]) Dequeue() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// Peek returns front of the Queue
func (q *Queue[T]) Peek() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// GetValues returns values
func (q *Queue[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }
