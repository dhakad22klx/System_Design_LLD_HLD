package iterator

// Concrete Iterator
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {

	return u.index < len(u.users)
	// if u.index < len(u.users) {
	// 	return true
	// }
	// return false

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
