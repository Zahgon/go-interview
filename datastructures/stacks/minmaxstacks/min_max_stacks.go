package minmaxstacks

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New[T any](compare func(i, j T) bool, values ...T) *Stack[T] {
	_ = "STUB: not implemented"
	return nil
}

// Stack stack structure
type Stack[T any] struct {
	valuesStack *slicestacks.Stack[T]
	minMaxStack *slicestacks.Stack[[]T]
	compare     func(i, j T) bool
}

// Push add to the stack
func (s *Stack[T]) Push(values ...T) { _ = "STUB: not implemented"; return }

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the stack
func (s *Stack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Clear clears stack
func (s *Stack[T]) Clear() { _ = "STUB: not implemented"; return }

// Pop remove from the stack
func (s *Stack[T]) Pop() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// Peek returns top of the stack
func (s *Stack[T]) Peek() (res T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// GetValues returns values
func (s *Stack[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }

// GetMinMax returns the min and max values
func (s *Stack[T]) GetMinMax() (min T, max T, err error) {
	_ = "STUB: not implemented"
	return *new(T), *new(T), nil
}
