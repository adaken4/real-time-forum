package domain

import "time"

type Session struct {
	ID           int
	UserID       int
	SessionToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
