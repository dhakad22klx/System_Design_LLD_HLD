package entities

import (
	"sync"

	"stack.overflow/enums"
)

type Post struct {
	Content
	voteCount  int
	voters     map[string]enums.VoteType
	comments   []*Comment
	observers  []IPostObserver
	isQuestion bool
	mu         sync.Mutex
}

func NewPost(id string, body string, author *User, isQuestion bool) Post {
	return Post{
		Content:    NewContent(id, body, author),
		voters:     make(map[string]enums.VoteType),
		comments:   make([]*Comment, 0),
		observers:  make([]IPostObserver, 0),
		isQuestion: isQuestion,
	}
}

func (p *Post) AddObserver(observer IPostObserver) {
	p.observers = append(p.observers, observer)
}

func (p *Post) notifyObservers(event Event) {
	for _, observer := range p.observers {
		observer.OnPostEvent(event)
	}
}

func (p *Post) Vote(user *User, voteType enums.VoteType) {
	p.mu.Lock()
	defer p.mu.Unlock()

	userID := user.GetID()
	existingVote, voted := p.voters[userID]
	if voted && existingVote == voteType {
		return // Already voted
	}

	scoreChange := 0
	if voted { // User is changing their vote
		if voteType == enums.UPVOTE {
			scoreChange = 2
		} else {
			scoreChange = -2
		}
	} else { // New vote
		if voteType == enums.UPVOTE {
			scoreChange = 1
		} else {
			scoreChange = -1
		}
	}

	p.voters[userID] = voteType
	p.voteCount += scoreChange

	eventType := enums.UPVOTE_QUESTION
	if p.isQuestion {
		if voteType == enums.UPVOTE {
			eventType = enums.UPVOTE_QUESTION
		} else {
			eventType = enums.DOWNVOTE_QUESTION
		}
	} else {
		if voteType == enums.UPVOTE {
			eventType = enums.UPVOTE_ANSWER
		} else {
			eventType = enums.DOWNVOTE_ANSWER
		}
	}

	p.notifyObservers(NewEvent(eventType, user, p))
}

func (p *Post) GetVoteCount() int {
	return p.voteCount
}
