package common

import (
	"encoding/json"
	"net/http"
)

func JsonEncode(w http.ResponseWriter, req any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
