package post

import (
	"glass-photo/internal/common"
	"net/http"
)

func getAllPostHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	posts, err := getAllPost(ctx)
	if err != nil {
		http.Error(w, "failed to fetch all posts", http.StatusServiceUnavailable)
		return
	}

	common.JsonEncode(w, posts)

}
