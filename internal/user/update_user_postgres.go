package user

import (
	"context"
	"database/sql"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func updateUserAvatarAndDescriptionPostgres(ctx context.Context, userID int, patch UserPatchRequest) (UserPage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var query string
	var args []interface{}

	switch {
	case patch.UserImg != nil && patch.Description != nil:
		query = `
			UPDATE users 
			SET user_img = $1, description = $2
			WHERE id = $3
			RETURNING id, username, user_img, description
		`
		args = []interface{}{patch.UserImg, patch.Description, userID}

	case patch.UserImg != nil:
		query = `
			UPDATE users 
			SET user_img = $1
			WHERE id = $2
			RETURNING id, username, user_img, description
		`
		args = []interface{}{patch.UserImg, userID}

	case patch.Description != nil:
		query = `
			UPDATE users 
			SET description = $1
			WHERE id = $2
			RETURNING id, username, user_img, description
		`
		args = []interface{}{patch.Description, userID}

	default:
		return UserPage{}, fmt.Errorf("no fields provided for update (user_img or description)")
	}

	var updatedProfile UserPage
	err := db.DB.GetContext(ctx, &updatedProfile, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserPage{}, sql.ErrNoRows
		}
		return UserPage{}, fmt.Errorf("failed to update user avatar/description in db: %w", err)
	}

	return updatedProfile, nil
}
