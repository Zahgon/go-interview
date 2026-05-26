package hashsets

// New factory that creates a hash set
func New[T comparable](values ...T) *HashSet[T] { _ = "STUB: not implemented"; return nil }

// HashSet datastructure
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// Add adds values to the set
func (s *HashSet[T]) Add(values ...T) { _ = "STUB: not implemented"; return }

// Remove removes values from the set
func (s *HashSet[T]) Remove(values ...T) { _ = "STUB: not implemented"; return }

// Contains checks if the value is in the set
func (s *HashSet[T]) Contains(value T) bool { _ = "STUB: not implemented"; return false }

// ContainsAll checks if all values are in the set
func (s *HashSet[T]) ContainsAll(values ...T) bool { _ = "STUB: not implemented"; return false }

// ContainsAny checks if any of the values are in the set
func (s *HashSet[T]) ContainsAny(values ...T) bool { _ = "STUB: not implemented"; return false }

// Merge the two sets
func (s *HashSet[T]) Merge(sets ...*HashSet[T]) { _ = "STUB: not implemented"; return }

// Clear clears set
func (s *HashSet[T]) Clear() { _ = "STUB: not implemented"; return }

// GetValues returns values
func (s *HashSet[T]) GetValues() []T { _ = "STUB: not implemented"; return nil }

// IsEmpty checks if the set is empty
func (s *HashSet[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the set
func (s *HashSet[T]) Size() int {
	_ = "STUB: not implemented"

	// Common set functions
	return 0
}

// Copy makes an identical copy of the set
func (s *HashSet[T]) Copy() *HashSet[T] { _ = "STUB: not implemented"; return nil }

// Union makes a set that has all of the elements in either of two sets
func (s *HashSet[T]) Union(ss *HashSet[T]) *HashSet[T] { _ = "STUB: not implemented"; return nil }

// Intersection makes a set that has only the elements common to both of two sets
func (s *HashSet[T]) Intersection(ss *HashSet[T]) *HashSet[T] {
	_ = "STUB: not implemented"
	return nil
}

// SymmetricDifference makes a set that has elements that are in one of two sets, but not both
func (s *HashSet[T]) SymmetricDifference(ss *HashSet[T]) *HashSet[T] {
	_ = "STUB: not implemented"
	return nil
}

// Subtraction makes a set with the elements that are in the first set, but not the second
func (s *HashSet[T]) Subtraction(ss *HashSet[T]) *HashSet[T] { _ = "STUB: not implemented"; return nil }
