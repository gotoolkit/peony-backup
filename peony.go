package peony

type (
	// User represent a user account.
	User struct {
		Username string `json:"Username"`
		Password string `json:"Password,omitempty"`
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

	UserService interface {
		User(username string) (*User, error)
	}
)
