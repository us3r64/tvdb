package tvdb

import (
	"context"
	"encoding/json"
)

type loginRequest struct {
	APIKey string `json:"apikey"`
	Pin    string `json:"pin,omitempty"`
}

type loginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Status string `json:"status"`
}

// Login retrieves a token to be used for TVDB API requests.
// The token is stored in the Client struct.
// https://thetvdb.github.io/v4-api/#/Login/post_login
func (c *Client) Login(ctx context.Context) error {
	reqURL, err := reqFullURL(baseURL, "login")
	if err != nil {
		return err
	}

	req := loginRequest{
		APIKey: c.apiKey,
	}
	if c.pin != "" {
		req.Pin = c.pin
	}

	resp, err := c.postRequest(ctx, reqURL, "", req)
	if err != nil {
		return err
	}

	lr := loginResponse{}
	err = json.Unmarshal(resp, &lr)
	if err != nil {
		return err
	}

	c.token = lr.Data.Token
	return nil
}
