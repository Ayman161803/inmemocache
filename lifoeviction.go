package inmemocache

import (
	"errors"
)

func LIFOEvictionPolicyGet[Key comparable, Val any](key Key, cache *Cache[Key, Val]) (any, error) {
	if elem, ok := cache.Data[key]; ok {
		return elem.Value.(*CacheEntry[Key, Val]).Val, nil
	}
	return nil, errors.New("Could not retrieve value from the cache")
}

func LIFOEvictionPolicyPut[Key comparable, Val any](key Key, value Val, cache *Cache[Key, Val]) error {
	if len(cache.Data)+1 > cache.Capacity {
		// Evict the least recently used element

		tail := cache.Els.Front()

		//remove from Cache map
		delete(cache.Data, tail.Value.(*CacheEntry[Key, Val]).Key)

		//remove from linkedlist
		cache.Els.Remove(tail)
	}

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

	return nil
}
