package entities

type IContent interface {
	GetID() string
	GetBody() string
	GetAuthor() *User
}
