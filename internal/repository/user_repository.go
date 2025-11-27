package repository

import (
	"context"
	"real-time-forum/internal/domain"
)

// ==============================
// USER REPOSITORY
// ==============================
type UserRepository interface {
	Create(ctx context.Context, user *domain.User, hashedPassword string) error
	FindByEmailOrNickname(ctx context.Context, identifier string) (*domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	UpdateProfile(ctx context.Context, user *domain.User) error
	GetAllForStatus() ([]domain.UserStatus, error)
}

// ==============================
// USER STATUS / PRESENCE REPOSITORY
// ==============================
type UserStatusRepository interface {
	SetOnline(ctx context.Context, userID int) error
	SetOffline(ctx context.Context, userID int) error
	GetOnlineUsers(ctx context.Context) ([]domain.UserStatus, error)
	GetStatus(ctx context.Context, userID int) (*domain.UserStatus, error)
}
