package HTTPHandlers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

func (h *handler) SearchUser(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithField("requestID", middleware.GetReqID(r.Context()))

	var (
		err error

		usernameParam *string
		sexParam      *string
		ageParam      *int
	)

	queryValues := r.URL.Query()
	username, ok := queryValues["username"]
	if ok && len(username) >= 0 {
		usernameParam = &username[0]
	}

	sex, ok := queryValues["sex"]
	if ok && len(sex) >= 0 {
		sexParam = &sex[0]
	}

	age, ok := queryValues["age"]
	if ok && len(age) >= 0 {
		ageInt, err := strconv.Atoi(age[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ageParam = &ageInt
	}

	users, err := h.store.Search(r.Context(), usernameParam, sexParam, ageParam)
	if err != nil {
		log.Warn(errors.Wrap(err, "cannot search users"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersResp := make([]User, len(users))
	for i, user := range users {
		usersResp[i] = CreateFromModel(user)
	}

	payload, err := json.Marshal(usersResp)
	if err != nil {
		log.Error(errors.Wrap(err, "cannot marshal users to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = h.SendResponse(w, payload)
	if err != nil {
		log.Warn(err)
	}
}
