package mediator

/*
The Mediator Design Pattern is a behavioral pattern that defines
an object (the Mediator) to encapsulate how a set of objects interact.

Mediator is a behavioral design pattern that reduces coupling between components of a program
by making them communicate indirectly, through a special mediator object.

The Mediator makes it easy to modify, extend and reuse individual components because
they’re no longer dependent on the dozens of other classes.

Observer Vs Mediator
Observer — One source, many listeners.
Mediator — Many components, one coordinator.


When to use which
Observer — when an event happens and multiple things need to react independently:

Order placed → email + SMS + inventory update
User logged in → analytics + session + audit log
Price changed → charts + alerts + history

Mediator — when multiple components need to coordinate with each other and the coordination logic is complex:

Chat room — messages routed between users
Air traffic control — planes coordinating runway access
UI form — field A changes → field B gets disabled → field C resets
Workflow engine — step A completes → decide whether to run B or C
*/

func TestMediatorPattern() {
	mediator := NewChatMediator()

	alice := NewUser("Alice")
	bob := NewUser("Bob")
	carol := NewUser("Carol")

	mediator.addUser(alice)
	mediator.addUser(bob)
	mediator.addUser(carol)

	alice.SendMessage("Hello, everyone!")
	bob.SendMessage("Hi Alice!")
	carol.SendMessage("Hey folks!")
}
