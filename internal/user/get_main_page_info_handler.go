package user

import (
	"glass-photo/internal/common"
	"net/http"
)

func GetMainPageInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.Context().Value("user_id").(int)

	userInfo, err := getMainPageInfoPostgres(ctx, userID)
	if err != nil {
		http.Error(w, "failed to get profile", http.StatusBadRequest)
		return
	}
	common.JsonEncode(w, userInfo)
}
