package db

import (
	"database/sql"
	"real-time-forum/internal/repository"
)

// UserRepositoryImpl implements the UserRepository interface
// providing concrete database operations for user management
type UserRepositoryImpl struct {
	db *sql.DB // database connection instance
}

// NewUserRepository creates and returns a new instance of UserRepositoryImpl
// initialized with the provided database connection
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}
