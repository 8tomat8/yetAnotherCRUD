package middlewares

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	"github.com/8tomat8/Qm9yeXMtSHVsaWk-/logger"
)

func init() {
	logger.Log.Out = ioutil.Discard // Logs redirected to /dev/null
}

type handlerMock struct {
	requestHandler func(w http.ResponseWriter, r *http.Request)
}

func (h handlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.requestHandler(w, r)
}

func TestContentType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		responseCode int
		ctxValue     string
		request      http.Request
		recorder     *httptest.ResponseRecorder
	}{
		{
			"Without Content-Type",
			http.StatusUnsupportedMediaType,
			"",
			http.Request{
				Method: "POST",
				Header: http.Header{ContentTypeKey: []string{""}},
			},
			httptest.NewRecorder(),
		},
		{
			"Valid Content-Type",
			http.StatusOK,
			"",
			http.Request{
				Method: "POST",
				Header: http.Header{ContentTypeKey: []string{"application/json"}},
			},
			httptest.NewRecorder(),
		},
	}

	for _, c := range testCases {
		c := c // To avoid data races
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			var request *http.Request
			h := handlerMock{func(_ http.ResponseWriter, r *http.Request) {
				request = r
			}}

			resHandler := ContentType(h)

			resHandler.ServeHTTP(c.recorder, c.request.WithContext(context.Background()))

			if c.recorder.Code != c.responseCode {
				t.Errorf("Expected HTTP status code %d, got %d", c.responseCode, c.recorder.Code)
			}

			// To filter negative test cases
			if c.recorder.Code != http.StatusOK {
				return
			}

			resultMediaType := GetMediaType(request)
			if resultMediaType != c.request.Header[ContentTypeKey][0] {
				t.Errorf(`Expected media type "%s", got "%s"`, c.request.Header[ContentTypeKey][0], resultMediaType)
			}
		})
	}
}
