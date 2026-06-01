package entities

type IPostObserver interface {
	OnPostEvent(e Event)
}
