package adapter

import (
	"fmt"
)

func TestAdapterPattern() {
	// Using the in-house payment processor
	fmt.Println("Using In-House Payment Processor:")
	inHouseProcessor := &InHousePaymentProcessor{}
	checkoutService := NewCheckoutService(inHouseProcessor)
	checkoutService.ProcessCheckout(100.0, "USD")

	fmt.Println("\nUsing Legacy Gateway (via Adapter):")
	// Using the legacy gateway through the adapter
	legacyGateway := &LegacyGateway{}
	legacyAdapter := NewLegacyGatewayAdapter(legacyGateway)
	checkoutService = NewCheckoutService(legacyAdapter)
	checkoutService.ProcessCheckout(150.0, "EUR")
}
