package post

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func GetPostByID(ctx context.Context, postID int) ([]PostWithComments, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	query := `
        SELECT
    p.id AS post_id,
    p.title,
    p.post_img,
    p.created_at,
    u.username AS post_author,
    u.id AS post_author_id
FROM
    posts AS p
    JOIN users AS u ON p.user_id = u.id
WHERE
    p.id = $1
`

	var post []PostWithComments
	err := db.DB.SelectContext(ctx, &post, query, postID)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %w", err)
	}

	return post, nil
}
