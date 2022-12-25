package users

type IFactory interface {
	NewUserId(r IRepository) (IdUser, error)
}
