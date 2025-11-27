package domain

import "time"

type Message struct {
	ID         int
	FromUserID int
	ToUserID   int
	Content    string
	FileURL    *string
	ReadStatus bool
	CreatedAt  time.Time
}
