package observer

// Abstract Subject
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
