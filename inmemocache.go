package inmemocache

import "container/list"

// The EvictionPolicyGet function signatue shows how a custom get policy should behave
type EvictionPolicyGet[Key comparable, Val any] func(key Key, cache *Cache[Key, Val]) (any, error)

// The EvictionPolicyPut function signatue shows how a custom put policy should behave
type EvictionPolicyPut[Key comparable, Val any] func(key Key, val Val, cache *Cache[Key, Val]) error

// Each Entry in Cache is a Cache Entry
type CacheEntry[Key comparable, Val any] struct {
	Key Key
	Val Val
}

// The Cache has CacheEntries stored in a map and a linkedlist. map has Key type as key and the LinkedList Node reference as value.
// Linkedlist here is used for easier modification with respect to priority
type Cache[Key comparable, Val any] struct {
	Capacity          int
	Data              map[Key]*list.Element
	Els               list.List
	evictionPolicyGet EvictionPolicyGet[Key, Val]
	evictionPolicyPut EvictionPolicyPut[Key, Val]
}

// Creates a new cache with a given capacity, key-val type and get/put policies
func NewCache[Key comparable, Val any](Capacity int, evictionPolicyGet EvictionPolicyGet[Key, Val], evictionPolicyPut EvictionPolicyPut[Key, Val]) *Cache[Key, Val] {
	return &Cache[Key, Val]{
		Capacity:          Capacity,
		Data:              make(map[Key]*list.Element),
		Els:               *list.New(),
		evictionPolicyGet: evictionPolicyGet,
		evictionPolicyPut: evictionPolicyPut,
	}
}

// Wrapper for Cache get policy that exposes the access to cache
func (cache *Cache[Key, Val]) Get(key Key) (any, error) {
	return cache.evictionPolicyGet(key, cache)
}

// Wrapper for Cache put policy that exposes the write access to cache data
func (cache *Cache[Key, Val]) Put(key Key, val Val) error {
	return cache.evictionPolicyPut(key, val, cache)
}
