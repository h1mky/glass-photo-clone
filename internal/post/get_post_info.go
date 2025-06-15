package post

import (
	"fmt"
	"glass-photo/internal/db"
)

func GetPostByID(postID int) ([]PostWithComments, error) {
	query := `
        SELECT
            p.id AS post_id,
            p.title,
            p.post_img,
            p.created_at,
            u.username AS post_author,
            c.content AS comment_text,
            c.created_at AS comment_created_at,
            cu.id AS comment_user_id,
            cu.username AS comment_author
        FROM
            posts AS p
            JOIN users AS u ON p.user_id = u.id
            LEFT JOIN comments AS c ON c.post_id = p.id
            LEFT JOIN users AS cu ON c.user_id = cu.id
        ORDER BY CREATED_AT DESC
        WHERE
            p.id = $1
            `

	var post []PostWithComments
	err := db.DB.Select(&post, query, postID)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %w", err)
	}
	return post, nil
}
