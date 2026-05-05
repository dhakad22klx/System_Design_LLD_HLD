package observer

// Observer Interface
type Observer interface {
	update(string)
	getID() string
}
