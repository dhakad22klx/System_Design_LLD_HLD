package enums

type GameStatus string

const (
	GAME_IN_PROGRESS GameStatus = "IN_PROGRESS"
	WINNER_X         GameStatus = "WINNER_X"
	WINNER_O         GameStatus = "WINNER_O"
	DRAW             GameStatus = "DRAW"
	NO_GAME_EXIST    GameStatus = "NO_GAME_EXIST"
)
