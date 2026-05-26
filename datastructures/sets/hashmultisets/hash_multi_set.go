package hashmultisets

// New factory that creates a new Hash Multi Set
func New[T comparable](values ...T) *HashMultiSet[T] { _ = "STUB: not implemented"; return nil }

// MultiSetPair a set's key/count pair
type MultiSetPair[T comparable] struct {
	Key   T
	Count int
}

// HashMultiSet a data structure representing a set with counts
type HashMultiSet[T comparable] struct {
	data map[T]int
}

// Merge merge multiple sets
func (s *HashMultiSet[T]) Merge(sets ...*HashMultiSet[T]) { _ = "STUB: not implemented"; return }

// Add adds a value to the set
func (s *HashMultiSet[T]) Add(values ...T) { _ = "STUB: not implemented"; return }

// IncrementBy increments a value's count by a number
func (s *HashMultiSet[T]) IncrementBy(value T, count int) { _ = "STUB: not implemented"; return }

// GetValues returns a list of the set's values
func (s *HashMultiSet[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }

// Contains checks if a value is in the set
func (s *HashMultiSet[T]) Contains(value T) bool { _ = "STUB: not implemented"; return false }

// ContainsAll checks if all values are in the set
func (s *HashMultiSet[T]) ContainsAll(values ...T) bool { _ = "STUB: not implemented"; return false }

// ContainsAny checks if any values are in the set
func (s *HashMultiSet[T]) ContainsAny(values ...T) bool { _ = "STUB: not implemented"; return false }

// GetCount returns count associated with the value
func (s *HashMultiSet[T]) GetCount(value T) int { _ = "STUB: not implemented"; return 0 }

// Remove removes a value
func (s *HashMultiSet[T]) Remove(values ...T) { _ = "STUB: not implemented"; return }

// Clear clears the set
func (s *HashMultiSet[T]) Clear() { _ = "STUB: not implemented"; return }

// IsEmpty checks if the set is empty
func (s *HashMultiSet[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the set
func (s *HashMultiSet[T]) Size() int {
	_ = "STUB: not implemented"

	// GetTopValues returns values ordered in descending order
	return 0
}

func (s *HashMultiSet[T]) GetTopValues() []MultiSetPair[T] { _ = "STUB: not implemented"; return nil }
