package tvdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"
)

const baseURL = "https://api4.thetvdb.com/v4"

// MaxRequestPerSecond max requests per second
// Client can do to the TVDB API.
const MaxRequestPerSecond = 4

var (
	rate     = time.Second/MaxRequestPerSecond + time.Millisecond*20
	throttle = time.NewTicker(rate)
)

// Client used to do requests to the TVDB API (https://thetvdb.github.io/v4-api/).
type Client struct {
	apiKey string
	pin    string
	token  string
	client *http.Client
}

// New provisions Client with TVDB API KEY, PIN and custom http.Client values.
func New(apiKey, pin string) *Client {
	return &Client{
		apiKey: apiKey,
		pin:    pin,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 1024,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 60 * time.Second,
				}).DialContext,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   15 * time.Second,
				ResponseHeaderTimeout: 30 * time.Second,
				ExpectContinueTimeout: 15 * time.Second,
			},
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) getRequest(ctx context.Context, url, token string) ([]byte, error) {
	<-throttle.C

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 { // Success!
		return body, nil
	}

	return body, fmt.Errorf("getRequest failed with status: %d", resp.StatusCode)
}

func (c *Client) postRequest(ctx context.Context, url, token string, reqBody interface{}) ([]byte, error) {
	<-throttle.C

	json, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 { // Success!
		return body, nil
	}

	return body, fmt.Errorf("postRequest failed with status: %d", resp.StatusCode)
}

func reqFullURL(baseURL string, elem ...string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	elem = append([]string{u.Path}, elem...)
	u.Path = path.Join(elem...)
	return u.String(), nil
}
