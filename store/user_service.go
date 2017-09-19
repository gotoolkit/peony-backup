package store

import (
	"github.com/gotoolkit/peony"
)

// UserService represents a service for managing users.
type UserService struct {
	store *Store
}

// User returns a user by username
func (service *UserService) User(username string) (*peony.User, error) {
	return nil, nil
}

// UserByUsername returns a user by username.
func (service *UserService) UserByUsername(username string) (*peony.User, error) {
	user := &peony.User{}
	user.ID = 123
	user.Username = "demo"
	user.Password = "demo"
	return user, nil
}
