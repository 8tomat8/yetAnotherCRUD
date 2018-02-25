package HTTPHandlers

import (
	"context"

	"github.com/8tomat8/yetAnotherCRUD/api/middlewares"
	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Storage interface {
	Create(context.Context, *entity.User) error
	Delete(context.Context, int) error
	Update(context.Context, *entity.User) error
	Search(ctx context.Context, username, sex *string, age *int) ([]entity.User, error)
}

// apiHandler will handle all requests for /api path
func UsersHandler(storage Storage, logger *logrus.Logger) chi.Router {
	r := chi.NewRouter()

	r.Mount("/users", objectsHandler(storage, logger))
	return r
}

// objectsHandler will handle all requests for /objects path
func objectsHandler(storage Storage, logger *logrus.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.ContentType("application/json; charset=UTF-8"))

	HTTPHandler := handler{store: storage, logger: logger}

	r.Post("/", HTTPHandler.CreateUser)
	r.Get("/", HTTPHandler.SearchUser)
	r.Put("/{id}", HTTPHandler.UpdateUser)
	r.Delete("/{id}", HTTPHandler.DeleteUser)
	return r
}
