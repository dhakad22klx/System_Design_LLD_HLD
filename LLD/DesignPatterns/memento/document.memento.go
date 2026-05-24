package memento

type State struct {
	content   string
	cursorPos int
	fontSize  int
}

// Memento — the snapshot. Immutable after creation.
// complete state captured
type DocumentMemento struct {
	state State
}

func NewDocumentMementoState(d *Document) *DocumentMemento {
	return &DocumentMemento{
		state: State{
			content:   d.content,
			cursorPos: d.cursorPos,
			fontSize:  d.fontSize,
		},
	}
}
