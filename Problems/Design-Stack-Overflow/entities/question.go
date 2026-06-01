package entities

import (
	"fmt"
	"sync/atomic"
)

type Question struct {
	Post
	title          string
	tags           []*Tag
	answers        []*Answer
	acceptedAnswer *Answer
	questionState  IQuestionState
}

var questionIdCounter atomic.Int64

func NewQuestion(body string, title string, author *User, tags ...*Tag) *Question {
	// tags is []*Tag{}  — empty slice, length 0
	return &Question{
		Post:           NewPost(fmt.Sprintf("question-%d", questionIdCounter.Add(1)), body, author, true),
		title:          title,
		tags:           tags,
		answers:        make([]*Answer, 0),
		acceptedAnswer: nil,
		questionState:  &QuestionOpenState{},
	}
}

func (q *Question) AddAnswer(answer *Answer) {
	q.questionState.addAnswer(q, answer)
}

func (q *Question) AcceptAnswer(answer *Answer) {
	q.questionState.acceptAnswer(q, answer)
}

func (q *Question) setState(state IQuestionState) {
	q.questionState = state
}

func (q *Question) GetTitle() string {
	return q.title
}

func (q *Question) GetTags() []*Tag {
	return q.tags
}

func (q *Question) GetAnswers() []*Answer {
	return q.answers
}

func (q *Question) GetAcceptedAnswer() *Answer {
	return q.acceptedAnswer
}
