package state

// Context for state pattern : Holds current state and content
type Document struct {
	state   DocumentState
	content string
}


func NewDocument() *Document {
	return &Document{state: &DraftState{}}
}
 
func (d *Document) setState(s DocumentState) { d.state = s }
func (d *Document) setContent(c string)      { d.content = c }
func (d *Document) Content() string          { return d.content }
 
func (d *Document) Edit(content string)  { d.state.edit(d, content) }
func (d *Document) SubmitForReview()     { d.state.submitForReview(d) }
func (d *Document) Approve()             { d.state.approve(d) }
func (d *Document) Reject()              { d.state.reject(d) }
func (d *Document) Unpublish()           { d.state.unpublish(d) }