package repository

import (
	"context"
	"real-time-forum/internal/domain"
)

// ==============================
// POST REPOSITORY
// ==============================
type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) (int, error)
	GetAll(ctx context.Context, category string, currentUserId int) ([]domain.Post, error)
	
	// Detail view
	GetByID(ctx context.Context, id, currentUserId int) (*domain.PostDetail, error)

	// Basic retrieval for create/update workflows
	GetPostByID(ctx context.Context, id, currentUserId int) (*domain.Post, error)

	// Reactions
	UpdateReaction(ctx context.Context, postID, userID int, reactionType string) error
	GetReactionCounts(ctx context.Context, postID int) (int, int, error)
	GetUserReaction(ctx context.Context, postID, userID int) (string, error)

	// Metrics
	IncrementCommentCount(ctx context.Context, postID int) error
	DecrementCommentCount(ctx context.Context, postID int) error
}
