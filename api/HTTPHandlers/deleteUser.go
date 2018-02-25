package HTTPHandlers

import (
	"net/http"

	"strconv"

	"github.com/8tomat8/yetAnotherCRUD/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

// DeleteObject handles requests for objects deletion
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithField("requestID", middleware.GetReqID(r.Context()))

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Info(errors.Wrap(err, "cannot parse user id"))
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.store.Delete(r.Context(), userID)
	if err != nil {
		if err == storage.ErrNotFound {
			log.WithField("userID", userID).Debug(err)
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
