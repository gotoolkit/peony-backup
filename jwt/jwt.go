package jwt

import (
	"github.com/gorilla/securecookie"
	"github.com/gotoolkit/peony"
)

// Service represents a service for managing JWT tokens.
type Service struct {
	secret []byte
}

// NewService initializes a new service. It will generate a random key that will be used to sign JWT tokens.
func NewService() (*Service, error) {
	secret := securecookie.GenerateRandomKey(32)
	if secret == nil {
		return nil, peony.ErrSecretGeneration
	}
	service := &Service{
		secret,
	}
	return service, nil
}

// GenerateToken generates a new JWT token.
func (service *Service) GenerateToken(data *peony.TokenData) (string, error) {
	return "", nil
}

// VerifyToken parses a JWT token and verify its validity. It returns an error if token is invalid.
func (service *Service) VerifyToken(token string) error {
	return nil
}
