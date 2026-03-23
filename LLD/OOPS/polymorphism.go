package main

import "fmt"

// Go achieves runtime polymorphism through interfaces, not inheritance.
// Any type that has a Send() method satisfies the Notifier interface.

type Notifier interface {
    Send()
}

type EmailNotification struct {
    Recipient string
    Message   string
    Subject   string
}

func (e *EmailNotification) Send() {
    fmt.Printf("Sending EMAIL to %s | Subject: %s\n", e.Recipient, e.Subject)
}

type SMSNotification struct {
    Recipient   string
    Message     string
    PhoneNumber string
}

func (s *SMSNotification) Send() {
    fmt.Printf("Sending SMS to %s | Message: %s\n", s.PhoneNumber, s.Message)
}

type PushNotification struct {
    Recipient   string
    Message     string
    DeviceToken string
}

func (p *PushNotification) Send() {
    fmt.Printf("Sending PUSH to device %s... | Alert: %s\n",
        p.DeviceToken[:8], p.Message)
}

// func main() {
//     notifications := []Notifier{
//         &EmailNotification{"alice@example.com", "Your order shipped!", "Order Update"},
//         &SMSNotification{"Bob", "Code: 482910", "+1-555-0123"},
//         &PushNotification{"Charlie", "New message", "d8a3f4b2c1e5a9b7"},
//     }

//     for _, n := range notifications {
//         n.Send()
//     }
// }