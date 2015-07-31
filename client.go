package steam

import (
	"fmt"
	"net/http"
)

// a steam API client, not tied to any particular game
type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key: key}
}

func (c *Client) Get(iface, method, version string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.steampowered.com/%s/%s/%s/?key=%s", iface, method, version, c.key)
	return http.Get(url)
}
