package strategy

import (
	"tictactoe/entities"
	"tictactoe/enums"
)

type DiagonalWinningStrategy struct{}

func (strategy *DiagonalWinningStrategy) CheckWin(board *entities.Board, row int, col int, symbol enums.Symbol) bool {
	size := board.GetSize()

	// Check main diagonal (top-left to bottom-right)
	mainDiagonalWin := true
	for i := 0; i < size; i++ {
		if board.GetCell(i, i).GetSymbol() != symbol {
			mainDiagonalWin = false
			break
		}
	}
	if mainDiagonalWin {
		return true
	}

	// Check anti-diagonal (top-right to bottom-left)
	for i := 0; i < size; i++ {
		if board.GetCell(i, size-1-i).GetSymbol() != symbol {
			return false
		}
	}
	return true
}
