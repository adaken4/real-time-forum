-- +goose Up
-- +goose StatementBegin
-- Sessions: For finding all sessions related to a user ID
CREATE INDEX idx_sessions_user_id ON sessions(user_id);

-- FK Indexes: Speed up joins when retrieving content/related data by user ID
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_post_reactions_user_id ON post_reactions(user_id);
CREATE INDEX idx_comment_reactions_user_id ON comment_reactions(user_id);
CREATE INDEX idx_messages_from_user_id ON messages(from_user_id);
CREATE INDEX idx_messages_to_user_id ON messages(to_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_sessions_user_id;
DROP INDEX idx_posts_user_id;
DROP INDEX idx_comments_user_id;
DROP INDEX idx_post_reactions_user_id;
DROP INDEX idx_comment_reactions_user_id;
DROP INDEX idx_messages_from_user_id;
DROP INDEX idx_messages_to_user_id;
-- +goose StatementEnd
