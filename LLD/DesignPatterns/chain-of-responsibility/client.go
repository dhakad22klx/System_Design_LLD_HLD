package chainofresponsibility

import "fmt"

/*
The Chain of Responsibility Design Pattern is a behavioral pattern that
lets you pass requests along a chain of handlers, allowing each handler to decide
whether to process the request or pass it to the next handler in the chain.

This pattern is useful when:

- A request must be handled by one of many possible handlers,
  and you don’t want the sender to be tightly coupled to any specific one.
- You want to decouple request logic from the code that processes it.
- You want to flexibly add, remove, or reorder handlers without changing the client code.
*/

func TestChainOfResponsibility() {
	// Build handlers
	hundreds := NewHundredDollarHandler()
	fifties := NewFiftyDollarHandler()
	twenties := NewTwentyDollarHandler()
	tens := NewTenDollarHandler()

	// Wire the chain
	hundreds.setNext(fifties)
	fifties.setNext(twenties)
	twenties.setNext(tens)

	fmt.Println("--- Withdrawing $380 ---")
	request1 := &CashRequest{amount: 380}
	hundreds.dispense(request1)
	fmt.Printf("Remaining: $%d\n", request1.amount)

	fmt.Println("\n--- Withdrawing $275 ---")
	request2 := &CashRequest{amount: 275}
	hundreds.dispense(request2)
	fmt.Printf("Remaining: $%d\n", request2.amount)
}
