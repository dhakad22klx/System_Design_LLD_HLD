package mediator

import "fmt"

// User represents a concrete colleague
type ChatUser struct {
	name     string
	mediator Mediator
}

// NewUser creates a new User
func NewUser(name string) *ChatUser {
	return &ChatUser{name: name}
}

func (u *ChatUser) SetMediator(mediator Mediator) {
	u.mediator = mediator
}

func (u *ChatUser) GetName() string {
	return u.name
}

func (u *ChatUser) ReceiveMessage(message string, from User) {
	fmt.Printf("%s received: %s from %s\n", u.name, message, from.GetName())
}

// SendMessage allows the User to send a message via the mediator
func (u *ChatUser) SendMessage(message string) {
	if u.mediator != nil {
		u.mediator.SendMessage(message, u)
	}
}
