package users

import (
	"fickle/domain/errors"
	"net/mail"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func (p *UpdateUserParam) validate(id IdUser, r IRepository) error {
	if p.Name == nil && p.Email == nil && p.Password == nil {
		return &errors.ErrNoUpdate{}
	}
	if p.Name != nil && *p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       p.Name,
			Description: "cannot be empty",
		}
	}
	if p.Email != nil {
		if *p.Email == "" {
			return &errors.ErrValidation{
				Property:    "Email",
				Given:       p.Email,
				Description: "cannot be empty",
			}
		}
		if _, err := mail.ParseAddress(*p.Email); err != nil {
			return &errors.ErrValidation{
				Property:    "Email",
				Given:       p.Email,
				Description: "is invalid email address",
			}
		}
		if u, err := r.FindUsers(QueryUserParam{Email: p.Email}); err != nil {
			return err
		} else if len(u) != 0 && u[0].Id != id {
			return &errors.ErrValidation{
				Property:    "Email",
				Given:       p.Email,
				Description: "was already used",
			}
		} else if len(u) != 0 && u[0].Id == id {
			p.Email = nil
		}
	}
	if p.Password != nil && utf8.RuneCountInString(*p.Password) < 8 {
		return &errors.ErrValidation{
			Property:    "Password",
			Given:       p.Password,
			Description: "must be at least 8 characters",
		}
	}
	return nil
}

// Update implements iUser
func (u *User) Update(r IRepository, p UpdateUserParam) (User, error) {
	// Validate parameter
	err := p.validate(u.Id, r)
	if err != nil {
		return User{}, err
	}

	if p.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*p.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, err
		}
		hashedStr := string(hashed)
		p.Password = &hashedStr
	}

	return r.UpdateUser(u.Id, p)
}
