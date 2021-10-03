package tvdb

import (
	"context"
	"testing"
	"time"
)

func TestClient_Updates(t *testing.T) {
	since := time.Now().AddDate(0, 0, -1).Unix()

	type args struct {
		ctx            context.Context
		UpdatesRequest UpdatesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1-day-action-create", args{context.Background(), UpdatesRequest{Since: since, Type: "series", Action: "create"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.Updates(tt.args.ctx, tt.args.UpdatesRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Updates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Data) == 0 {
				t.Errorf("Client.Updates() = %v, want not empty", got)
			}
		})
	}
}
