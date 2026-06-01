package entities

import (
	"fmt"
	"sync/atomic"
)

var userIdCounter atomic.Int64

type User struct {
	id         string
	name       string
	reputation atomic.Int64
}

func NewUser(name string) *User {
	return &User{
		id:   fmt.Sprintf("user-%d", userIdCounter.Add(1)),
		name: name,
	}
}

func (u *User) UpdateReputation(change int) {
	u.reputation.Add(int64(change))
}

func (u *User) GetID() string      { return u.id }
func (u *User) GetName() string    { return u.name }
func (u *User) GetReputation() int { return int(u.reputation.Load()) }
