package comment

import (
	"glass-photo/internal/common"
	"net/http"
)

func deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.Context().Value("user_id").(int)

	commentID, err := common.GetParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentUserId, err := getCommentByID(ctx, commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if userID != commentUserId.UserId {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	err = deleteComment(ctx, commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
