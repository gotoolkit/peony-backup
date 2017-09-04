package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Option functional options
type Option func(*Store)

// Dialect option orm dialect
func Dialect(dialect string) Option {
	return func(s *Store) {
		s.dialect = dialect
	}
}

// Args option orm args
func Args(arg string) Option {
	return func(s *Store) {
		s.args = arg
	}
}

// Store defines the implementation of peony.DataStore using
type Store struct {
	dialect string
	args    string

	// Services
	UserService *UserService

	db *gorm.DB
}

// NewStore initializes a new Store and the associated services
func NewStore(opts ...Option) *Store {
	s := &Store{
		UserService: &UserService{},
	}

	s.UserService.store = s

	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Open opens and initializes the Mysql database.
func (store *Store) Open() error {
	db, err := gorm.Open(store.dialect, store.args)
	if err != nil {
		return err
	}
	store.db = db
	return nil
}

// Close closes the Mysql database.
func (store *Store) Close() error {
	if store.db != nil {
		return store.db.Close()
	}
	return nil
}
