package user

import (
	"context"
	"glass-photo/internal/db"
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

	return profile, nil
}
