package HTTPHandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/8tomat8/yetAnotherCRUD/api/HTTPHandlers/mocks"
	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/matryer/is"
	"github.com/stretchr/testify/mock"
)

func Test_handler_SearchUser(t *testing.T) {
	storageMock := &mocks.Storage{}
	handler := UsersHandler(storageMock, logger)
	storageMock.On("Search", mock.Anything, (*string)(nil), (*string)(nil), (*int)(nil)).Return(storageUsers, nil).Once()
	storageMock.On("Search", mock.Anything, (*string)(nil), stringPtr("male"), (*int)(nil)).Return([]entity.User{storageUsers[0]}, nil).Once()
	storageMock.On("Search", mock.Anything, stringPtr("Username433"), (*string)(nil), (*int)(nil)).Return([]entity.User{storageUsers[2]}, nil).Once()
	storageMock.On("Search", mock.Anything, (*string)(nil), (*string)(nil), intPtr(26)).Return([]entity.User{storageUsers[1]}, nil).Once()

	is := is.New(t)
	tests := []struct {
		name                  string
		request               *http.Request
		expectedStatus        int
		expectedStorageMethod string
		expectedStorageError  error
		expectedStorageData   []entity.User
		expectedResponse      string
	}{
		{
			name: "All users",
			request: httptest.NewRequest(
				http.MethodGet,
				"/users",
				nil,
			),
			expectedStatus:        http.StatusOK,
			expectedStorageMethod: "Search",
			expectedStorageData:   storageUsers,
			expectedResponse:      `[{"UserID":431,"Username":"Username431","Password":"SuperMegaPassword431","Firstname":"Firstname431","Lastname":"Lastname431","Sex":"male","Birthdate":"15-12-1990"},{"UserID":432,"Username":"Username432","Password":"SuperMegaPassword432","Firstname":"Firstname432","Lastname":"Lastname432","Sex":"female","Birthdate":"15-12-1991"},{"UserID":433,"Username":"Username433","Password":"SuperMegaPassword433","Firstname":"Firstname433","Lastname":"Lastname433","Sex":"female","Birthdate":"15-12-1992"}]`,
		},
		{
			name: "Filter by sex",
			request: httptest.NewRequest(
				http.MethodGet,
				"/users?sex=male",
				nil,
			),
			expectedStatus:        http.StatusOK,
			expectedStorageMethod: "Search",
			expectedStorageData:   []entity.User{storageUsers[0]},
			expectedResponse:      `[{"UserID":431,"Username":"Username431","Password":"SuperMegaPassword431","Firstname":"Firstname431","Lastname":"Lastname431","Sex":"male","Birthdate":"15-12-1990"}]`,
		},
		{
			name: "Filter by username",
			request: httptest.NewRequest(
				http.MethodGet,
				"/users?username=Username433",
				nil,
			),
			expectedStatus:        http.StatusOK,
			expectedStorageMethod: "Search",
			expectedStorageData:   []entity.User{storageUsers[2]},
			expectedResponse:      `[{"UserID":433,"Username":"Username433","Password":"SuperMegaPassword433","Firstname":"Firstname433","Lastname":"Lastname433","Sex":"female","Birthdate":"15-12-1992"}]`,
		},
		{
			name: "Filter by age",
			request: httptest.NewRequest(
				http.MethodGet,
				"/users?age=26",
				nil,
			),
			expectedStatus:        http.StatusOK,
			expectedStorageMethod: "Search",
			expectedStorageData:   []entity.User{storageUsers[1]},
			expectedResponse:      `[{"UserID":432,"Username":"Username432","Password":"SuperMegaPassword432","Firstname":"Firstname432","Lastname":"Lastname432","Sex":"female","Birthdate":"15-12-1991"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, tt.request)
			is.Equal(recorder.Code, tt.expectedStatus)

			is.Equal(storageMock.Calls[0].Method, tt.expectedStorageMethod)
			is.Equal(recorder.Body.String(), tt.expectedResponse)
		})
	}
	storageMock.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
func intPtr(s int) *int {
	return &s
}
