package comment

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func postCommentPostgres(ctx context.Context, comment CommentPost, userID, postID int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	query := `
INSERT INTO comments (user_id, post_id, content, created_at)
VALUES ($1, $2, $3, NOW())`

	if _, err := db.DB.ExecContext(ctx, query, userID, postID, comment.Content); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}
