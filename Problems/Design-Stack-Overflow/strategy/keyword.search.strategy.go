package strategy

import (
	"strings"

	"stack.overflow/entities"
)

type KeywordSearchStrategy struct {
	keyword string
}

func NewKeywordSearchStrategy(keyword string) *KeywordSearchStrategy {
	return &KeywordSearchStrategy{
		keyword: keyword,
	}
}

func (s *KeywordSearchStrategy) Filter(q []*entities.Question) []*entities.Question {
	result := make([]*entities.Question, 0)

	for _, question := range q {
		questionTitle := strings.ToLower(question.GetTitle())
		questionBody := strings.ToLower(question.GetBody())

		if strings.Contains(questionTitle, s.keyword) || strings.Contains(questionBody, s.keyword) {
			result = append(result, question)
		}
	}

	return result
}
