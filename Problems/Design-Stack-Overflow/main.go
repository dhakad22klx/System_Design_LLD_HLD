package main

import (
	"fmt"

	"stack.overflow/entities"
	"stack.overflow/enums"
	"stack.overflow/strategy"
)

func main() {
	service := NewStackOverflowService()

	// 1. Create Users
	alice := service.CreateUser("Alice")
	bob := service.CreateUser("Bob")
	charlie := service.CreateUser("Charlie")
	deepak := service.CreateUser("Deepak Dhakad")

	fmt.Printf("Create user IDs %s, %s, %s.\n", alice.GetID(), bob.GetID(), charlie.GetID())

	// 2. Alice posts a question
	fmt.Println("--- Alice posts a question ---")
	javaTag := entities.NewTag("java")
	designPatternsTag := entities.NewTag("design-patterns")
	tags := []*entities.Tag{javaTag, designPatternsTag}
	question := service.PostQuestion(alice.GetID(), "How to implement Observer Pattern?", "Details about Observer Pattern...", tags)
	printReputations(alice, bob, charlie)

	// 3. Bob and Charlie post answers
	fmt.Println("\n--- Bob and Charlie post answers ---")
	bobAnswer := service.PostAnswer(bob.GetID(), question.GetID(), "You can use the java.util.Observer interface.")
	charlieAnswer := service.PostAnswer(charlie.GetID(), question.GetID(), "A better way is to create your own Observer interface.")
	printReputations(alice, bob, charlie)

	fmt.Printf("IDs of Posts %s, %s, %s.\n", question.GetID(), bobAnswer.GetID(), charlieAnswer.GetID())
	// 4. Voting happens
	fmt.Println("\n--- Voting Occurs ---")
	service.VoteOnPost(alice.GetID(), question.GetID(), enums.UPVOTE)    // Alice upvotes her own question
	service.VoteOnPost(bob.GetID(), charlieAnswer.GetID(), enums.UPVOTE) // Bob upvotes Charlie's answer
	service.VoteOnPost(alice.GetID(), bobAnswer.GetID(), enums.DOWNVOTE) // Alice downvotes Bob's answer
	printReputations(alice, bob, charlie)

	// 5. Alice accepts Charlie's answer
	fmt.Println("\n--- Alice accepts Charlie's answer ---")
	service.AcceptAnswer(question.GetID(), charlieAnswer.GetID())
	printReputations(alice, bob, charlie)

	//Testing state pattern implementation
	service.PostAnswer(deepak.GetID(), question.GetID(), "You can use the java.util.Observer interface.")
	service.AcceptAnswer(question.GetID(), charlieAnswer.GetID())

	// 6.1 Search for Questions by Keyword
	fmt.Println("\n Question Search by keyword: 'implement'---")
	keywordSearchStrategy := strategy.NewKeywordSearchStrategy("implement")
	service.SetQuestionSearchStrategy(keywordSearchStrategy)
	searchResults := service.SearchQuestions()
	for _, question := range searchResults {
		fmt.Println("  - Found: " + question.GetTitle())
	}
	// 6.2 Search for Questions by User
	fmt.Println("\n Question Search by user: 'alice'---")
	userSearchStrategy := strategy.NewUserSearchStrategy(alice)
	service.SetQuestionSearchStrategy(userSearchStrategy)
	searchResults = service.SearchQuestions()
	for _, question := range searchResults {
		fmt.Println("  - Found: " + question.GetTitle())
	}
	// 6.3 Search for Questions by Tags
	fmt.Println("\n Question Search by tags: user 'java'---")
	tagSearchStrategy := strategy.NewTagSearchStrategy(javaTag)
	service.SetQuestionSearchStrategy(tagSearchStrategy)
	searchResults = service.SearchQuestions()
	for _, question := range searchResults {
		fmt.Println("  - Found: " + question.GetTitle())
	}
}

func printReputations(users ...*entities.User) {
	fmt.Println("--- Current Reputations ---")
	for _, user := range users {
		fmt.Printf("%s: %d\n", user.GetName(), user.GetReputation())
	}
}
