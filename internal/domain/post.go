package domain

import "time"

type Post struct {
	ID           int
	UserID       int
	Author       string
	Title        string
	Content      string
	Category     string
	FileURL      *string
	CommentCount int
	LikeCount    int
	DislikeCount int
	UserReaction *string // 'like', 'dislike', or null
	CreatedAt    time.Time
}

type PostDetail struct {
	Post     Post
	Comments []Comment
}
