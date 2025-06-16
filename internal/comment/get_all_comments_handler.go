package comment

import (
	"glass-photo/internal/common"
	"net/http"
)

func getAllCommentsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	params, err := common.GetParams(r)
	if err != nil {
		http.Error(w, "error with fetch params", http.StatusBadRequest)
		return
	}

	comments, err := getAllComments(ctx, params)
	if err != nil {
		http.Error(w, "error with fetch comments", http.StatusBadRequest)
		return
	}
	common.JsonEncode(w, comments)

}
