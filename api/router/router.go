package router

import (
	"time"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	defaultTimeout = 60
)

func NewRouter(handlers ...http.Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer) // Could be improved to use custom lib to log errors
	r.Use(middleware.Timeout(defaultTimeout * time.Second))
	r.Use(middleware.RequestID)

	for _, handler := range handlers {
		r.Mount("/api", handler)
	}

	return r
}
