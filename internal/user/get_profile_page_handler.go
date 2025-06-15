package user

import (
	"glass-photo/internal/common"
	"net/http"
)

func getProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, err := common.GetParams(r)
	if err != nil {
		http.Error(w, "failed to get params", http.StatusBadRequest)
		return
	}
	profile, err := getProfilePage(ctx, params)
	if err != nil {
		http.Error(w, "failed to get profile", http.StatusBadRequest)
		return
	}
	common.JsonEncode(w, profile)
}
