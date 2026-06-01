package entities

import "stack.overflow/enums"

type QuestionOpenState struct{}

func (state *QuestionOpenState) addAnswer(question *Question, answer *Answer) {
	question.mu.Lock()
	question.answers = append(question.answers, answer)
	question.mu.Unlock()
}

func (state *QuestionOpenState) acceptAnswer(question *Question, answer *Answer) {
	question.mu.Lock()
	defer question.mu.Unlock()

	if question.GetAuthor().GetID() != answer.GetAuthor().GetID() {
		question.acceptedAnswer = answer
		answer.SetAccepted(true)
		question.setState(&QuestionClosedState{})
		question.notifyObservers(NewEvent(enums.ACCEPT_ANSWER, answer.GetAuthor(), &answer.Post))
	}
}
