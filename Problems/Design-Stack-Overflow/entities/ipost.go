package entities

import (
	"stack.overflow/enums"
)

type IPost interface {
	IContent // promoted — all Content methods included
	// Observer management
	AddObserver(observer IPostObserver)

	// Voting
	Vote(user *User, voteType enums.VoteType)
	GetVoteCount() int
}
