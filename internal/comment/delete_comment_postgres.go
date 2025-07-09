package comment

import (
	"context"
	"glass-photo/internal/db"
	"log"
	"time"
)

func deleteComment(ctx context.Context, commentId int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `DELETE FROM comments WHERE id = $1;`

	_, err = db.DB.ExecContext(ctx, query, commentId)
	if err != nil {
		log.Printf("ERROR: deleteCommentPostgres - Failed to delete Comment: %v", err)
		return
	}
	return nil

}
