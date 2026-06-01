package entities

import (
	"fmt"
)

type QuestionClosedState struct{}

func (state *QuestionClosedState) addAnswer(question *Question, answer *Answer) {
	fmt.Printf("Answer can not be added to closed question : %s.\n", question.GetID())
}

func (state *QuestionClosedState) acceptAnswer(question *Question, answer *Answer) {
	fmt.Printf("A answer (id : %s) for this question (id : %s) already accepted.\n", question.acceptedAnswer.GetID(), question.GetID())
}
