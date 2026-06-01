package strategy

import (
	"strings"

	"stack.overflow/entities"
)

type TagSearchStrategy struct {
	tag *entities.Tag
}

func NewTagSearchStrategy(tag *entities.Tag) *TagSearchStrategy {
	return &TagSearchStrategy{
		tag: tag,
	}
}

func (s *TagSearchStrategy) Filter(q []*entities.Question) []*entities.Question {

	result := make([]*entities.Question, 0)

	for _, question := range q {
		for _, tag := range question.GetTags() {
			if strings.EqualFold(tag.GetName(), s.tag.GetName()) { //case insensitive comparision
				result = append(result, question)
				break
			}
		}
	}

	return result
}
