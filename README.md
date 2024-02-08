# Installation

```bash
go get github.com/Ayman161803/inmemocache
```
# Usage

## LRU Cache Example
```go
package main

import (
	"fmt"

	. "github.com/Ayman161803/inmemocache"
)

func main() {

	cache := NewCache(3, LRUEvictionPolicyGet[string, string], LRUEvictionPolicyPut[string, string])

	cache.Put("hello", "jj")
	cache.Put("kk", "jjkk")
	cache.Put("kkk", "jjkk")

	cache.Get("hello")
	cache.Put("hell", "jj")

	//get val for 'kk'. Now since 'kk' was least recently used element, it was evicted when 'hell' key was entered into cache
	val, err := cache.Get("kk")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'kk' is", val)
	}

	//value from key 'hello' should get retrieved successfully
	val, err = cache.Get("hello")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'hello' is", val)
	}
}
```

## LIFO Cache Example

```go
package main

import (
	"fmt"

	. "github.com/Ayman161803/inmemocache"
)

func main() {

	cache := NewCache(3, LIFOEvictionPolicyGet[string, string], LIFOEvictionPolicyPut[string, string])

	cache.Put("hello", "jj")
	cache.Put("kk", "jjkk")
	cache.Put("kkk", "jjkk")

	cache.Get("hello")
	cache.Put("hell", "jj")

	//  'kkk' is the last entry in cache. hence, it is evicted when 'hell' is added.
	// We have a cache hit for 'hello' and a cache miss for 'kkk'
	val, err := cache.Get("kkk")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'kkk' is", val)
	}

	val, err = cache.Get("hello")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'hello' is", val)
	}
}
```

## FIFO Cache Example

```go
package main

import (
	"fmt"

	. "github.com/Ayman161803/inmemocache"
)

func main() {

	cache := NewCache(3, FIFOEvictionPolicyGet[string, string], FIFOEvictionPolicyPut[string, string])

	cache.Put("hello", "jj")
	cache.Put("kk", "jjkk")
	cache.Put("kkk", "jjkk")

	cache.Get("hello")
	cache.Put("hell", "jj")

	//get val for 'kk'. Even though 'hello' key was recentky used, the cache entry for 'hello' is oldest.
	//Hence, we have a cache miss for 'hello' and a hit for 'kk'
	val, err := cache.Get("kk")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'kk' is", val)
	}

	val, err = cache.Get("hello")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'hello' is", val)
	}
}
```

## Custum Get/Put Policy
The example policy below always returns "Unga Bunga" for a get and throws an error for a put
 ```go
 package main

import (
	"errors"
	"fmt"

	. "github.com/Ayman161803/inmemocache"
)

func CustomGetPolicy[Key comparable, Val any](key Key, cache *Cache[Key, Val]) (any, error) {
	return "Unga Bunga", nil
}

func CustomPutPolicy[Key comparable, Val any](key Key, val Val, cache *Cache[Key, Val]) error {
	return errors.New("I always throw errors")
}

func main() {

	cache := NewCache(3, CustomGetPolicy[string, string], CustomPutPolicy[string, string])

	// get always returns "Unga Bunga"

	val, err := cache.Get("kk")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	} else {
		fmt.Println("Value for key 'kk' is", val)
	}

	// Put always throws an error
	err = cache.Put("hello", "hi?")
	if err != nil {
		fmt.Println("Cache miss error: ", err)
	}
}
```
