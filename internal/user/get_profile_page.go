package user

import (
	"context"
	"glass-photo/internal/db"
	"time"
)

func getProfilePage(ctx context.Context, id int) (UserPage, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
SELECT
    u.id,
    u.username,
    u.description
FROM users AS u
         JOIN posts AS p ON p.user_id = u.id
WHERE u.id = $1
ORDER BY p.created_at DESC;
`

	var profile UserPage
	if err := db.DB.SelectContext(ctx, &profile, query, id); err != nil {
	}
	return profile, nil
}
