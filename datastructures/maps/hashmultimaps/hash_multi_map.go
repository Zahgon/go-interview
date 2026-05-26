package hashmultimaps

// New factory that creates a new Hash Multi Map
func New[K, V comparable]() *HashMultiMap[K, V] { _ = "STUB: not implemented"; return nil }

// HashMultiMap a data structure representing a map of keys with lists of values
type HashMultiMap[K, V comparable] struct {
	data map[K][]V
}

// Merge merge multiple multi maps
func (s *HashMultiMap[K, V]) Merge(maps ...*HashMultiMap[K, V]) { _ = "STUB: not implemented"; return }

// Put key/value pair to the multi map
func (s *HashMultiMap[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

// PutAll put key/values to the multi map
func (s *HashMultiMap[K, V]) PutAll(key K, values ...V) { _ = "STUB: not implemented"; return }

// GetKeys returns a list of the multi map's keys
func (s *HashMultiMap[K, V]) GetKeys() []K { _ = "STUB: not implemented"; return nil }

// Contains checks if a key is in the multi map
func (s *HashMultiMap[K, V]) Contains(key K) bool { _ = "STUB: not implemented"; return false }

// ContainsAll checks if all keys are in the multi map
func (s *HashMultiMap[K, V]) ContainsAll(keys ...K) bool { _ = "STUB: not implemented"; return false }

// ContainsAny checks if any keys are in the multi map
func (s *HashMultiMap[K, V]) ContainsAny(keys ...K) bool { _ = "STUB: not implemented"; return false }

// GetValues returns values associated with the key
func (s *HashMultiMap[K, V]) GetValues(key K) []V { _ = "STUB: not implemented"; return nil }

// RemoveKey removes a key and all its values
func (s *HashMultiMap[K, V]) RemoveKey(keys ...K) { _ = "STUB: not implemented"; return }

// Remove removes a value from a key's values
func (s *HashMultiMap[K, V]) Remove(key K, value V) { _ = "STUB: not implemented"; return }

// Clear clears the multiMap
func (s *HashMultiMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

// IsEmpty checks if the multiMap is empty
func (s *HashMultiMap[K, V]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the multiMap
func (s *HashMultiMap[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func getIndex[V comparable](value V, values ...V) int { _ = "STUB: not implemented"; return 0 }
