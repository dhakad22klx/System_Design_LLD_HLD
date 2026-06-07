package main

import (
	"fmt"
	"sync"
	"tictactoe/entities"
	"tictactoe/enums"
	"tictactoe/observer"
	"tictactoe/strategy"
)

type TicTacToeSystem struct {
	scoreboard  *observer.Scoreboard
	currentGame *entities.Game
	mu          sync.Mutex
}

var ticTacToeSystemInstance *TicTacToeSystem
var ticTacToeSystemLock sync.Mutex

func GetTicTacToeSystemInstance() *TicTacToeSystem {
	ticTacToeSystemLock.Lock()
	defer ticTacToeSystemLock.Unlock()

	if ticTacToeSystemInstance == nil {
		ticTacToeSystemInstance = &TicTacToeSystem{scoreboard: observer.NewScoreboard()}
	}
	return ticTacToeSystemInstance
}

func (system *TicTacToeSystem) CreateGame(player1 *entities.Player, player2 *entities.Player) *entities.Game {
	system.mu.Lock()
	defer system.mu.Unlock()

	system.currentGame = entities.NewGame(player1, player2, 3)
	system.currentGame.AddObserver(system.scoreboard)
	system.currentGame.SetWinnigStrategies([]entities.IWinningStrategy{&strategy.ColumnWinningStrategy{},
		&strategy.RowWinningStrategy{}, &strategy.DiagonalWinningStrategy{}})
	fmt.Println("New game started:", player1.GetName(), "vs", player2.GetName())
	return system.currentGame
}

func (system *TicTacToeSystem) MakeMove(player *entities.Player, row int, col int) error {
	system.mu.Lock()
	defer system.mu.Unlock()

	if system.currentGame == nil {
		return fmt.Errorf("no active game. call createGame first")
	}
	fmt.Printf("%s plays at (%d, %d)\n", player.GetName(), row, col)
	system.currentGame.MakeMove(row, col)
	system.currentGame.PrintBoard()

	return nil
}

func (system *TicTacToeSystem) GetGameStatus() (enums.GameStatus, error) {
	system.mu.Lock()
	defer system.mu.Unlock()

	if system.currentGame == nil {
		return enums.NO_GAME_EXIST, fmt.Errorf("no active game.")
	}
	return system.currentGame.GetStatus(), nil
}

func (system *TicTacToeSystem) PrintScoreboard() {
	system.scoreboard.PrintScoreboard()
}

// For testing: reset the singleton
func ResetTicTacToeSystemInstance() {
	ticTacToeSystemLock.Lock()
	defer ticTacToeSystemLock.Unlock()
	ticTacToeSystemInstance = nil
}
