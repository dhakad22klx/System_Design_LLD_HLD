package main

import "sync"

type LRUCache[K comparable, V any] struct {
	capacity    int
	keyValueMap map[K]*Node[K, V]
	list        *DoublyLinkedList[K, V]
	mu          sync.Mutex
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	// Initialize capacity, create empty map, create new DoublyLinkedList.
	return &LRUCache[K, V]{
		capacity:    capacity,
		keyValueMap: make(map[K]*Node[K, V]),
		list:        newDoublyLinkedList[K, V](),
	}
}

func (cache *LRUCache[K, V]) GetKey(key K) (V, bool) {
	//Thread safety
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// Check if key exist (if exist then return value, true else zeroV, false)
	node, ok := cache.keyValueMap[key]
	if !ok {
		var zeroV V
		return zeroV, false
	}

	cache.list.MoveToFront(node)

	return node.value, true

}

func (cache *LRUCache[K, V]) PutKey(key K, value V) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// If key already exist : Just update the value and move to the first.
	if node, ok := cache.keyValueMap[key]; ok {
		node.value = value
		cache.list.MoveToFront(node)
		return
	}

	//Case : key is not present

	//check for capacity : if capacity is full then evict else just put.
	if len(cache.keyValueMap) == cache.capacity {
		//evict the last record and delete from the keyvalue map
		if last := cache.list.RemoveLast(); last != nil {
			delete(cache.keyValueMap, last.key)
		}
	}
	//add this key to first
	newNode := newNode(key, value)
	cache.list.AddFirst(newNode)
	cache.keyValueMap[key] = newNode
}

func (cache *LRUCache[K, V]) Remove(key K) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	node, ok := cache.keyValueMap[key]
	if !ok {
		return
	}

	cache.list.Remove(node)
	delete(cache.keyValueMap, key)
}
