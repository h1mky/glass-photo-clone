package comment

import (
	"glass-photo/internal/common"
	"net/http"
)

func GetAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := common.GetParams(r)
	if err != nil {
		http.Error(w, "error with fetch params", http.StatusBadRequest)
		return
	}

}
