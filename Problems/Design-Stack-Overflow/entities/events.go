package entities

import (
	"stack.overflow/enums"
)

type Event struct {
	eventType  enums.EventType
	actor      *User
	targetPost *Post
}

func NewEvent(eventType enums.EventType, actor *User, targetPost *Post) Event {
	return Event{
		eventType:  eventType,
		actor:      actor,
		targetPost: targetPost,
	}
}

func (e Event) GetType() enums.EventType {
	return e.eventType
}

func (e Event) GetActor() *User {
	return e.actor
}

func (e Event) GetTargetPost() *Post {
	return e.targetPost
}
