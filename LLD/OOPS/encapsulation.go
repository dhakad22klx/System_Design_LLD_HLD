package main

import "fmt"

type PaymentProcessor struct {
    cardNumber string //first letter is lower case, so it is private to a package.
    amount     float64
}

func maskCardNumber(cardNumber string) string { ///first letter is lower case, so function is private to a package.
    return "****-****-****-" + cardNumber[len(cardNumber)-4:]
}

func NewPaymentProcessor(cardNumber string, amount float64) *PaymentProcessor {
    return &PaymentProcessor{
        cardNumber: maskCardNumber(cardNumber),
        amount:     amount,
    }
}

func (p *PaymentProcessor) ProcessPayment() {
    fmt.Printf("Processing payment of $%.2f for card %s\n", p.amount, p.cardNumber)
}

// func main() {
//     payment := NewPaymentProcessor("1234567812345678", 250.00)
//     payment.ProcessPayment()
// }