package repository

import (
	"context"
	"real-time-forum/internal/domain"
)

// ==============================
// MESSAGE REPOSITORY
// ==============================
type MessageRepository interface {
	Send(ctx context.Context, msg *domain.Message) (int, error)
	GetConversation(ctx context.Context, userID1, userID2 int) ([]domain.Message, error)
	MarkAsRead(ctx context.Context, messageID int) error
	MarkConversationAsRead(ctx context.Context, fromUserID, toUserID int) error
	Delete(ctx context.Context, messageID int) error
}
