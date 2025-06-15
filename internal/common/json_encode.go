package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonEncode(w http.ResponseWriter, req any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //
	if err := json.NewEncoder(w).Encode(req); err != nil {

		log.Printf("Error encoding JSON: %v\n", err)

	}
}
