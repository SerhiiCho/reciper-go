package middleware

import (
	"net/http"
)

// AppMiddleware sets http headers
func AppMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Max-Age", "604800")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
