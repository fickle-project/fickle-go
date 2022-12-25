package users

type iUserService interface {
	NewUser(IFactory, IRepository, CreateUserParam) (User, error)
	FindUser(IRepository, IdUser) (User, error)
	FindUsers(IRepository, QueryUserParam) ([]User, error)
}

func NewService() iUserService {
	return &userService{}
}

// Implements iUserService
type userService struct{}
