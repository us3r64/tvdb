package tvdb

import (
	"context"
	"encoding/json"
	"strconv"
)

// Series struct for series base record.
type Series struct {
	Abbreviation string `json:"abbreviation"`
	Aliases      []struct {
		Language string `json:"language"`
		Name     string `json:"name"`
	} `json:"aliases"`
	Country              string   `json:"country"`
	DefaultSeasonType    int      `json:"defaultSeasonType"`
	FirstAired           string   `json:"firstAired"`
	ID                   int      `json:"id"`
	Image                string   `json:"image"`
	IsOrderRandomized    bool     `json:"isOrderRandomized"`
	LastAired            string   `json:"lastAired"`
	Name                 string   `json:"name"`
	NameTranslations     []string `json:"nameTranslations"`
	NextAired            string   `json:"nextAired"`
	OriginalCountry      string   `json:"originalCountry"`
	OriginalLanguage     string   `json:"originalLanguage"`
	OverviewTranslations []string `json:"overviewTranslations"`
	Score                float64  `json:"score"`
	Slug                 string   `json:"slug"`
	Status               struct {
		ID          int    `json:"id"`
		KeepUpdated bool   `json:"keepUpdated"`
		Name        string `json:"name"`
		RecordType  string `json:"recordType"`
	} `json:"status"`
}

// SeriesResponse struct for series response.
type SeriesResponse struct {
	Data   Series `json:"data"`
	Status string `json:"status"`
}

// GetSeries returns a series by ID.
// https://thetvdb.github.io/v4-api/#/Series/getSeriesBase
func (c *Client) GetSeries(ctx context.Context, ID int) (SeriesResponse, error) {
	sr := SeriesResponse{}

	reqURL, err := reqFullURL(baseURL, "series", strconv.Itoa(ID))
	if err != nil {
		return sr, err
	}

	resp, err := c.getRequest(ctx, reqURL, c.token)
	if err != nil {
		return sr, err
	}

	err = json.Unmarshal(resp, &sr)
	if err != nil {
		return sr, err
	}

	return sr, nil
}
