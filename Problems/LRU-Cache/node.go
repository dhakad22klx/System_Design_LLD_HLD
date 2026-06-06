package main

// [K comparable, V any] : The Type Parameter List. It says
// "This struct uses two placeholders, K and V, which can be comparable and any type."
type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

func newNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{key: key, value: value}
}
