package users

import (
	"fickle/domain/errors"
	"net/mail"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func (p *CreateUserParam) validate(r IRepository) error {
	if p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       "",
			Description: "cannot be empty",
		}
	}
	if p.Email == "" {
		return &errors.ErrValidation{
			Property:    "Email",
			Given:       "",
			Description: "cannot be empty",
		}
	}
	if _, err := mail.ParseAddress(p.Email); err != nil {
		return &errors.ErrValidation{
			Property:    "Email",
			Given:       p.Email,
			Description: "is invalid email address",
		}
	}
	if u, err := r.FindUsers(QueryUserParam{Email: &p.Email}); err != nil {
		return err
	} else if len(u) != 0 {
		return &errors.ErrValidation{
			Property:    "Email",
			Given:       p.Email,
			Description: "was already used",
		}
	}
	if utf8.RuneCountInString(p.Password) < 8 {
		return &errors.ErrValidation{
			Property:    "Email",
			Given:       p.Password,
			Description: "must be at least 8 characters",
		}
	}
	return nil
}

// NewUser implements iUserService
func (u *userService) NewUser(f IFactory, r IRepository, p CreateUserParam) (User, error) {
	// Validate parameter
	err := p.validate(r)
	if err != nil {
		return User{}, err
	}

	id, err := f.NewUserId(r)
	if err != nil {
		return User{}, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return r.AddUser(User{
		Id:           id,
		Name:         p.Name,
		Email:        p.Email,
		PasswordHash: string(hashed),
	})
}
