package users

// Remove implements iUser
func (u *User) Remove(r IRepository) error {
	return r.RemoveUser(u.Id)
}
