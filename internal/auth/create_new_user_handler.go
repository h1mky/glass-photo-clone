package auth

import (
	"fmt"
	"glass-photo/internal/common"
	"net/http"
)

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {

	var user UserInputRegister

	if err := common.JsonParse(w, r, user); err != nil {
		fmt.Errorf("error parsing json: %v", err)
		return
	}

}
