package entities

import (
	"fmt"
	"sync"
	"tictactoe/enums"
	"tictactoe/exception"
)

type Game struct {
	board              *Board
	players            [2]*Player
	currentPlayerIndex int
	status             enums.GameStatus
	winningStrategies  []IWinningStrategy
	observers          []IGameObserver
	mu                 sync.Mutex
}

func NewGame(player1 *Player, player2 *Player, boardSize int) *Game {
	game := &Game{
		board:              NewBoard(boardSize),
		players:            [2]*Player{player1, player2},
		currentPlayerIndex: 0,
		status:             enums.GAME_IN_PROGRESS,
		observers:          make([]IGameObserver, 0),
		winningStrategies:  make([]IWinningStrategy, 0),
	}
	return game
}

func (game *Game) SetWinnigStrategies(strategies []IWinningStrategy) {
	game.winningStrategies = strategies
}

func (game *Game) MakeMove(row int, col int) {
	game.mu.Lock()
	defer game.mu.Unlock()

	// Check if game is already over
	if game.status != enums.GAME_IN_PROGRESS {
		panic(exception.NewInvalidMoveError("Game is already over!"))
	}

	// Validate the move
	if !game.board.IsCellEmpty(row, col) {
		panic(exception.NewInvalidMoveError(fmt.Sprintf("Cell (%d, %d) is already occupied", row, col)))
	}

	// Place the symbol
	currentPlayer := game.players[game.currentPlayerIndex]
	game.board.PlaceSymbol(row, col, currentPlayer.GetSymbol())

	// Check for win
	if game.checkWin(row, col, currentPlayer.GetSymbol()) {
		if currentPlayer.GetSymbol() == enums.SYMBOL_CROSS {
			game.status = enums.WINNER_X
		} else {
			game.status = enums.WINNER_O
		}
		game.NotifyObservers()
		return
	}

	// Check for draw
	if game.board.IsFull() {
		game.status = enums.DRAW
		game.NotifyObservers()
		return
	}

	// Switch to next player
	game.currentPlayerIndex = (game.currentPlayerIndex + 1) % 2
}

func (game *Game) checkWin(row int, col int, symbol enums.Symbol) bool {
	for _, strategy := range game.winningStrategies {
		if strategy.CheckWin(game.board, row, col, symbol) {
			return true
		}
	}
	return false
}

func (game *Game) AddObserver(observer IGameObserver) {
	game.mu.Lock()
	defer game.mu.Unlock()
	game.observers = append(game.observers, observer)
}

func (game *Game) NotifyObservers() {
	for _, observer := range game.observers {
		observer.Update(game)
	}
}

func (game *Game) GetBoard() *Board {
	return game.board
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.players[game.currentPlayerIndex]
}

func (game *Game) GetStatus() enums.GameStatus {
	return game.status
}

func (game *Game) GetWinner() *Player {
	if game.status == enums.WINNER_X {
		if game.players[0].GetSymbol() == enums.SYMBOL_CROSS {
			return game.players[0]
		}
		return game.players[1]
	} else if game.status == enums.WINNER_O {
		if game.players[0].GetSymbol() == enums.SYMBOL_ZERO {
			return game.players[0]
		}
		return game.players[1]
	}
	return nil
}

func (game *Game) PrintBoard() {
	game.board.PrintBoard()
}
