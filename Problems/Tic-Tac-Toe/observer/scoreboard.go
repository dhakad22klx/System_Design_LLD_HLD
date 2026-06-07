package observer

import (
	"fmt"
	"sync"
	"tictactoe/entities"
)

type Scoreboard struct {
	scores map[string]int
	mu     sync.Mutex
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{scores: make(map[string]int)}
}

func (scoreboard *Scoreboard) Update(game *entities.Game) {
	winner := game.GetWinner()
	if winner != nil {
		scoreboard.RecordWin(winner)
		fmt.Println("Scoreboard updated:", winner.GetName(), "wins!")
	}
}

func (scoreboard *Scoreboard) RecordWin(player *entities.Player) {
	scoreboard.mu.Lock()
	defer scoreboard.mu.Unlock()
	scoreboard.scores[player.GetName()]++
}

func (scoreboard *Scoreboard) GetScore(playerName string) int {
	scoreboard.mu.Lock()
	defer scoreboard.mu.Unlock()
	return scoreboard.scores[playerName]
}

func (scoreboard *Scoreboard) PrintScoreboard() {
	scoreboard.mu.Lock()
	defer scoreboard.mu.Unlock()

	fmt.Println("\n===== SCOREBOARD =====")
	if len(scoreboard.scores) == 0 {
		fmt.Println("No games played yet.")
	} else {
		for name, score := range scoreboard.scores {
			fmt.Printf("%s: %d wins\n", name, score)
		}
	}
	fmt.Printf("======================\n\n")
}
