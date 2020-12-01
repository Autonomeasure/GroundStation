package middleware

import "net/http"

// Set the HTTP Access-Control-Allow-Origin header to '*'
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
