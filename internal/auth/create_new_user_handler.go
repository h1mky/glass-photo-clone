package auth

import (
	"fmt"
	"glass-photo/internal/common"
	"net/http"
)

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user UserInputRegister
	if err := common.JsonParse(w, r, &user); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return

	}

	user.Password = common.HashPasswordSHA256(user.Password)

	if err := common.ValidateStruct(user); err != nil {
		http.Error(w, "error on validate json", http.StatusBadRequest)
		return
	}

	if err := createNewUserPostgres(ctx, user); err != nil {
		http.Error(w, fmt.Sprintf("error creating new user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
