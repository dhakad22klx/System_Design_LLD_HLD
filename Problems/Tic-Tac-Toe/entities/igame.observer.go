package entities

type IGameObserver interface {
	Update(game *Game)
}
