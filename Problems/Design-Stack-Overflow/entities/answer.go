package entities

import (
	"fmt"
	"sync/atomic"
)

type Answer struct {
	Post
	isAccepted bool
}

var answerIdCounter atomic.Int64

func NewAnswer(body string, author *User) *Answer {
	return &Answer{
		Post:       NewPost(fmt.Sprintf("Answer-%d", answerIdCounter.Add(1)), body, author, false),
		isAccepted: false,
	}
}

func (a *Answer) SetAccepted(accepted bool) {
	a.isAccepted = accepted
}

func (a *Answer) IsAcceptedAnswer() bool {
	return a.isAccepted
}
