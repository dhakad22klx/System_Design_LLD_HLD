package iterator

// Abstract Iterator
type Iterator interface {
	hasNext() bool
	getNext() *User
}
