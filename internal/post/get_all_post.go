package post

import (
	"context"
	"glass-photo/internal/db"

	"log"
	"time"
)

func getAllPost(ctx context.Context) ([]Post, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	var posts []Post

	if err := db.DB.SelectContext(ctx, &posts, `SELECT * FROM posts ORDER BY created_at DESC `); err != nil {

		log.Println("Error fetching a posts", err)
		return nil, err

	}
	return posts, nil
}
