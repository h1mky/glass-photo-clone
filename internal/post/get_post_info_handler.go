package post

import (
	"github.com/go-chi/chi/v5"
	"glass-photo/internal/common"
	"net/http"
	"strconv"
)

func GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	postID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "wrong post ID", http.StatusBadRequest)
		return
	}

	post, err := GetPostByID(ctx, postID)
	if err != nil {
		http.Error(w, "unavailiable to get post", http.StatusServiceUnavailable)
		return
	}

	common.JsonEncode(w, post)
}
