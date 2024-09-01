package client

import (
	"net/http"
	"time"

	"github.com/johnconnor-sec/firecrawl/pkg/config"
)

type Client struct {
	http *http.Client
	cfg  *config.Config
}

func New(cfg *config.Config) *Client {
	return &Client{
		http: &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second},
		cfg:  cfg,
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.cfg.UserAgent)
	return c.http.Do(req)
}