package post

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func createPostPostgres(ctx context.Context, post CreatePostRequest, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO posts (user_id, post_img, title, description)
	VALUES ($1, $2, $3, $4)
	`
	_, err := db.DB.ExecContext(ctx, query, userID, post.PostImg, post.Title, post.Description)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}
