package iterator

/*
Iterator is a behavioral design pattern that lets you traverse elements of a collection
without exposing its underlying representation (list, stack, tree, etc.).

---

Iterator is a behavioral design pattern that allows sequential traversal through a
complex data structure without exposing its internal details.

Thanks to the Iterator, clients can go over elements of different
collections in a similar fashion using a single iterator interface.
*/

import "fmt"

func TestIteratorPattern() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
