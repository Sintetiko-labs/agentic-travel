package client

import ("bytes"; "fmt"; "io"; "net/http"; "github.com/fbelchi/travelkit/akamai"; tkbase "github.com/fbelchi/travelkit/base")

func (c *Client) apiGET(fullURL, referer string) ([]byte, int, error) {
	c.Throttle(); req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil { return nil, 0, err }
	c.SetAPIHeaders(req); if referer != "" { req.Header.Set("referer", referer) }
	req.Header.Set("origin", c.BaseURL); c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req); if err != nil { return nil, 0, err }
	defer resp.Body.Close(); body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20)); return body, resp.StatusCode, nil
}
func (c *Client) apiPOST(fullURL, referer string, payload []byte) ([]byte, int, error) {
	c.Throttle(); req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewReader(payload))
	if err != nil { return nil, 0, err }
	c.SetAPIHeaders(req); req.Header.Set("content-type", "application/json")
	if referer != "" { req.Header.Set("referer", referer) }
	req.Header.Set("origin", c.BaseURL); c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req); if err != nil { return nil, 0, err }
	defer resp.Body.Close(); body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20)); return body, resp.StatusCode, nil
}
func checkAPI(status int, body []byte, slug string) error {
	text := string(body)
	if akamai.IsWAFBlocked(status, text) || akamai.IsDenied(status, text) { return fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint(slug)) }
	if status < 200 || status >= 300 { return fmt.Errorf("HTTP %d: %s", status, tkbase.Truncate(text, 200)) }
	if !akamai.LooksLikeJSON(text) { return fmt.Errorf("non-JSON response — %s", akamai.NeedsSessionHint(slug)) }
	return nil
}
