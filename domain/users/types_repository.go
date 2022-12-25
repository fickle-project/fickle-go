package users

type IRepository interface {
	AddUser(User) (User, error)
	UpdateUser(IdUser, UpdateUserParam) (User, error)
	RemoveUser(IdUser) error
	FindUsers(QueryUserParam) ([]User, error)
	FindUser(IdUser) (User, error)
}
