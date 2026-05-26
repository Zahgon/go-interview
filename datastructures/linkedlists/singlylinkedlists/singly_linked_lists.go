package singlylinkedlists

import (
	"errors"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new singly linked lists
func New[T any](values ...T) *SinglyLinkedList[T] { _ = "STUB: not implemented"; return nil }

// SLLNode singly linked list node
type SLLNode[T any] struct {
	Value T
	Next  *SLLNode[T]
}

// SinglyLinkedList singly linked list structure
type SinglyLinkedList[T any] struct {
	count int
	head  *SLLNode[T]
	tail  *SLLNode[T]
}

// IsEmpty checks if the list is empty
func (s *SinglyLinkedList[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the list
func (s *SinglyLinkedList[T]) Size() int {
	_ = "STUB: not implemented"

	// Clear clears list
	return 0
}

func (s *SinglyLinkedList[T]) Clear() { _ = "STUB: not implemented"; return }

// GetValues returns values
func (s *SinglyLinkedList[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }

// GetIndexOf returns the index of the first occurence
func (s *SinglyLinkedList[T]) GetIndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

// Add add to the list
func (s *SinglyLinkedList[T]) Add(values ...T) { _ = "STUB: not implemented"; return }

// InsertAt insert value at specific index in the list
func (s *SinglyLinkedList[T]) InsertAt(index int, value T) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAt remove value from the list at specific index
func (s *SinglyLinkedList[T]) RemoveAt(index int) (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetValueAt returns value at a specific index in the list
func (s *SinglyLinkedList[T]) GetValueAt(index int) (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetFirstValue returns first value in the list
func (s *SinglyLinkedList[T]) GetFirstValue() (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetLastValue returns last value in the list
func (s *SinglyLinkedList[T]) GetLastValue() (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetHead returns head node of the list
func (s *SinglyLinkedList[T]) GetHead() *SLLNode[T] { _ = "STUB: not implemented"; return nil }

func (s *SinglyLinkedList[T]) getNode(index int) (*SLLNode[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
