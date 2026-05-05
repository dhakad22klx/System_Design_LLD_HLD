package observer

/*
Observer is a behavioral design pattern that allows some objects to
notify other objects about changes in their state.

The Observer pattern provides a way to subscribe and unsubscribe to and from
these events for any object that implements a subscriber interface.
---
Observer is a behavioral design pattern that lets you define a subscription mechanism to notify
multiple objects about any events that happen to the object they’re observing.
*/

func TestObserverPattern() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
