package middlewares

import (
	"net/http"
)

func ContentType(payload string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", payload)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
