package peony

type (
	// User represent a user account.
	User struct {
		ID       UserID `json:"Id"`
		Username string `json:"Username"`
		Password string `json:"Password,omitempty"`
	}
	// UserID represents a user identifier
	UserID int

	// UserRole represents the role of a user. It can be either an administrator
	// or a regular user
	UserRole int

	// TokenData represents the data embedded in a JWT token.
	TokenData struct {
		ID       UserID
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
		UserByUsername(username string) (*User, error)
	}

	// JWTService represents a service for managing JWT tokens
	JWTService interface {
		GenerateToken(data *TokenData) (string, error)
		ParseAndVerifyToken(token string) (*TokenData, error)
	}
	// // PermissionService represents a service for managing user permission
	// PermissionService interface {
	// 	Permission()
	// }
)
