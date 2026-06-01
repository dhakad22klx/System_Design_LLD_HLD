package observer

import (
	"stack.overflow/entities"
	"stack.overflow/enums"
)

type ReputationManager struct{}

const (
	QUESTION_UPVOTE_REP        = 5
	ANSWER_UPVOTE_REP          = 10
	ACCEPTED_ANSWER_REP        = 15
	DOWNVOTE_REP_PENALTY       = -1 // Penalty for the voter
	POST_DOWNVOTED_REP_PENALTY = -2 // Penalty for the post author
)

func NewReputationManager() *ReputationManager {
	return &ReputationManager{}
}

func (r *ReputationManager) OnPostEvent(event entities.Event) {
	postAuthor := event.GetTargetPost().GetAuthor()

	switch event.GetType() {
	case enums.UPVOTE_QUESTION:
		postAuthor.UpdateReputation(QUESTION_UPVOTE_REP)
	case enums.DOWNVOTE_QUESTION:
		postAuthor.UpdateReputation(DOWNVOTE_REP_PENALTY)
		event.GetActor().UpdateReputation(POST_DOWNVOTED_REP_PENALTY) // voter penalty
	case enums.UPVOTE_ANSWER:
		postAuthor.UpdateReputation(ANSWER_UPVOTE_REP)
	case enums.DOWNVOTE_ANSWER:
		postAuthor.UpdateReputation(DOWNVOTE_REP_PENALTY)
		event.GetActor().UpdateReputation(POST_DOWNVOTED_REP_PENALTY)
	case enums.ACCEPT_ANSWER:
		postAuthor.UpdateReputation(ACCEPTED_ANSWER_REP)
	}
}
