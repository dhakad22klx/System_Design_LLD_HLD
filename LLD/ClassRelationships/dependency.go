// A Dependency exists when one class relies on another to fulfill a responsibility,
// but does so without retaining a permanent reference to it.

package main

import (
	"fmt"
)

// SeatValidator checks seat availability
type SeatValidator struct{}

func (s *SeatValidator) IsAvailable(eventId, seatNumber string) bool {
	fmt.Printf("Checking seat %s for event %s\n", seatNumber, eventId)
	return true
}

// PaymentProcessor handles charging
type PaymentProcessor struct{}

func (p *PaymentProcessor) Charge(email string, amount float64) bool {
	fmt.Printf("Charging $%.2f to %s\n", amount, email)
	return true
}

// QRCodeGenerator generates QR codes
type QRCodeGenerator struct{}

func (q *QRCodeGenerator) Generate(eventId, seatNumber string) string {
	qrCode := "QR-" + eventId + "-" + seatNumber
	fmt.Printf("Generated QR code: %s\n", qrCode)
	return qrCode
}

// EmailService sends confirmation emails
type EmailService struct{}

func (e *EmailService) SendConfirmation(email, qrCode string) {
	fmt.Printf("Sending confirmation to %s with code %s\n", email, qrCode)
}

// TicketBookingService orchestrates the booking flow
type TicketBookingService struct{}

func (t *TicketBookingService) BookTicket(
	eventId, seatNumber, email string,
	amount float64,
	validator *SeatValidator,
	payment *PaymentProcessor,
	qrGenerator *QRCodeGenerator,
	emailService *EmailService,
) bool {
	if !validator.IsAvailable(eventId, seatNumber) {
		fmt.Println("Seat not available.")
		return false
	}
	if !payment.Charge(email, amount) {
		fmt.Println("Payment failed.")
		return false
	}
	qrCode := qrGenerator.Generate(eventId, seatNumber)
	emailService.SendConfirmation(email, qrCode)
	fmt.Println("Booking confirmed!")
	return true
}

func testDependency() {
	bookingService := &TicketBookingService{}

	validator := &SeatValidator{}
	payment := &PaymentProcessor{}
	qrGenerator := &QRCodeGenerator{}
	emailService := &EmailService{}

	bookingService.BookTicket(
		"CONF-2025", "A12", "alice@example.com",
		99.99, validator, payment, qrGenerator, emailService,
	)
}

// Why This Design Works
// All dependencies are method parameters. TicketBookingService has zero fields.
// Every collaborator comes in through bookTicket() and disappears when the method returns.
// This is pure dependency with no structural coupling.
