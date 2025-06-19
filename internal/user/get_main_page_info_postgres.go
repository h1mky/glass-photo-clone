package user

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func getMainPageInfoPostgres(ctx context.Context, userId int) (MainPageRegistered, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var userInfo MainPageRegistered

	query := `
SELECT id,user_img
FROM users
where id = $1
	`
	if err := db.DB.GetContext(ctx, &userInfo, query, userId); err != nil {
		return userInfo, fmt.Errorf("error getting user info: %w", err)
	}
	return userInfo, nil
}
