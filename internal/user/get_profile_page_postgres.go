package user

import (
	"context"
	"glass-photo/internal/db"
	"glass-photo/internal/post"
	"time"
)

func getProfilePage(ctx context.Context, id int) (UserPage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var profile UserPage
	userQuery := `
	SELECT id, username, user_img,description
	FROM users 
	WHERE id = $1
	`
	if err := db.DB.GetContext(ctx, &profile, userQuery, id); err != nil {
		return UserPage{}, err
	}

	postsQuery := `
	SELECT p.id, p.title, p.post_img, u.username
	FROM posts AS p
	JOIN users AS u ON u.id = p.user_id
	WHERE u.id = $1
	ORDER BY p.created_at DESC
	`

	var posts []post.MainPagePost
	if err := db.DB.SelectContext(ctx, &posts, postsQuery, id); err != nil {
		return UserPage{}, err
	}

	profile.Posts = posts
	return profile, nil
}
