package entities

import "time"

type Content struct {
	id         string
	body       string
	author     *User
	creationAt time.Time
}

func NewContent(id string, body string, auther *User) Content {
	return Content{
		id:         id,
		body:       body,
		author:     auther,
		creationAt: time.Now(),
	}
}

func (c *Content) GetID() string { return c.id }

func (c *Content) GetBody() string { return c.body }

func (c *Content) GetAuthor() *User { return c.author }
