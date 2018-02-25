package HTTPHandlers

import (
	"io"
	"net/http"
	"testing"
)

func Test_handler_SendResponse(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		payload []byte
	}
	tests := []struct {
		name    string
		h       handler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.SendResponse(tt.args.w, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("handler.SendResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handler_CloseBody(t *testing.T) {
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name    string
		h       handler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CloseBody(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("handler.CloseBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
