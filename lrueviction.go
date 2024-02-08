package inmemocache

import (
	"errors"
)

// LRU get policy that follows Least Recently Used rule for eviction
func LRUEvictionPolicyGet[Key comparable, Val any](key Key, cache *Cache[Key, Val]) (any, error) {
	if elem, ok := cache.Data[key]; ok {

		// increase the priority the element being access(used)
		cache.Els.MoveToFront(elem)
		return elem.Value.(*CacheEntry[Key, Val]).Val, nil
	}
	return nil, errors.New("Could not retrieve value from the cache")
}

// LRU put policy that follows Least Recently Used rule for eviction
func LRUEvictionPolicyPut[Key comparable, Val any](key Key, value Val, cache *Cache[Key, Val]) error {

	// alter the list node if entry is already there
	if elem, ok := cache.Data[key]; ok {
		cache.Els.PushFront(elem)
		listEntryOfKey := cache.Data[key]
		listEntryOfKey.Value = &CacheEntry[Key, Val]{Key: key, Val: value}
	} else {

		// create new node if no entry for key is present
		cacheEntryToBeAdded := &CacheEntry[Key, Val]{Key: key, Val: value}
		el := cache.Els.PushFront(cacheEntryToBeAdded)
		cache.Data[key] = el
	}

	if len(cache.Data) > cache.Capacity {
		// Evict the least recently used element

		tail := cache.Els.Front()
		for it := cache.Els.Front(); it.Next() != nil; it = it.Next() {
			tail = it
		}
		//remove from Cache map
		delete(cache.Data, tail.Value.(*CacheEntry[Key, Val]).Key)
		//remove from linkedlist
		cache.Els.Remove(tail)
	}

	return nil
}
