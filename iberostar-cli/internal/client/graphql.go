package client

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	graphQLPath    = "/api/graphql"
	graphQLReferer = "/es/hoteles"
)

func (c *Client) postGraphQL(payload, out any) error {
	c.Throttle()
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+graphQLPath, strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", c.BaseURL)
	req.Header.Set("referer", c.BaseURL+graphQLReferer)
	req.Header.Set("x-market", "ES")
	req.Header.Set("x-language", "es")
	c.ApplyCookie(req)
	return c.DoJSON(req, out)
}
