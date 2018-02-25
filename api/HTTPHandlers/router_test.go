package HTTPHandlers

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func TestUsersHandler(t *testing.T) {
	type args struct {
		storage Storage
		logger  *logrus.Logger
	}
	tests := []struct {
		name string
		args args
		want chi.Router
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UsersHandler(tt.args.storage, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_objectsHandler(t *testing.T) {
	type args struct {
		storage Storage
		logger  *logrus.Logger
	}
	tests := []struct {
		name string
		args args
		want chi.Router
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := objectsHandler(tt.args.storage, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("objectsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
