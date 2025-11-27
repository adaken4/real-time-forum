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
