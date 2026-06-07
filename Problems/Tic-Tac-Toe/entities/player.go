package entities

import (
	"fmt"
	"tictactoe/enums"
)

type Player struct {
	name   string
	symbol enums.Symbol
}

func NewPlayer(name string, symbol enums.Symbol) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetSymbol() enums.Symbol {
	return player.symbol
}

func (player *Player) String() string {
	return fmt.Sprintf("%s (%c)", player.name, player.symbol)
}
