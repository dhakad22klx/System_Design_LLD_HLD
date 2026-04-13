package factorymethod

// Creator Interface
type NotificationCreator interface {
	CreateNotification() Notification // The Factory Method
}

// We can define a base function to represent shared logic
// that uses the factory method.
func SendNotification(f NotificationCreator, message string) {
	n := f.CreateNotification()
	n.Send(message)
}

type EmailCreator struct{}

func (ec *EmailCreator) CreateNotification() Notification {
	return &EmailNotification{}
}

type SMSCreator struct{}

func (sc *SMSCreator) CreateNotification() Notification {
	return &SMSNotification{}
}
