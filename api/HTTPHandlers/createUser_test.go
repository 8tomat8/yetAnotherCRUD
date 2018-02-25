package HTTPHandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/8tomat8/yetAnotherCRUD/api/HTTPHandlers/mocks"
	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/matryer/is"
	"github.com/stretchr/testify/mock"
)

func Test_handler_CreateUser(t *testing.T) {
	is := is.New(t)
	storageMock := &mocks.Storage{}
	storageMock.On("Create", mock.Anything, mock.Anything).Return(nil)
	handler := UsersHandler(storageMock, logger)
	tests := []struct {
		name                  string
		request               *http.Request
		expectedStatus        int
		expectedStorageMethod string
		expectedResponse      string
	}{
		{
			name: "Valid",
			request: httptest.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewReader([]byte(`{"UserID":1,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`))),
			expectedStatus:        http.StatusCreated,
			expectedStorageMethod: "Create",
			expectedResponse:      `{"UserID":0,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`,
		},
		{
			name: "Empty fields",
			request: httptest.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewReader([]byte(`{"UserID":1,"Username":"","Password":"","Firstname":"","Lastname":"","Sex":"","Birthdate":"15-12-1991"}`))),
			expectedStatus:        http.StatusBadRequest,
			expectedStorageMethod: "Create",
		},
		{
			name: "Wrong date format",
			request: httptest.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewReader([]byte(`{"UserID":1,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"12-15-1991"}`))),
			expectedStatus:        http.StatusBadRequest,
			expectedStorageMethod: "Create",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, tt.request)
			is.Equal(recorder.Code, tt.expectedStatus)

			is.Equal(len(storageMock.Calls), 1)
			is.Equal(storageMock.Calls[0].Method, tt.expectedStorageMethod)

			if recorder.Code != http.StatusCreated {
				return
			}

			is.Equal(recorder.Body.String(), tt.expectedResponse)
			if !storageMock.Calls[0].Arguments.Get(1).(*entity.User).IsValid() {
				t.Error("saved user is invalid")
			}

		})
	}
}
