package strategy

/*
Strategy is a behavioral design pattern that turns a set of behaviors into objects
and makes them interchangeable inside original context object.

The original object, called context, holds a reference to a strategy object.
The context delegates executing the behavior to the linked strategy object.
In order to change the way the context performs its work,
other objects may replace the currently linked strategy object with another one.


----

The Strategy Design Pattern is a behavioral pattern that lets you define a family of algorithms,
encapsulate each one in its own class, and make them interchangeable at runtime.

---

Instead of embedding multiple algorithms inside a single class with conditional logic,
you extract each algorithm into its own strategy class. The main class (context) delegates the work to whichever strategy is currently plugged in.

This pattern becomes valuable when:

- You have multiple ways to perform the same operation, and the choice might change at runtime
- You want to avoid bloated conditional statements that select between different behaviors
- You need to isolate algorithm-specific data and logic from the code that uses it
- Different clients might need different algorithms for the same task

*/

func TestStrategyPattern() {
	cart := NewShoppingCart(100.0)

	// Use credit card payment
	creditCard := NewCreditCardPayment("1234-5678-9012-3456", "John Doe", "123", "12/25")
	cart.SetPaymentStrategy(creditCard)
	cart.Checkout()

	// Use PayPal payment
	paypal := NewPayPalPayment("john@example.com")
	cart.SetPaymentStrategy(paypal)
	cart.Checkout()
}
