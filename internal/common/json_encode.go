package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonEncode(w http.ResponseWriter, posts any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v\n", err)
	}
}
