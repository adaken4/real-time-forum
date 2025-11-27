package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver import with side effects
)

// SQLiteDB wraps sql.DB to provide SQLite-specific database operations
// and connection management with optimized pooling settings
type SQLiteDB struct {
	*sql.DB // Embed the standard sql.DB to inherit all its methods
}

// NewSQLiteDB creates and configures a new SQLite database connection
// It takes the database file path and returns a configured SQLiteDB instance
// Sets connection pool limits and lifetime for optimal performance
func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	// Open database connection using SQLite driver
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Configure connection pool settings for better performance
	db.SetMaxOpenConns(25)                 // Maximum number of open connections
	db.SetMaxIdleConns(25)                 // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime

	return &SQLiteDB{db}, nil
}
