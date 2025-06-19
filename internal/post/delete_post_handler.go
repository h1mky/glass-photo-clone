package post

import (
	"fmt"
	"glass-photo/internal/common"
	"log"
	"net/http"
)

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("user_id").(int)
	postID, err := common.GetParams(r)
	if err != nil {
		log.Printf("ERROR: deletePostHandler - Get post ID from request failed: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := deletePostPostgres(ctx, userID, postID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete post: %v", err), http.StatusInternalServerError)
		return
	}

	http.ResponseWriter(w).WriteHeader(http.StatusNoContent)

}
