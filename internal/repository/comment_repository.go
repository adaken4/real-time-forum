package repository

import (
	"context"
	"real-time-forum/internal/domain"
)

// ==============================
// COMMENT REPOSITORY
// ==============================
type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) (int, error)
	GetByPostID(ctx context.Context, postID, currentUserID int) ([]domain.Comment, error)

	// Reactions
	UpdateReaction(ctx context.Context, commentID, userID int, reactionType string) error
	GetReactionCounts(ctx context.Context, commentID int) (likeCount, dislikeCount int, err error)
	GetUserReaction(ctx context.Context, commentID, userID int) (string, error)
}
