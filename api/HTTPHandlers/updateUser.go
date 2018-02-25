package HTTPHandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/8tomat8/yetAnotherCRUD/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithField("requestID", middleware.GetReqID(r.Context()))
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Info(errors.Wrap(err, "cannot parse user id"))
		w.WriteHeader(http.StatusBadRequest)
	}

	var user User

	err = json.NewDecoder(r.Body).Decode(&user)
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

	// Why int32?!
	user.UserID = int32(userID)

	userModel := &entity.User{
		UserID:    user.UserID,
		Password:  user.Password,
		Birthdate: user.Birthdate.Time,
		Sex:       user.Sex,
		Lastname:  user.Lastname,
		Firstname: user.Firstname,
		Username:  user.Username,
	}

	if !userModel.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.store.Update(r.Context(), userModel)
	if err != nil {
		if err == storage.ErrNotFound {
			log.WithField("userID", userID).Debug(err)
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Warn(errors.Wrapf(err, "cannot update user %+v", user))
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	payload, err := json.Marshal(&user)
	if err != nil {
		log.Error(errors.Wrap(err, "cannot marshal user to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = h.SendResponse(w, payload)
	if err != nil {
		log.Warn(err)
	}
}
