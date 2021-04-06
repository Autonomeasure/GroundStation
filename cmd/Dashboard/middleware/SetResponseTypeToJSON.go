package middleware

import "net/http"

// Set the HTTP Content-Type header to application/json for the API routes
func SetResponseTypeToJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}