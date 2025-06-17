package auth

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func signInPostgres(ctx context.Context, user UserInputSignIn) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var id int
	
	err := db.DB.GetContext(ctx, &id, `
		SELECT id FROM users WHERE email = $1 AND password = $2
	`, user.Email, user.Password)

	if err != nil {
		return 0, fmt.Errorf("auth check failed: %w", err)
	}

	return id, nil
}
