package strategy

import (
	"tictactoe/entities"
	"tictactoe/enums"
)

type RowWinningStrategy struct{}

func (strategy *RowWinningStrategy) CheckWin(board *entities.Board, row int, col int, symbol enums.Symbol) bool {
	size := board.GetSize()
	for c := 0; c < size; c++ {
		if board.GetCell(row, c).GetSymbol() != symbol {
			return false
		}
	}
	return true
}
