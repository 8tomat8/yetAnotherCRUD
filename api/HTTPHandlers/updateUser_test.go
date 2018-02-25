package HTTPHandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"github.com/8tomat8/yetAnotherCRUD/api/HTTPHandlers/mocks"
	"github.com/8tomat8/yetAnotherCRUD/storage"
	"github.com/matryer/is"
	"github.com/stretchr/testify/mock"
)

func Test_handler_UpdateUser(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		name                  string
		request               *http.Request
		expectedStatus        int
		expectedStorageMethod string
		expectedStorageError  error
		expectedResponse      string
	}{
		{
			name: "Valid",
			request: httptest.NewRequest(
				http.MethodPut,
				"/users/123",
				bytes.NewReader([]byte(`{"UserID":1,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`)),
			),
			expectedStatus:        http.StatusOK,
			expectedStorageMethod: "Update",
			expectedResponse:      `{"UserID":123,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`,
		},
		{
			name: "Invalid ID",
			request: httptest.NewRequest(
				http.MethodPut,
				"/users/asd",
				nil,
			),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Id not found",
			request: httptest.NewRequest(
				http.MethodPut,
				"/users/321",
				bytes.NewReader([]byte(`{"UserID":1,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`)),
			),
			expectedStatus:        http.StatusNotFound,
			expectedStorageMethod: "Update",
			expectedStorageError:  storage.ErrNotFound,
		},
		{
			name: "Empty body",
			request: httptest.NewRequest(
				http.MethodPut,
				"/users/321",
				nil,
			),
			expectedStatus:        http.StatusBadRequest,
			expectedStorageMethod: "Update",
		},
		{
			name: "Invalid body - No username",
			request: httptest.NewRequest(
				http.MethodPut,
				"/users/321",
				bytes.NewReader([]byte(`{"UserID":1,"Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}`)),
			),
			expectedStatus:        http.StatusBadRequest,
			expectedStorageMethod: "Update",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := &mocks.Storage{}
			handler := UsersHandler(storageMock, logger)
			storageMock.On("Update", mock.Anything, mock.Anything).Return(tt.expectedStorageError)
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, tt.request)
			is.Equal(recorder.Code, tt.expectedStatus)

			if recorder.Code != http.StatusBadRequest {
				is.Equal(storageMock.Calls[0].Method, tt.expectedStorageMethod)
				is.Equal(len(storageMock.Calls), 1)
			}
		})
	}
}
