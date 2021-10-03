package tvdb

import (
	"context"
	"testing"
)

func TestClient_Login(t *testing.T) {
	type fields struct {
		apikey string
		pin    string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"no key, no pin", fields{}, args{context.Background()}, true},
		{"key, no pin", fields{testAPIKey, ""}, args{context.Background()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.apikey, tt.fields.pin)
			if err := c.Login(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Client.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
