package entities

type IQuestionState interface {
	addAnswer(question *Question, answer *Answer)
	acceptAnswer(question *Question, answer *Answer)
}
