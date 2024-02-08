package inmemocache

import "fmt"

func main() {
	cache := NewCache(5, LRUEvictionPolicyGet[string, string], LRUEvictionPolicyPut[string, string])

	cache.Put("hello", "jj")
	val, err := cache.Get("hello")

	if err == nil {
		fmt.Println(val)
	}
}
