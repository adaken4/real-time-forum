package domain

import "time"

type Comment struct {
	ID           int
	PostID       int
	UserID       int
	Author       string
	Content      string
	FileURL      *string
	LikeCount    int
	DislikeCount int
	UserReaction *string // 'like', 'dislike', or null
	CreatedAt    time.Time
}
