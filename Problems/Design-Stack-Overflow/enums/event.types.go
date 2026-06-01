package enums

type EventType int

const (
	UPVOTE_QUESTION EventType = iota
	DOWNVOTE_QUESTION
	UPVOTE_ANSWER
	DOWNVOTE_ANSWER
	ACCEPT_ANSWER
)
