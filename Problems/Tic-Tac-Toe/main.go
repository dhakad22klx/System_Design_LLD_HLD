package main

import (
	"fmt"
	"tictactoe/entities"
	"tictactoe/enums"
)

func main() {
	system := GetTicTacToeSystemInstance()

	alice := entities.NewPlayer("Alice", enums.SYMBOL_CROSS)
	bob := entities.NewPlayer("Bob", enums.SYMBOL_ZERO)

	// Game 1: Alice wins
	fmt.Println("========== GAME 1 ==========")
	system.CreateGame(alice, bob)

	system.MakeMove(alice, 0, 0) // X at (0,0)
	system.MakeMove(bob, 1, 0)   // O at (1,0)
	system.MakeMove(alice, 0, 1) // X at (0,1)
	system.MakeMove(bob, 1, 1)   // O at (1,1)
	system.MakeMove(alice, 0, 2) // X at (0,2) - Alice wins!

	gameStatus, _ := system.GetGameStatus()
	fmt.Println("Game 1 Result:", gameStatus)

	// Game 2: Bob wins
	fmt.Println("\n========== GAME 2 ==========")
	system.CreateGame(alice, bob)

	system.MakeMove(alice, 0, 0) // X at (0,0)
	system.MakeMove(bob, 1, 2)   // O at (1,1) - center
	system.MakeMove(alice, 0, 1) // X at (0,1)
	system.MakeMove(bob, 0, 2)   // O at (0,2)
	system.MakeMove(alice, 2, 0) // X at (2,0)
	system.MakeMove(bob, 1, 2)   // O at (2,2) - Bob wins diagonal!
	system.MakeMove(bob, 2, 2)   // O at (2,2) - Bob wins diagonal!
	gameStatus, _ = system.GetGameStatus()
	fmt.Println("Game 2 Result:", gameStatus)

	// Game 3: Draw
	fmt.Println("\n========== GAME 3 ==========")
	system.CreateGame(alice, bob)

	system.MakeMove(alice, 0, 0) // X
	system.MakeMove(bob, 0, 1)   // O
	system.MakeMove(alice, 0, 2) // X
	system.MakeMove(bob, 1, 1)   // O
	system.MakeMove(alice, 1, 0) // X
	system.MakeMove(bob, 1, 2)   // O
	system.MakeMove(alice, 2, 1) // X
	system.MakeMove(bob, 2, 0)   // O
	system.MakeMove(alice, 2, 2) // X - Draw!
	gameStatus, _ = system.GetGameStatus()
	fmt.Print("Game 3 Result:", gameStatus)

	// Final scoreboard
	system.PrintScoreboard()
}
