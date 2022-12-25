package users

// FindUser implements iUserService
func (*userService) FindUser(r IRepository, id IdUser) (User, error) {
	return r.FindUser(id)
}

// FindUsers implements iUserService
func (*userService) FindUsers(r IRepository, p QueryUserParam) ([]User, error) {
	return r.FindUsers(p)
}
