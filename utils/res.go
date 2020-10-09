package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON will return data as JSON
func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahKeByte, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahKeByte))
}
