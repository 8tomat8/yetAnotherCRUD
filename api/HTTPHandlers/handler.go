package HTTPHandlers

import (
	"net/http"

	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type handler struct {
	store  Storage
	logger *logrus.Logger
}

func (handler) SendResponse(w http.ResponseWriter, payload []byte) error {
	n, err := w.Write(payload)
	if err != nil {
		return err
	} else if n != len(payload) {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("not all data was written to response")
	}
	return nil
}

func (handler) CloseBody(body io.ReadCloser) error {
	_, err := io.Copy(ioutil.Discard, body)
	if err != nil {
		return errors.Wrap(err, "cannot clean body")
	}

	return errors.Wrap(body.Close(), "cannot close body")
}
