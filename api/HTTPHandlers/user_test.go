package HTTPHandlers

import (
	"reflect"
	"testing"

	"github.com/8tomat8/yetAnotherCRUD/entity"
)

func TestCreateFromModel(t *testing.T) {
	type args struct {
		u entity.User
	}
	tests := []struct {
		name string
		args args
		want User
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateFromModel(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFromModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
