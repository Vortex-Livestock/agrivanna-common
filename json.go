package common

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WirteError(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, map[string]string{"error": message})
}
