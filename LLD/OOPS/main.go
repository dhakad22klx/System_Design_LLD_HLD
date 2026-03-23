package main



import (
	"fmt"
	
)


func main(){

	//Calling Abstraction
	fmt.Println("Calling Abstraction")
	var dbClient *DatabaseClient = NewDatabaseClient(10, 3)
	dbClient.Connect("localhost", 8080)


	//Calling Encapsulation 
	fmt.Println()
	fmt.Println("Calling Encapsulation")
	payment := NewPaymentProcessor("1234567812345678", 250.00)
	payment.ProcessPayment()

	//Calling Inheritance
	fmt.Println()
	fmt.Println("Calling Inheritance")
	var email *EmailNotification_ = &EmailNotification_{
        Notification: NewNotification("alice@example.com", "Your order has been shipped!"),
        Subject:      "Order Update",
    }

	email.Send()
	fmt.Println()


    sms := &SMSNotification_{
        Notification: NewNotification("Bob", "Your verification code is 482910."),
        PhoneNumber:  "+1-555-0123",
    }
    sms.Send()

    fmt.Println()

    push := &PushNotification_{
        Notification: NewNotification("Charlie", "New message from Alice"),
        DeviceToken:  "d8a3f4b2c1e5a9b7",
        Priority:     "high",
    }
    push.Send()

	//Calling Polymorphism
	fmt.Println()
	fmt.Println("Calling Polymorphism")
    notifications := []Notifier{
        &EmailNotification{"alice@example.com", "Your order shipped!", "Order Update"},
        &SMSNotification{"Bob", "Code: 482910", "+1-555-0123"},
        &PushNotification{"Charlie", "New message", "d8a3f4b2c1e5a9b7"},
    }

    for _, n := range notifications {
        n.Send()
    }
}