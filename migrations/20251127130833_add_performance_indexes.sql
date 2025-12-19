-- +goose Up
-- +goose StatementBegin
-- Messages
CREATE INDEX idx_messages_users ON messages(from_user_id, to_user_id, created_at);
CREATE INDEX idx_messages_created_at ON messages(created_at DESC);

-- Sessions
CREATE UNIQUE INDEX idx_sessions_token ON sessions(session_token); 
-- Note: Tokens should likely be unique for a primary lookup

-- Posts & Comments
CREATE INDEX idx_posts_category ON posts(category, created_at DESC);
CREATE INDEX idx_comments_post_id ON comments(post_id, created_at ASC);

-- Reactions
CREATE INDEX idx_post_reactions ON post_reactions(post_id, reaction_type);
CREATE INDEX idx_comment_reactions ON comment_reactions(comment_id, reaction_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Messages
DROP INDEX idx_messages_users;
DROP INDEX idx_messages_created_at;

-- Sessions
DROP INDEX idx_sessions_token;

-- Posts & Comments
DROP INDEX idx_posts_category;
DROP INDEX idx_comments_post_id;

-- Reactions
DROP INDEX idx_post_reactions;
DROP INDEX idx_comment_reactions;
-- +goose StatementEnd
