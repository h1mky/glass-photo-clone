package comment

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func getCommentByID(ctx context.Context, CommentID int) (CommentUser, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	query := `
	SELECT 
	id,user_id,post_id,content,created_at
	FROM comments
	WHERE id = $1
	`
	var comment CommentUser

	err := db.DB.GetContext(ctx, &comment, query, CommentID)
	if err != nil {
		return comment, fmt.Errorf("error getting Comment: %w", err)
	}
	return comment, nil
}
