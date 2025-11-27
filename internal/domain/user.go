package domain

import "time"

type User struct {
	ID           int
	Email        string
	Nickname     string
	PasswordHash string
	FirstName    string
	LastName     string
	Age          int
	Gender       string
	CreatedAt    time.Time
}

type UserStatus struct {
	ID        int
	UserID    int
	Nickname  string
	Online    bool
	LastSeen  time.Time
	UpdatedAt time.Time
}
