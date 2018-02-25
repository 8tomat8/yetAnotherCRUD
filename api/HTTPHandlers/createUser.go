package HTTPHandlers

import (
	"net/http"

	"encoding/json"

	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

// CreateUser handles requests for creating new users
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithField("requestID", middleware.GetReqID(r.Context()))

	var userReq User
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		log.Info(errors.Wrap(err, "cannot unmarshal user from body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func() {
		err := h.CloseBody(r.Body)
		if err != nil {
			log.Error(err)
		}
	}()

	user := entity.User{
		Password:  userReq.Password,
		Birthdate: userReq.Birthdate.Time,
		Sex:       userReq.Sex,
		Lastname:  userReq.Lastname,
		Firstname: userReq.Firstname,
		Username:  userReq.Username,
	}

	if !user.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.store.Create(r.Context(), &user)
	if err != nil {
		log.Warn(errors.Wrapf(err, "cannot create user %+v", user))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(CreateFromModel(user))
	if err != nil {
		log.Error(errors.Wrap(err, "cannot marshal user to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = h.SendResponse(w, payload)
	if err != nil {
		log.Warn(err)
	}
}
