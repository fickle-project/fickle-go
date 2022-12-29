package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func verifyPasswordHash(hashed []byte, password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashed, password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
