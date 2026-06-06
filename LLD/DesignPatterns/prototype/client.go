package prototype

import "fmt"

/*
The Prototype Design Pattern is a creational design pattern that lets you create new objects by cloning existing ones,
instead of instantiating them from scratch.

It’s particularly useful in situations where:

- Creating a new object is expensive, time-consuming, or resource-intensive.
- You want to avoid duplicating complex initialization logic.
- You need many similar objects with only slight differences.

The Prototype Pattern allows you to create new instances by cloning a pre-configured prototype object,
ensuring consistency while reducing boilerplate and complexity.

---
Prototype is a creational design pattern that lets you copy existing objects
without making your code dependent on their classes.

Say you have an object, and you want to create an exact copy of it. How would you do it?
First, you have to create a new object of the same class.
Then you have to go through all the fields of the original object and copy their values over to the new object.

Nice! But there’s a catch. Not all objects can be copied that way because some of the object’s fields may be private
and not visible from outside of the object itself.

----
Prototype is a creational design pattern that allows cloning objects, even complex ones,
without coupling to their specific classes.

All prototype classes should have a common interface that makes it possible to copy objects even if
their concrete classes are unknown. Prototype objects can produce full copies
since objects of the same class can access each other’s private fields.
*/
func TestPrototype() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
