# Installation

```bash
go get github.com/Ayman161803/inmemocache
```

# Features
- Provides inbuilt standard cache policies - LRU, FIFO, LIFO
- Provides an interface to give customized get/put policies to the cache.
- Thread safety has been implemented by read and write mutex(Look [here](https://github.com/Ayman161803/inmemocache/blob/80849a9ebdcc2606dd35eaddb78f0e2e3c6ea254/inmemocache.go#L45)). This ensures that Read-Write and Write-Write conflicts are prevented. 
    
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

