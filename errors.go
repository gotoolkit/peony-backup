package peony

// JWT errors.
const (
	ErrSecretGeneration = Error("Unable to generate secret key")
	ErrInvalidJWTToken  = Error("Invalid JWT token")
)

// Error represents an application error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}
