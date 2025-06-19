package post

import (
	"context"
	"glass-photo/internal/db"
	"time"
)

func getUserPostsPostgres(ctx context.Context, UserID int) ([]MainPagePost, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var posts []MainPagePost

	query := `
SELECT p.id, p.title, p.post_img, u.username
	FROM posts AS p
	JOIN users AS u ON u.id = p.user_id
	WHERE u.id = $1
	ORDER BY p.created_at DESC
	`

	err := db.DB.SelectContext(ctx, &posts, query, UserID)
	if err != nil {
		return posts, err
	}
	return posts, nil
}
