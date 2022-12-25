package inmemory

import (
	"fickle/domain/users"

	"github.com/google/uuid"
)

func NewFactoryUsers() users.IFactory {
	return &fUsers{}
}

type fUsers struct{}

// NewUserId implements users.IFactory
func (*fUsers) NewUserId(r users.IRepository) (users.IdUser, error) {
	return users.IdUser(uuid.NewString()), nil
}
