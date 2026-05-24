package memento

import "fmt"

// Originator — the object whose state we snapshot
type Document struct {
	content   string
	cursorPos int
	fontSize  int
}

// Save — creates a snapshot of current state
func (d *Document) Save() *DocumentMemento {
	return NewDocumentMementoState(d)
}

// Restore — goes back to a previous snapshot
func (d *Document) Restore(m *DocumentMemento) {
	fmt.Println("Restoring to previous state")
	// fmt.Printf("Document memento : content %s , cursorPos %d, fontSize %d", m.state.content, m.state.cursorPos, m.state.fontSize)
	// fmt.Println()
	d.content = m.state.content
	d.cursorPos = m.state.cursorPos
	d.fontSize = m.state.fontSize
}

func (d *Document) Print() {
	fmt.Printf("Content: %q | Cursor: %d | Font: %d\n",
		d.content, d.cursorPos, d.fontSize)
}
