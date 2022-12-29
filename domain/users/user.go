package users

type iUser interface {
	Update(IRepository, UpdateUserParam) (User, error)
	Remove(IRepository) error
	Verify(password []byte) (bool, error)
}

type (
	CreateUserParam struct {
		Name     string
		Email    string
		Password string
	}

	UpdateUserParam struct {
		Name     *string
		Email    *string
		Password *string
	}

	QueryUserParam struct {
		Name        *string
		NameContain *string
		Email       *string
	}

	IdUser string
)

// Implements
var _ iUser = &User{}

// Implements iUser
type User struct {
	Id           IdUser
	Name         string
	Email        string
	PasswordHash string
}
