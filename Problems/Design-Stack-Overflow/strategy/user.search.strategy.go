package strategy

import "stack.overflow/entities"

type UserSearchStrategy struct {
	user *entities.User
}

func NewUserSearchStrategy(user *entities.User) *UserSearchStrategy {
	return &UserSearchStrategy{
		user: user,
	}
}

func (s *UserSearchStrategy) Filter(q []*entities.Question) []*entities.Question {
	result := make([]*entities.Question, 0)

	for _, question := range q {
		if question.GetAuthor().GetID() == s.user.GetID() {
			result = append(result, question)
		}
	}

	return result
}
