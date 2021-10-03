package tvdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// UpdatesRequest struct for updates request.
type UpdatesRequest struct {
	Since  int64
	Type   string // Available values : artwork, award_nominees, companies, episodes, lists, people, seasons, series, seriespeople, artworktypes, award_categories, awards, company_types, content_ratings, countries, entity_types, genres, languages, movies, movie_genres, movie_status, peopletypes, seasontypes, sourcetypes, tag_options, tags, translatedcharacters, translatedcompanies, translatedepisodes, translatedlists, translatedmovies, translatedpeople, translatedseasons, translatedserierk
	Action string // Available values : delete, update
	Page   int
}

// UpdatesResponse struct for updates response.
type UpdatesResponse struct {
	Data   []UpdatesResult `json:"data"`
	Status string          `json:"status"`
	Links  struct {
		Prev string `json:"prev"`
		Self string `json:"self"`
		Next string `json:"next"`
	} `json:"links"`
}

// UpdatesResult struct for updates data result response.
type UpdatesResult struct {
	EntityType string `json:"entityType"`
	Method     string `json:"method"`
	RecordID   int    `json:"recordId"`
	TimeStamp  int    `json:"timeStamp"`
}

// Updates returns updated entities based on UpdatesRequest params.
// https://thetvdb.github.io/v4-api/#/Updates/updates
func (c *Client) Updates(ctx context.Context, updatesRequest UpdatesRequest) (UpdatesResponse, error) {
	ur := UpdatesResponse{}

	reqURL, err := reqFullURL(baseURL, "updates")
	if err != nil {
		return ur, err
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return ur, err
	}

	q := u.Query()
	if updatesRequest.Since == 0 {
		return ur, fmt.Errorf("since required")
	}
	q.Set("since", strconv.FormatInt(updatesRequest.Since, 10))
	if updatesRequest.Type != "" {
		q.Set("type", updatesRequest.Type)
	}
	if updatesRequest.Action != "" {
		q.Set("action", updatesRequest.Action)
	}
	if updatesRequest.Page != 0 {
		q.Set("page", strconv.Itoa(updatesRequest.Page))
	}
	u.RawQuery = q.Encode()

	resp, err := c.getRequest(ctx, u.String(), c.token)
	if err != nil {
		return ur, err
	}

	err = json.Unmarshal(resp, &ur)
	if err != nil {
		return ur, err
	}

	return ur, nil
}
