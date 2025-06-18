package auth

import (
	"context"
	"fmt"
	"glass-photo/internal/db"
	"time"
)

func createNewUserPostgres(ctx context.Context, register UserInputRegister) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	query := `
INSERT INTO users (username, password,email,user_img) VALUES (:username, :password, :email, :user_img)
		`

	if _, err := db.DB.NamedExecContext(ctx, query, register); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
