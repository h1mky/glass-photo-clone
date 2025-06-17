package post

import (
	"glass-photo/internal/common"
	"net/http"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post CreatePostRequest
	userID := r.Context().Value("user_id").(int)

	if err := common.JsonParse(w, r, &post); err != nil {
		return
	}
	
	if err := common.ValidateStruct(post); err != nil {
		http.Error(w, `{"error":"invalid post input"}`, http.StatusBadRequest)
		return
	}

	if err := createPostPostgres(ctx, post, userID); err != nil {
		http.Error(w, `{"error":"failed to create post"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
