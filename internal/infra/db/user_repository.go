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

// FindByEmailOrNickname retrieves a user from the database by email or nickname
// It takes a context and identifier string that can be either email or nickname
// Returns the user domain object if found, nil if not found, or error if query fails
// Useful for login operations and checking duplicate registrations
func (r *UserRepositoryImpl) FindByEmailOrNickname(ctx context.Context, identifier string) (*domain.User, error) {
	query := `
		SELECT id, email, nickname, first_name, last_name, age, gender, password_hash, created_at
		FROM users
		WHERE email = ? OR nickname = ?
	`
	var user domain.User

	// Execute query and scan results into user struct
	err := r.db.QueryRowContext(ctx, query, identifier, identifier).Scan(
		&user.ID,
		&user.Email,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil without error when user is not found
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByID retrieves a user from the database by their unique identifier
// It takes a context and user ID integer, returns user domain object if found
// Returns nil if user doesn't exist, or error if query fails
// Excludes sensitive password hash from result for security
func (r *UserRepositoryImpl) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `
		SELECT id, email, nickname, first_name, last_name, age, gender, created_at
		FROM users
		WHERE id = ?
	`
	var user domain.User

	// Execute query with user ID and scan results into user struct
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found - return nil without error
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// UpdateProfile modifies an existing user's profile information in the database
// It takes a context and user domain object containing updated profile data
// Returns error if the update operation fails
// Updates all user profile fields except password and creation timestamp
func (r *UserRepositoryImpl) UpdateProfile(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET email = ?, nickname = ?, first_name = ?, last_name = ?, age = ?, gender = ?
		WHERE id = ?
	`
	// Execute update query with user profile data and ID
	_, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Nickname,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		user.ID,
	)
	return err
}

// GetAllForStatus retrieves all users with their online status information
// Returns a slice of UserStatus containing user ID, nickname, online status, 
// last seen timestamp, and update timestamp
// Uses LEFT JOIN to include users without explicit status entries, defaulting 
// online status to false and using creation time as fallback for timestamps
// Results are ordered alphabetically by nickname for consistent presentation
func (r *UserRepositoryImpl) GetAllForStatus() ([]domain.UserStatus, error) {
	ctx := context.Background()

	query := `
        -- Select all users (LEFT JOIN ensures users without a status entry are still included)
        SELECT 
            u.id, 
            u.nickname,
            COALESCE(us.online, FALSE),   -- Default status to FALSE if no entry exists
            COALESCE(us.last_seen, u.created_at), -- Use user creation time as fallback
            COALESCE(us.updated_at, u.created_at)
        FROM users u
        LEFT JOIN user_statuses us ON u.id = us.user_id
        ORDER BY u.nickname ASC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersStatus []domain.UserStatus

	for rows.Next() {
		var uStatus domain.UserStatus

		// Scan the columns into the UserStatus struct fields
		err := rows.Scan(
			&uStatus.ID,
			&uStatus.Nickname,
			&uStatus.Online,
			&uStatus.LastSeen,
			&uStatus.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Set the redundant UserID field to the primary ID for completeness
		uStatus.UserID = uStatus.ID

		usersStatus = append(usersStatus, uStatus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usersStatus, nil
}
