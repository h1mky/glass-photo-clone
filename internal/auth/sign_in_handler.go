package auth

import (
	"fmt"
	"glass-photo/internal/common"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user UserInputSignIn

	if err := common.JsonParse(w, r, &user); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := common.ValidateStruct(user); err != nil {
		http.Error(w, fmt.Sprintf("invalid struct: %v", err), http.StatusBadRequest)
		return
	}

	user.Password = common.HashPasswordSHA256(user.Password)

	userId, err := signInPostgres(ctx, user)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to sign in: %v", err), http.StatusBadRequest)
		return
	}

	token, err := common.GenerateJWT(userId, user.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generate token: %v", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}
