package post

import (
	"context"
	"glass-photo/internal/db"
	"log"
	"time"
)

func deletePostPostgres(ctx context.Context, targetUserID, postID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `DELETE FROM posts WHERE user_id = $1 AND id = $2`
	if _, err := db.DB.ExecContext(ctx, query, targetUserID, postID); err != nil {
		log.Printf("ERROR: deletePostPostgres - Failed to delete post: %v", err)
		return
	}
	return nil
}
