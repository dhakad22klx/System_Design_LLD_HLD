package main

import "fmt"

type Notifier interface {
	Send()
}

/*
Concrete implementations
*/

type EmailNotification struct {
	Recipient string
	Message   string
	Subject   string
}

func (e *EmailNotification) Send() {
	fmt.Printf("Sending EMAIL to %s | Subject: %s\n", e.Recipient, e.Subject)
}

type SMSNotification struct {
	PhoneNumber string
	Message     string
}

func (s *SMSNotification) Send() {
	fmt.Printf("Sending SMS to %s | Message: %s\n", s.PhoneNumber, s.Message)
}

/*
Service that depends on abstraction
*/
type NotificationService struct {
	notifier Notifier
}

/*
Dependency Injection via constructor
*/
func NewNotificationService(n Notifier) *NotificationService {
	return &NotificationService{
		notifier: n,
	}
}

/*
Business logic
*/
func (s *NotificationService) Notify() {
	s.notifier.Send()
}

func testDependencyInjection() {

	email := &EmailNotification{
		Recipient: "alice@email.com",
		Message:   "Order shipped",
		Subject:   "Order Update",
	}

	sms := &SMSNotification{
		PhoneNumber: "+91XXXXXX",
		Message:     "OTP 1234",
	}

	emailService := NewNotificationService(email)
	smsService := NewNotificationService(sms)

	emailService.Notify()
	smsService.Notify()
}