package auth

import (
	"context"
	"glass-photo/internal/db"
	"time"
)

func createNewUserPostgres(ctx context.Context, register UserInputRegister) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	query := `
INSERT INTO users (username, password,email) VALUES (:username, :password, :email)
		`

	if _, err := db.DB.NamedExecContext(ctx, query, register); err != nil {
		return err
	}
	return nil
}
