//in go, Inheritance type pattern acheived through composition; struct embedding and interfaces.

package main

import(
	"fmt"
	"time"
)

//Defining struct
type Notification struct{
	Recipient string
	Message string
	Timestamp string
}


//New Object creation
func NewNotification(recipient string, message string) Notification {
	return Notification{
		Recipient: recipient,
		Message: message,
		Timestamp: time.Now().Format("2001-09-22 22:02:03"),
	}
}

//Pointer receiver
func (n *Notification) FormatHeader() string {
	return "["+n.Timestamp+" To :"+n.Recipient+"]"
}

// EmailNofication_ embeds Notification (methods like FormatHearder are promoted)
type EmailNotification_ struct{
	Notification
	Subject string
}


func (e *EmailNotification_) Send(){
    fmt.Println(e.FormatHeader()) //method promoted from Notification
    fmt.Printf("Subject: %s\n", e.Subject)
    fmt.Printf("Body: %s\n", e.Message)
    fmt.Println("Status: Email delivered")
}



// SMSNotification_ embeds Notification
type SMSNotification_ struct {
    Notification
    PhoneNumber string
}

func (s *SMSNotification_) Send() {
    fmt.Println(s.FormatHeader())
    fmt.Printf("Phone: %s\n", s.PhoneNumber)
    fmt.Printf("SMS: %s\n", s.Message)
    fmt.Printf("Status: SMS sent (%d chars)\n", len(s.Message))
}

// PushNotification_ embeds Notification
type PushNotification_ struct {
    Notification
    DeviceToken string
    Priority    string
}

func (p *PushNotification_) Send() {
    fmt.Println(p.FormatHeader())
    fmt.Printf("Device: %s...\n", p.DeviceToken[:8])
    fmt.Printf("Priority: %s\n", p.Priority)
    fmt.Printf("Alert: %s\n", p.Message)
    fmt.Println("Status: Push notification delivered")
}


// func main(){

// 	var email *EmailNotification_ = &EmailNotification_{
//         Notification: NewNotification("alice@example.com", "Your order has been shipped!"),
//         Subject:      "Order Update",
//     }

// 	email.Send()
// 	fmt.Println()


//     sms := &SMSNotification_{
//         Notification: NewNotification("Bob", "Your verification code is 482910."),
//         PhoneNumber:  "+1-555-0123",
//     }
//     sms.Send()

//     fmt.Println()

//     push := &PushNotification_{
//         Notification: NewNotification("Charlie", "New message from Alice"),
//         DeviceToken:  "d8a3f4b2c1e5a9b7",
//         Priority:     "high",
//     }
//     push.Send()
// }