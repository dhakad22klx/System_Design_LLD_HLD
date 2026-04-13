package factorymethod

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (e *EmailNotification) Send(msg string) {
	fmt.Printf("Sending Email: %s\n", msg)
}

type SMSNotification struct{}

func (s *SMSNotification) Send(msg string) {
	fmt.Printf("Sending SMS: %s\n", msg)
}
