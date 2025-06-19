package post

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func GetPostByID(ctx context.Context, postID int) (PostInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

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

	var post PostInfo
	err := db.DB.GetContext(ctx, &post, query, postID)
	if err != nil {
		return post, fmt.Errorf("error getting post: %w", err)
	}

	return post, nil
}
