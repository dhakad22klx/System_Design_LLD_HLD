package state

/*
The State Design Pattern is a behavioral design pattern that lets an object change its behavior
when its internal state changes, as if it were switching to a different class at runtime.
----
State is a behavioral design pattern that lets an object alter its behavior when its internal state changes.
It appears as if the object changed its class

The pattern extracts state-related behaviors into separate state classes and
forces the original object to delegate the work to an instance of these classes, instead of acting on its own.
----
It’s particularly useful in situations where:

An object can be in one of many distinct states, each with different behavior.
The object’s behavior depends on current context, and that context changes over time.
You want to avoid large, monolithic if-else or switch statements that check for every possible state.
*/
func TestStatePattern() {
	doc := NewDocument()

	doc.Edit("First draft of the article.")
	doc.Approve() // Cannot approve a draft
	doc.SubmitForReview()
	doc.Edit("Trying to edit") // Cannot edit while under review
	doc.Reject()               // Back to draft
	doc.Edit("Revised draft.")
	doc.SubmitForReview()
	doc.Approve()              // Published
	doc.Edit("Trying to edit") // Cannot edit a published document
	doc.Unpublish()            // Back to draft
}
