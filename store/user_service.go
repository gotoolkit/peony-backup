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
