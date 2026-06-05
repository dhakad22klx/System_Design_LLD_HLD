package main

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

func newNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{key: key, value: value}
}
