package memento

/*
The Memento Design Pattern is a behavioral design pattern that lets you capture and
store an object’s internal state so it can be restored later, without violating encapsulation.

It’s particularly useful in situations where:

-You need to implement undo/redo functionality.
-You want to support checkpointing or versioning of an object’s state.
-You want to separate the concerns of state storage from state management logic.

---
Memento is a behavioral design pattern that lets you save and
restore the previous state of an object without revealing the details of its implementation.
---
Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in future.

The Memento doesn’t compromise the internal structure of the object it works with,
as well as data kept inside the snapshots.

# Memento Vs Command Pattern

Command thinks in ACTIONS:

	"Insert 'Hello' was done → to undo, delete 5 chars"
	"Delete 6 chars was done → to undo, insert ' World'"

	Undo logic lives INSIDE the command object.
	Only stores what's relevant to reverse THIS action.

Memento thinks in SNAPSHOTS:

	"Before you made changes — here's the full photograph"
	"Restore? Here, put everything back to this photograph"

	No undo logic anywhere — just save and restore.
	Stores the ENTIRE object state regardless of what changed.

# Where Command wins, where Memento wins

Command is better when:

Actions are discrete and well-defined — insert, delete, move
You want a detailed audit log — "who did what when"
You want redo as well as undo — commands can replay forward
State is large — storing full snapshots would be expensive

Memento is better when:

Undo logic is too complex to reverse algorithmically
State is interconnected — many fields depend on each other
You don't want the undo mechanism to know about internals
You need to jump to arbitrary points — not just one step back

Complex state rollback :
Hard in Command— undo logic gets messy
Easy in memento— just restore snapshot

Memory cost is high in memento.
*/
func TestMementoPattern() {

	doc := &Document{content: "Hello", cursorPos: 5, fontSize: 12}
	history := &History{} //snapshot manager

	// snapshot 1
	history.Push(doc.Save())
	doc.Print()

	// make changes
	doc.content = "Hello World"
	doc.cursorPos = 11
	doc.fontSize = 14

	// snapshot 2
	history.Push(doc.Save())
	doc.Print()

	// drastic change
	doc.content = "Oops everything is wrong"
	doc.fontSize = 99
	doc.Print()

	// restore to snapshot 2
	doc.Restore(history.Pop())
	doc.Print()

	// restore to snapshot 1
	doc.Restore(history.Pop())
	doc.Print()
}
