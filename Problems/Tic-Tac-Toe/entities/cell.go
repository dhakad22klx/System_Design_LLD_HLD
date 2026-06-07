package entities

import (
	. "tictactoe/enums"
)

type Cell struct {
	symbol Symbol
}

func NewCell() *Cell {
	return &Cell{symbol: SYMBOL_EMPTY}
}

func (cell *Cell) GetSymbol() Symbol {
	return cell.symbol
}

func (cell *Cell) SetSymbol(symbol Symbol) {
	cell.symbol = symbol
}

func (cell *Cell) IsEmpty() bool {
	return cell.symbol == SYMBOL_EMPTY
}
