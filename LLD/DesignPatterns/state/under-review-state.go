package state

import (
	"fmt"
)

// concrete state
type UnderReviewState struct{}

func (s *UnderReviewState) edit(doc *Document, content string) {
	fmt.Println("Cannot edit while under review.")
}

func (s *UnderReviewState) submitForReview(doc *Document) {
	fmt.Println("Document is already under review.")
}

func (s *UnderReviewState) approve(doc *Document) {
	fmt.Println("Document approved and published.")
	doc.setState(&PublishedState{})
}

func (s *UnderReviewState) reject(doc *Document) {
	fmt.Println("Document rejected. Returning to draft.")
	doc.setState(&DraftState{})
}

func (s *UnderReviewState) unpublish(doc *Document) {
	fmt.Println("Document is not published yet.")
}
