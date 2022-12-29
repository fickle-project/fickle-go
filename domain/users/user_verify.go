package users

// Verify implements iUser
func (u *User) Verify(password []byte) (bool, error) {
	return verifyPasswordHash([]byte(u.PasswordHash), password)
}
