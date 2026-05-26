package slicestacks

import (
	"errors"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New[T any](values ...T) *Stack[T] { _ = "STUB: not implemented"; return nil }

// Stack stack structure
type Stack[T any] struct {
	array []T
}

// Push add to the stack
func (s *Stack[T]) Push(values ...T) { _ = "STUB: not implemented"; return }

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the stack
func (s *Stack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Clear clears stack
func (s *Stack[T]) Clear() {
	_ = "STUB: not implemented"

	// Pop remove from the stack
	return
}

func (s *Stack[T]) Pop() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// Peek returns top of the stack
func (s *Stack[T]) Peek() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// GetValues returns values
func (s *Stack[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }
