package domain

type ReactionType string

const (
	Like    ReactionType = "like"
	Dislike ReactionType = "dislike"
	Remove  ReactionType = "remove"
)

type Reaction struct {
	Success      bool
	LikeCount    int
	DislikeCount int
	UserReaction *string
}
