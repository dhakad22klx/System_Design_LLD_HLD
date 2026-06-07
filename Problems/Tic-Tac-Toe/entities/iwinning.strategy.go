package entities

import (
	"tictactoe/enums"
)

type IWinningStrategy interface {
	CheckWin(board *Board, row int, col int, symbol enums.Symbol) bool
}
