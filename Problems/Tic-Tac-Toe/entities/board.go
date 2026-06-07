package entities

import (
	"fmt"
	"strings"
	"tictactoe/enums"
	"tictactoe/exception"
)

type Board struct {
	grid [][]*Cell
	size int
}

func NewBoard(size int) *Board {
	board := &Board{grid: make([][]*Cell, size), size: size}
	board.initializeBoard()
	return board
}

func (board *Board) initializeBoard() {
	for i := 0; i < board.size; i++ {
		board.grid[i] = make([]*Cell, board.size)
		for j := 0; j < board.size; j++ {
			board.grid[i][j] = NewCell()
		}
	}
}

func (board *Board) GetSize() int {
	return board.size
}
func (board *Board) GetCell(row int, col int) *Cell {
	return board.grid[row][col]
}

func (board *Board) IsCellEmpty(row int, col int) bool {
	board.validatePosition(row, col)
	return board.grid[row][col].IsEmpty()
}

func (board *Board) IsFull() bool {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			if board.grid[i][j].IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (board *Board) PlaceSymbol(row int, col int, symbol enums.Symbol) {
	err := board.validatePosition(row, col)
	if err != nil {
		panic(fmt.Sprintf("Invalid Placement of Symbol %v", err))
	}
	board.grid[row][col].SetSymbol(symbol)
}

func (board *Board) validatePosition(row int, col int) error {
	if row < 0 || row >= board.size || col < 0 || col >= board.size {
		return exception.NewInvalidMoveError(fmt.Sprintf("Position (%d, %d) is out of bounds", row, col))
	}
	return nil
}

func (board *Board) PrintBoard() {
	fmt.Println()
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			fmt.Printf(" %c ", board.grid[i][j].GetSymbol())
			if j < board.size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < board.size-1 {
			fmt.Println(strings.Repeat("-", board.size*4-1))
		}
	}
	fmt.Println()
}
