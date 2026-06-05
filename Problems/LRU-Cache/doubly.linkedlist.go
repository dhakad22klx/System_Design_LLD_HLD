package main

type DoublyLinkedList[K comparable, V any] struct {
	head *Node[K, V]
	tail *Node[K, V]
}

func newDoublyLinkedList[K comparable, V any]() *DoublyLinkedList[K, V] {

	// Create dummy head and tail nodes with zero key/value.
	var zeroK K
	var zeroV V
	head := newNode(zeroK, zeroV)
	tail := newNode(zeroK, zeroV)

	head.next = tail
	tail.prev = head

	return &DoublyLinkedList[K, V]{head: head, tail: tail}
}

func (list *DoublyLinkedList[K, V]) AddFirst(node *Node[K, V]) {

	// Insert node right after head(dummy head)

	headNext := list.head.next
	node.next = headNext
	node.prev = list.head
	headNext.prev = node
	list.head.next = node
}

func (list *DoublyLinkedList[K, V]) Remove(node *Node[K, V]) {
	if node == nil {
		return
	}

	left := node.prev
	right := node.next

	if left == nil || right == nil {
		return
	}

	left.next = right
	right.prev = left

	node.next = nil
	node.prev = nil

}

func (list *DoublyLinkedList[K, V]) MoveToFront(node *Node[K, V]) {
	// Move an existing node to the front.
	// Remove it first, then add it to front.
	list.Remove(node)
	list.AddFirst(node)
}

func (list *DoublyLinkedList[K, V]) RemoveLast() *Node[K, V] {

	if list.tail.prev == list.head {
		return nil
	}

	last := list.tail.prev

	list.Remove(last)

	return last
}
