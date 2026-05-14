package state

import "fmt"

type PublishedState struct{}

func (s *PublishedState) edit(doc *Document, content string) {
	fmt.Println("Cannot edit a published document. Unpublish first.")
}

func (s *PublishedState) submitForReview(doc *Document) {
	fmt.Println("Document is already published.")
}

func (s *PublishedState) approve(doc *Document) {
	fmt.Println("Document is already published.")
}

func (s *PublishedState) reject(doc *Document) {
	fmt.Println("Cannot reject a published document.")
}

func (s *PublishedState) unpublish(doc *Document) {
	fmt.Println("Document unpublished. Returning to draft.")
	doc.setState(&DraftState{})
}
