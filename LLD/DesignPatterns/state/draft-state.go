package state

import "fmt"

type DraftState struct{}

func (s *DraftState) edit(doc *Document, content string) {
	fmt.Println("Editing document:", content)
	doc.setContent(content)
}

func (s *DraftState) submitForReview(doc *Document) {
	fmt.Println("Document submitted for review.")
	doc.setState(&UnderReviewState{})
}

func (s *DraftState) approve(doc *Document) {
	fmt.Println("Cannot approve a draft. Submit for review first.")
}

func (s *DraftState) reject(doc *Document) {
	fmt.Println("Cannot reject a draft. Submit for review first.")
}

func (s *DraftState) unpublish(doc *Document) {
	fmt.Println("Document is already a draft.")
}
