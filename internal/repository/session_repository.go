package repository

import (
	"context"
	"real-time-forum/internal/domain"
)

// ==============================
// SESSION REPOSITORY
// ==============================
type SessionRepository interface {
	Create(ctx context.Context, session *domain.Session) error
	GetByToken(ctx context.Context, token string) (*domain.Session, error)
	DeleteByToken(ctx context.Context, token string) error
	DeleteUserSessions(ctx context.Context, userID int) error
	DeleteExpiredSessions(ctx context.Context) error
}
