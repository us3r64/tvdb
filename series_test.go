package tvdb

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
)

func TestClient_GetSeries(t *testing.T) {
	fireflySeries := `{
		"status": "success",
		"data": {
		  "id": 78874,
		  "name": "Firefly",
		  "slug": "firefly",
		  "image": "https://artworks.thetvdb.com/banners/posters/78874-2.jpg",
		  "nameTranslations": ["ces","deu","eng","fra","heb","hun","ita","kor","nld","pol","por","rus","spa","swe","zho","hrv"],
		  "overviewTranslations": ["ces","deu","eng","fra","heb","hun","ita","nld","pol","por","rus","spa","zho","hrv"],
		  "aliases": [
			{
			  "language": "deu",
			  "name": "Firefly - Der Aufbruch der Serenity"
			},
			{
			  "language": "eng",
			  "name": "Serenity"
			}
		  ],
		  "firstAired": "2002-09-20",
		  "lastAired": "2003-07-28",
		  "nextAired": "",
		  "score": 9.5,
		  "status": {
			"id": 2,
			"name": "Ended",
			"recordType": "series",
			"keepUpdated": false
		  },
		  "originalCountry": "usa",
		  "originalLanguage": "eng",
		  "defaultSeasonType": 1,
		  "isOrderRandomized": false,
		  "lastUpdated": "2020-01-26 14:49:00",
		  "averageRuntime": 48
		}
	}`

	var fireflySeriesResponse SeriesResponse
	err := json.Unmarshal([]byte(fireflySeries), &fireflySeriesResponse)
	if err != nil {
		t.Fatalf("Client.GetSeries() error = %v", err)
	}

	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    SeriesResponse
		wantErr bool
	}{
		{"no-ID", args{context.Background(), 0}, SeriesResponse{}, true},
		{"ok-ID-Firefly", args{context.Background(), 78874}, fireflySeriesResponse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.GetSeries(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetSeries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}
