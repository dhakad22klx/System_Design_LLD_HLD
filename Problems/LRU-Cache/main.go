package main

import "fmt"

func main() {
	cache := NewLRUCache[string, int](3)

	cache.PutKey("a", 1)
	cache.PutKey("b", 2)
	cache.PutKey("c", 3)

	if value, ok := cache.GetKey("a"); ok {
		fmt.Printf("Value %d, found %t\n", value, ok)
	}

	cache.PutKey("d", 4)

	fmt.Println(cache.GetKey("b")) // 0 (zero value for int when missing)
}
