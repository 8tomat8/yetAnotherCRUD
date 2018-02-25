package HTTPHandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/8tomat8/yetAnotherCRUD/api/HTTPHandlers/mocks"
	"github.com/8tomat8/yetAnotherCRUD/storage"
	"github.com/matryer/is"
	"github.com/stretchr/testify/mock"
)

func Test_handler_DeleteUser(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		name                  string
		request               *http.Request
		expectedStatus        int
		expectedStorageMethod string
		expectedStorageError  error
	}{
		{
			name: "Valid",
			request: httptest.NewRequest(
				http.MethodDelete,
				"/users/123",
				nil,
			),
			expectedStatus:        http.StatusNoContent,
			expectedStorageMethod: "Delete",
		},
		{
			name: "Invalid ID",
			request: httptest.NewRequest(
				http.MethodDelete,
				"/users/asd",
				nil,
			),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Id not found",
			request: httptest.NewRequest(
				http.MethodDelete,
				"/users/321",
				nil,
			),
			expectedStatus:        http.StatusNotFound,
			expectedStorageMethod: "Delete",
			expectedStorageError:  storage.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := &mocks.Storage{}
			handler := UsersHandler(storageMock, logger)
			storageMock.On("Delete", mock.Anything, mock.Anything).Return(tt.expectedStorageError)
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, tt.request)
			is.Equal(recorder.Code, tt.expectedStatus)
			is.Equal(len(storageMock.Calls), 1)

			if recorder.Code != http.StatusBadRequest {
				is.Equal(storageMock.Calls[0].Method, tt.expectedStorageMethod)
			}
		})
	}
}
