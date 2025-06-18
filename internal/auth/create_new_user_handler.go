package auth

import (
	"fmt"
	"glass-photo/internal/common"
	"net/http"
)

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user UserInputRegister
	defaultUserPhoto := "https://www.shutterstock.com/image-vector/blank-avatar-photo-place-holder-600nw-1095249842.jpg"

	if err := common.JsonParse(w, r, &user); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := common.ValidateStruct(user); err != nil {
		http.Error(w, fmt.Sprintf("invalid struct: %v", err), http.StatusBadRequest)
		return
	}

	user.Password = common.HashPasswordSHA256(user.Password)
	user.UserPhoto = defaultUserPhoto

	if err := createNewUserPostgres(ctx, user); err != nil {
		http.Error(w, fmt.Sprintf("error creating new user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
