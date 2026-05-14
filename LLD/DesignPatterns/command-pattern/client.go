package commandpattern

/*
Command is behavioral design pattern that converts requests or simple operations into objects.

The conversion allows deferred or remote execution of commands, storing command history, etc.

------------

Command is a behavioral design pattern that turns a request into a stand-alone object that contains all
information about the request. This transformation lets you pass requests as a method arguments,
delay or queue a request’s execution, and support undoable operations.

-----

It’s particularly useful in situations where:

You want to encapsulate operations as objects.
You need to queue, delay, or log requests.
You want to support undo/redo functionality.
You want to decouple the object that invokes an operation from the one that knows how to perform it.
*/

import "fmt"

// ── TextEditor───

type TextEditor struct {
	content string
}

func (e *TextEditor) append(text string) {
	e.content += text
}

// deleteLast removes up to `count` characters from the end and returns what was removed.
func (e *TextEditor) deleteLast(count int) string {
	start := len(e.content) - count
	if start < 0 {
		start = 0
	}
	deleted := e.content[start:]
	e.content = e.content[:start]
	return deleted
}

func (e *TextEditor) getContent() string { return e.content }

type EditorCommand interface {
	execute()
	undo()
}

// ── TypeCommand

type TypeCommand struct {
	editor *TextEditor
	text   string
}

func (c *TypeCommand) execute() {
	c.editor.append(c.text)
	fmt.Printf("Typed: %q\n", c.text)
}

func (c *TypeCommand) undo() {
	c.editor.deleteLast(len(c.text))
	fmt.Printf("Undo type: %q\n", c.text)
}

// ── DeleteCommand

type DeleteCommand struct {
	editor      *TextEditor
	count       int
	deletedText string
}

func (c *DeleteCommand) execute() {
	c.deletedText = c.editor.deleteLast(c.count)
	fmt.Printf("Deleted: %q\n", c.deletedText)
}

func (c *DeleteCommand) undo() {
	c.editor.append(c.deletedText)
	fmt.Printf("Undo delete: restored %q\n", c.deletedText)
}

// ── EditorInvoker
// EditorInvoker manages undo/redo stacks using slices as stacks.
// Go has no built-in stack; a slice with append/pop is idiomatic.
type EditorInvoker struct {
	undoStack []EditorCommand
	redoStack []EditorCommand
}

func (inv *EditorInvoker) execute(cmd EditorCommand) {
	cmd.execute()
	inv.undoStack = append(inv.undoStack, cmd)
	inv.redoStack = inv.redoStack[:0] // clear redo stack
}

func (inv *EditorInvoker) undo() {
	n := len(inv.undoStack)
	if n == 0 {
		fmt.Println("Nothing to undo.")
		return
	}
	cmd := inv.undoStack[n-1]
	inv.undoStack = inv.undoStack[:n-1]
	cmd.undo()
	inv.redoStack = append(inv.redoStack, cmd)
}

func (inv *EditorInvoker) redo() {
	n := len(inv.redoStack)
	if n == 0 {
		fmt.Println("Nothing to redo.")
		return
	}
	cmd := inv.redoStack[n-1]
	inv.redoStack = inv.redoStack[:n-1]
	cmd.execute()
	inv.undoStack = append(inv.undoStack, cmd)
}

func TestCommandPattern() {
	editor := &TextEditor{}
	invoker := &EditorInvoker{}

	type1 := &TypeCommand{editor: editor, text: "Hello"}
	type2 := &TypeCommand{editor: editor, text: " World"}
	type3 := &TypeCommand{editor: editor, text: "!"}

	invoker.execute(type1)
	invoker.execute(type2)
	invoker.execute(type3)
	fmt.Printf("Content: %q\n", editor.getContent())

	fmt.Println("\n--- Undo ---")
	invoker.undo()
	fmt.Printf("Content: %q\n", editor.getContent())
	invoker.undo()
	fmt.Printf("Content: %q\n", editor.getContent())

	fmt.Println("\n--- Redo ---")
	invoker.redo()
	fmt.Printf("Content: %q\n", editor.getContent())

	fmt.Println("\n--- New operation clears redo ---")
	del := &DeleteCommand{editor: editor, count: 3}
	invoker.execute(del)
	fmt.Printf("Content: %q\n", editor.getContent())
	invoker.redo()

	fmt.Println("\n--- Undo delete ---")
	invoker.undo()
	fmt.Printf("Content: %q\n", editor.getContent())
}
