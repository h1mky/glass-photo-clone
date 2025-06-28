package comment

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func getAllComments(ctx context.Context, idPost int) ([]Comments, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	var comments []Comments

	query := `
SELECT c.id,c.post_id,c.user_id,c.content,c.created_at,u.username, u.user_img
FROM comments as c
JOIN users as u on c.user_id = u.id
WHERE c.post_id = $1
ORDER BY created_at DESC
	`

	if err := db.DB.SelectContext(ctx, &comments, query, idPost); err != nil {
		return comments, fmt.Errorf("error getting comments: %w", err)
	}
	return comments, nil
}
