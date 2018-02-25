package middlewares

import (
	"net/http"
)

func ContentType(payload string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			w.Header().Set("Content-Type", payload)
		}
		return http.HandlerFunc(fn)
	}
}
