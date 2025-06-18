package comment

import (
	"glass-photo/internal/common"
	"net/http"
)

func postCommentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var comment CommentPost

	postID, err := common.GetParams(r)

	if err != nil {
		return
	}

	userID := r.Context().Value("user_id").(int)

	if err := common.JsonParse(w, r, &comment); err != nil {
		return
	}

	if err := common.ValidateStruct(comment); err != nil {
		http.Error(w, "error on validate json", http.StatusBadRequest)
		return
	}
	if err := postCommentPostgres(ctx, comment, userID, postID); err != nil {
		http.Error(w, "failed to create comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
