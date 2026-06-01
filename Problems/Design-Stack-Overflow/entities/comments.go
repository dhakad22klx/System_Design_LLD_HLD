package entities

import (
	"fmt"
	"sync/atomic"
)

type Comment struct {
	Content
}

var commentIdCounter atomic.Int64

func NewComment(body string, author *User) Comment {
	return Comment{Content: NewContent(fmt.Sprintf("Comment-%d", commentIdCounter.Add(1)), body, author)}
}
