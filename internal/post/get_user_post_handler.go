package post

import (
	"glass-photo/internal/common"
	"net/http"
)

func getUserPostHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID, err := common.GetParams(r)
	if err != nil {
		http.Error(w, "failed to get params", http.StatusBadRequest)
		return
	}

	posts, err := getUserPostsPostgres(ctx, userID)
	if err != nil {
		http.Error(w, "failed to get profile", http.StatusBadRequest)
		return
	}
	common.JsonEncode(w, posts)
}
