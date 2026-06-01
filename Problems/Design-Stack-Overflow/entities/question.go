package entities

import (
	"fmt"
	"sync/atomic"

	"stack.overflow/enums"
)

type Question struct {
	Post
	title          string
	tags           []*Tag
	answers        []*Answer
	acceptedAnswer *Answer
	isClosed       bool
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
		isClosed:       false,
	}
}

func (q *Question) AddAnswer(answer *Answer) {
	q.mu.Lock()
	q.answers = append(q.answers, answer)
	q.mu.Unlock()
}

func (q *Question) AcceptAnswer(answer *Answer) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.GetAuthor().GetID() != answer.GetAuthor().GetID() && q.acceptedAnswer == nil {
		q.acceptedAnswer = answer
		q.isClosed = true
		answer.SetAccepted(true)
		q.notifyObservers(NewEvent(enums.ACCEPT_ANSWER, answer.GetAuthor(), &answer.Post))
	}
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
