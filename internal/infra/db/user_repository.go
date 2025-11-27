package db

import (
	"context"
	"database/sql"
	"real-time-forum/internal/domain"
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

// Create inserts a new user record into the database
// It takes a context, user domain object, and hashed password string
// Returns error if the operation fails or if retrieving the last insert ID fails
// On success, it updates the user object with the generated ID
func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User, hashedPassword string) error {
	query := `
		INSERT INTO users (email, nickname, first_name, last_name, age, gender, password_hash, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Execute the insert query with user data
	result, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Nickname,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		hashedPassword,
		user.CreatedAt,
	)
	if err != nil {
		return err
	}

	// Retrieve the auto-generated ID from the insert operation
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Update the user domain object with the new ID
	user.ID = int(id)
	return nil
}
