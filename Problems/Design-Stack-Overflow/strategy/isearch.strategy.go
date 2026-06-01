package strategy

import "stack.overflow/entities"

type ISearchStrategy interface {
	Filter(q []*entities.Question) []*entities.Question
}
