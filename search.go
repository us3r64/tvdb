package tvdb

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// SearchRequest struct for search request.
type SearchRequest struct {
	Q        string
	Query    string // additional search query param
	Type     string // restrict results to entity type movie|series|person|company
	RemoteID string // restrict results to remote id
	Year     int    // restrict results to a year for movie|series
	Offset   int
	Limit    int
}

// SearchResponse struct for search response.
type SearchResponse struct {
	Data   []SearchResult `json:"data"`
	Status string         `json:"status"`
}

// SearchResult struct for search data result response.
type SearchResult struct {
	Aliases              []string          `json:"aliases"`
	Companies            []string          `json:"companies"`
	CompanyType          string            `json:"companyType"`
	Country              string            `json:"country"`
	Director             string            `json:"director"`
	ExtendedTitle        string            `json:"extendedTitle"`
	Genres               []string          `json:"genres"`
	ID                   string            `json:"id"`
	ImageURL             string            `json:"imageUrl"`
	Name                 string            `json:"name"`
	NameTranslated       string            `json:"nameTranslated"`
	OfficialList         string            `json:"officialList"`
	Overview             string            `json:"overview"`
	OverviewTranslated   []string          `json:"overview_translated"`
	Posters              []string          `json:"posters"`
	PrimaryLanguage      string            `json:"primaryLanguage"`
	PrimaryType          string            `json:"primaryType"`
	Status               string            `json:"status"`
	TranslationsWithLang []string          `json:"translationsWithLang"`
	TvdbID               string            `json:"tvdb_id"`
	Type                 string            `json:"type"`
	Year                 string            `json:"year"`
	Thumbnail            string            `json:"thumbnail"`
	Poster               string            `json:"poster"`
	Translations         map[string]string `json:"translations"`
	IsOfficial           bool              `json:"is_official"`
	RemoteIds            []struct {
		ID         string `json:"id"`
		Type       int    `json:"type"`
		SourceName string `json:"sourceName"`
	} `json:"remoteIds"`
	Network   string            `json:"network"`
	Title     string            `json:"title"`
	Overviews map[string]string `json:"overviews"`
}

// Search queries for records based on SearchRequest params.
// https://thetvdb.github.io/v4-api/#/Search/getSearchResults
func (c *Client) Search(ctx context.Context, searchRequest SearchRequest) (SearchResponse, error) {
	sr := SearchResponse{}

	reqURL, err := reqFullURL(baseURL, "search")
	if err != nil {
		return sr, err
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return sr, err
	}

	q := u.Query()
	if searchRequest.Q != "" {
		q.Set("q", searchRequest.Q)
	}
	if searchRequest.Query != "" {
		q.Set("query", searchRequest.Query)
	}
	if searchRequest.Type != "" {
		q.Set("type", searchRequest.Type)
	}
	if searchRequest.RemoteID != "" {
		q.Set("remote_id", searchRequest.RemoteID)
	}
	if searchRequest.Year != 0 {
		q.Set("year", strconv.Itoa(searchRequest.Year))
	}
	if searchRequest.Offset != 0 {
		q.Set("offset", strconv.Itoa(searchRequest.Offset))
	}
	if searchRequest.Limit != 0 {
		q.Set("limit", strconv.Itoa(searchRequest.Limit))
	}
	u.RawQuery = q.Encode()

	resp, err := c.getRequest(ctx, u.String(), c.token)
	if err != nil {
		return sr, err
	}

	err = json.Unmarshal(resp, &sr)
	if err != nil {
		return sr, err
	}

	return sr, nil
}
