package state

// DocumentState that defines interface for all the document states.
type DocumentState interface {
	edit(context *Document, content string)
	submitForReview(context *Document)
	approve(context *Document)
	reject(context *Document)
	unpublish(context *Document)
}
