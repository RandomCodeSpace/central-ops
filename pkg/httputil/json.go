package httputil

import (
	"encoding/json"
	"net/http"
)

// WriteJSON sets Content-Type to application/json, writes the status code,
// and encodes v as JSON. Encoding errors are silently dropped.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// WriteError writes a JSON error response: {"error": msg}.
func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}
