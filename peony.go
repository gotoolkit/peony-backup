package peony

type (
	// User represent a user account.
	User struct {
		Username string `json:"Username"`
		Password string `json:"Password,omitempty"`
	}

	// TokenData represents the data embedded in a JWT token.
	TokenData struct {
		Username string
	}

	// DataStore defines the interface to manage the data.
	DataStore interface {
		Open() error
		Close() error
	}

	// Server defines the interface to serve the data.
	Server interface {
		Start() error
	}

	// UserService represents a service for managing users.
	UserService interface {
		User(username string) (*User, error)
	}

	// JWTService represents a service for managing JWT tokens
	JWTService interface {
		GenerateToken(data *TokenData) (string, error)
		VerifyToken(token string) error
	}
	// // PermissionService represents a service for managing user permission
	// PermissionService interface {
	// 	Permission()
	// }
)
