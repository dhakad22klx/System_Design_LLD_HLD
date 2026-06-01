package main

import (
	"sync"

	"stack.overflow/entities"
	"stack.overflow/enums"
	"stack.overflow/observer"
	"stack.overflow/strategy"
)

type StackOverflowService struct {
	users                  map[string]*entities.User
	questions              map[string]*entities.Question
	answers                map[string]*entities.Answer
	reputationManager      entities.IPostObserver
	questionSearchStrategy strategy.ISearchStrategy
	mu                     sync.RWMutex
}

func NewStackOverflowService() *StackOverflowService {
	return &StackOverflowService{
		users:             make(map[string]*entities.User),
		questions:         make(map[string]*entities.Question),
		answers:           make(map[string]*entities.Answer),
		reputationManager: observer.NewReputationManager(),
	}
}

func (s *StackOverflowService) CreateUser(name string) *entities.User {
	user := entities.NewUser(name)
	s.mu.Lock()
	s.users[user.GetID()] = user
	s.mu.Unlock()
	return user
}

func (s *StackOverflowService) PostQuestion(userID string, title string, body string, tags []*entities.Tag) *entities.Question {
	s.mu.RLock()
	author := s.users[userID]
	s.mu.RUnlock()

	question := entities.NewQuestion(title, body, author, tags...)
	question.AddObserver(s.reputationManager)

	s.mu.Lock()
	s.questions[question.GetID()] = question
	s.mu.Unlock()
	return question
}

func (s *StackOverflowService) PostAnswer(userID string, questionID string, body string) *entities.Answer {
	s.mu.RLock()
	author := s.users[userID]
	question := s.questions[questionID]
	s.mu.RUnlock()

	answer := entities.NewAnswer(body, author)
	answer.AddObserver(s.reputationManager)
	question.AddAnswer(answer)

	s.mu.Lock()
	s.answers[answer.GetID()] = answer
	s.mu.Unlock()
	return answer
}

func (s *StackOverflowService) VoteOnPost(userID string, postID string, voteType enums.VoteType) {
	s.mu.RLock()
	user := s.users[userID]
	s.mu.RUnlock()

	post := s.findPostByID(postID)
	post.Vote(user, voteType)
}

func (s *StackOverflowService) AcceptAnswer(questionID string, answerID string) {
	s.mu.RLock()
	question := s.questions[questionID]
	answer := s.answers[answerID]
	s.mu.RUnlock()
	question.AcceptAnswer(answer)
}

func (s *StackOverflowService) SetQuestionSearchStrategy(strategy strategy.ISearchStrategy) {
	s.questionSearchStrategy = strategy
}
func (s *StackOverflowService) SearchQuestions() []*entities.Question {
	s.mu.RLock()
	results := make([]*entities.Question, 0, len(s.questions))
	for _, question := range s.questions {
		results = append(results, question)
	}
	s.mu.RUnlock()

	// Search by applying set strategy
	results = s.questionSearchStrategy.Filter(results)
	return results
}

func (s *StackOverflowService) GetUser(userID string) *entities.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users[userID]
}

func (s *StackOverflowService) findPostByID(postID string) *entities.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if question, ok := s.questions[postID]; ok {
		return &question.Post
	}
	if answer, ok := s.answers[postID]; ok {
		return &answer.Post
	}

	panic("Post not found")
}
