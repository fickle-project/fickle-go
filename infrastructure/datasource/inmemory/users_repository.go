package inmemory

import (
	"fickle/domain/errors"
	"fickle/domain/users"
	"strings"
)

func NewRepositoryUsers() users.IRepository {
	return &rUsers{}
}

type rUsers struct {
	data []users.User
}

// AddUser implements users.IRepository
func (r *rUsers) AddUser(u users.User) (users.User, error) {
	r.data = append(r.data, u)
	return u, nil
}

// FindUser implements users.IRepository
func (r *rUsers) FindUser(id users.IdUser) (users.User, error) {
	for _, u := range r.data {
		if u.Id == id {
			return u, nil
		}
	}
	return users.User{}, &errors.ErrNotFound{Object: "User", Id: string(id)}
}

// FindUsers implements users.IRepository
func (r *rUsers) FindUsers(q users.QueryUserParam) ([]users.User, error) {
	u := filter(r.data, func(u users.User) bool {
		selected := true
		if q.Name != nil {
			selected = selected && u.Name == *q.Name
		}
		if q.NameContain != nil {
			selected = selected && strings.Contains(u.Name, *q.NameContain)
		}
		if q.Email != nil {
			selected = selected && u.Email == *q.Email
		}
		return selected
	})
	return u, nil
}

// RemoveUser implements users.IRepository
func (r *rUsers) RemoveUser(id users.IdUser) error {
	l := len(r.data)
	r.data = filter(r.data, func(u users.User) bool { return u.Id != id })
	if l == len(r.data) {
		return &errors.ErrNotFound{Object: "User", Id: string(id)}
	}
	return nil
}

// UpdateUser implements users.IRepository
func (r *rUsers) UpdateUser(id users.IdUser, p users.UpdateUserParam) (users.User, error) {
	u, err := r.FindUser(id)
	if err != nil {
		return users.User{}, err
	}

	err = r.RemoveUser(id)
	if err != nil {
		return users.User{}, err
	}

	if p.Name != nil {
		u.Name = *p.Name
	}
	if p.Email != nil {
		u.Email = *p.Email
	}
	if p.Password != nil {
		u.PasswordHash = *p.Password
	}
	return r.AddUser(u)
}
