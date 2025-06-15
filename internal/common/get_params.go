package common

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetParams(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.Atoi(idStr)
}
