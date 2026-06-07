package strategy

import (
	"tictactoe/entities"
	"tictactoe/enums"
)

type ColumnWinningStrategy struct{}

func (strategy *ColumnWinningStrategy) CheckWin(board *entities.Board, row int, col int, symbol enums.Symbol) bool {
	size := board.GetSize()
	for r := 0; r < size; r++ {
		if board.GetCell(r, col).GetSymbol() != symbol {
			return false
		}
	}
	return true
}
