package memento

// Caretaker — manages the snapshots, knows WHEN to save/restore
// but never looks INSIDE the snapshot

type History struct {
	snapshots []*DocumentMemento
}

func (h *History) Push(m *DocumentMemento) {
	h.snapshots = append(h.snapshots, m)
}

func (h *History) Pop() *DocumentMemento {
	if len(h.snapshots) == 0 {
		return nil
	}
	last := h.snapshots[len(h.snapshots)-1]
	h.snapshots = h.snapshots[:len(h.snapshots)-1]
	return last
}
