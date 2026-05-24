package mediator

// User defines the interface for participants in the communication
type User interface {
	SetMediator(mediator Mediator)
	GetName() string
	ReceiveMessage(message string, from User)
}
