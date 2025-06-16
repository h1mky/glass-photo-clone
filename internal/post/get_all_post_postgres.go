package post

import (
	"context"
	"glass-photo/internal/db"
	"log"
	"time"
)

func getAllPost(ctx context.Context) ([]MainPagePost, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	var posts []MainPagePost

	query := `
		SELECT
			p.id,
			p.title,
			p.post_img,
			u.username
		FROM posts AS p
		JOIN users AS u ON p.user_id = u.id
		ORDER BY p.created_at DESC
	`
	if err := db.DB.SelectContext(ctx, &posts, query); err != nil {

		log.Println("Error fetching a posts", err)
		return nil, err

	}
	return posts, nil
}
