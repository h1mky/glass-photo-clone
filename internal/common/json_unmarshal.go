package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonParse(w http.ResponseWriter, r *http.Request, req interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"failed to read body"}`, http.StatusBadRequest)
		return err
	}

	if err := json.Unmarshal(body, req); err != nil {
		http.Error(w, `{"error":"invalid JSON format"}`, http.StatusBadRequest)
		return err
	}

	return nil
}
