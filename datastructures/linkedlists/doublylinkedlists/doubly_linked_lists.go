package doublylinkedlists

import (
	"errors"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new doubly linked lists
func New[T any](values ...T) *DoublyLinkedList[T] { _ = "STUB: not implemented"; return nil }

// DLLNode doubly linked list node
type DLLNode[T any] struct {
	Value    T
	Previous *DLLNode[T]
	Next     *DLLNode[T]
}

// DoublyLinkedList doubly linked list structure
type DoublyLinkedList[T any] struct {
	count int
	head  *DLLNode[T]
	tail  *DLLNode[T]
}

// IsEmpty checks if the list is empty
func (s *DoublyLinkedList[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the list
func (s *DoublyLinkedList[T]) Size() int {
	_ = "STUB: not implemented"

	// Clear clears list
	return 0
}

func (s *DoublyLinkedList[T]) Clear() { _ = "STUB: not implemented"; return }

// GetValues returns values from head to tail
func (s *DoublyLinkedList[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }

// GetReverseValues returns values from tail to head
func (s *DoublyLinkedList[T]) GetReverseValues() []T { _ = "STUB: not implemented"; return nil }

// GetIndexOf returns the index of the first occurence
func (s *DoublyLinkedList[T]) GetIndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

// GetLastIndexOf returns the index of the last occurence
func (s *DoublyLinkedList[T]) GetLastIndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

// Add add to the list
func (s *DoublyLinkedList[T]) Add(values ...T) { _ = "STUB: not implemented"; return }

// InsertAt insert value at specific index in the list
func (s *DoublyLinkedList[T]) InsertAt(index int, value T) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAt remove value from the list at specific index
func (s *DoublyLinkedList[T]) RemoveAt(index int) (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetValueAt returns value at a specific index in the list
func (s *DoublyLinkedList[T]) GetValueAt(index int) (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetFirstValue returns first value in the list
func (s *DoublyLinkedList[T]) GetFirstValue() (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetLastValue returns last value in the list
func (s *DoublyLinkedList[T]) GetLastValue() (res T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetHead returns head node of the list
func (s *DoublyLinkedList[T]) GetHead() *DLLNode[T] {
	_ = "STUB: not implemented"

	// GetTail returns tail node of the list
	return nil
}

func (s *DoublyLinkedList[T]) GetTail() *DLLNode[T] { _ = "STUB: not implemented"; return nil }

func (s *DoublyLinkedList[T]) getNode(index int) (*DLLNode[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
