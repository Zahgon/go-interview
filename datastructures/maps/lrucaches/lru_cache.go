package lrucaches

import (
	"container/list"
)

// New factory that creates a new LRU Cache
func New[K comparable, V any](capacity int) *LRUCache[K, V] { _ = "STUB: not implemented"; return nil }

// LRUEntry holds a key/value pair
type LRUEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// LRUCache a data structure representing a map of keys with lists of values
type LRUCache[K comparable, V any] struct {
	capacity   int
	linkedList *list.List
	hashMap    map[K]*list.Element
}

// Merge merge multiple lru caches
func (s *LRUCache[K, V]) Merge(caches ...*LRUCache[K, V]) { _ = "STUB: not implemented"; return }

// Put key/value pair into the cache
func (s *LRUCache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

// GetKeys returns a list of the cache's keys in most recently used order
func (s *LRUCache[K, V]) GetKeys() []K { _ = "STUB: not implemented"; return nil }

// GetEntries returns a list of the cache's entries in most recently used order
func (s *LRUCache[K, V]) GetEntries() []*LRUEntry[K, V] { _ = "STUB: not implemented"; return nil }

// Contains checks if a key is in cache
func (s *LRUCache[K, V]) Contains(key K) bool { _ = "STUB: not implemented"; return false }

// ContainsAll checks if all keys are in the cache
func (s *LRUCache[K, V]) ContainsAll(keys ...K) bool { _ = "STUB: not implemented"; return false }

// ContainsAny checks if any keys are in the cache
func (s *LRUCache[K, V]) ContainsAny(keys ...K) bool { _ = "STUB: not implemented"; return false }

// GetValue returns value associated with the key
func (s *LRUCache[K, V]) GetValue(key K) (res V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Remove removes a key and its value
func (s *LRUCache[K, V]) Remove(keys ...K) { _ = "STUB: not implemented"; return }

// Clear clears the multiMap
func (s *LRUCache[K, V]) Clear() { _ = "STUB: not implemented"; return }

// IsEmpty checks if the multiMap is empty
func (s *LRUCache[K, V]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Size returns size of the multiMap
func (s *LRUCache[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *LRUCache[K, V]) remove(key K) { _ = "STUB: not implemented"; return }
