package mediator

// ChatMediator is a concrete mediator that manages users
type ChatMediator struct {
	users []User
}

// NewChatMediator creates a new ChatMediator
func NewChatMediator() *ChatMediator {
	return &ChatMediator{users: make([]User, 0)}
}

// AddUser adds a user to the chat
func (m *ChatMediator) addUser(user User) {
	m.users = append(m.users, user)
	user.SetMediator(m)
}

// SendMessage sends a message from one user to all others
func (m *ChatMediator) SendMessage(message string, sender User) {
	for _, user := range m.users {
		if user != sender {
			user.ReceiveMessage(message, sender)
		}
	}
}
