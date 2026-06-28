// Package base provides a reusable HTTP client skeleton for travel CLIs.
package base

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/cookies"
	"github.com/fbelchi/travelkit/ratelimit"
	"github.com/fbelchi/travelkit/session"
	"github.com/fbelchi/travelkit/transport"
)

const DefaultUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

// Client is the shared HTTP transport used by scaffolded travel CLIs.
type Client struct {
	HTTP      *http.Client
	Jar       cookies.Jar
	BaseURL   string
	UserAgent string
	Cookie    string
	Pacer     *ratelimit.Pacer
	EnvPrefix string
}

// New builds a client with Chrome-like TLS and optional env-based rate limiting.
func New(baseURL, envPrefix string) *Client {
	jar := cookies.NewJar()
	hc := &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
	}
	if tr, err := transport.NewChromeTransport(); err == nil {
		hc.Transport = tr
	}
	prefix := strings.ToUpper(strings.ReplaceAll(envPrefix, "-", "_"))
	if prefix == "" {
		prefix = "TRAVEL"
	}
	c := &Client{
		HTTP:      hc,
		Jar:       jar,
		BaseURL:   strings.TrimRight(baseURL, "/"),
		UserAgent: DefaultUA,
		Cookie:    strings.TrimSpace(os.Getenv(prefix + "_COOKIE")),
		Pacer:     ratelimit.NewPacerFromEnv(prefix),
		EnvPrefix: prefix,
	}
	c.LoadPersistedCookies()
	return c
}

// LoadPersistedCookies merges env cookie with on-disk session (env wins).
func (c *Client) LoadPersistedCookies() {
	if c.Cookie != "" {
		cookies.SetJar(c.Jar, c.BaseURL, c.Cookie)
		return
	}
	d, err := session.Load(c.EnvPrefix)
	if err != nil || d.Cookie == "" {
		return
	}
	c.Cookie = d.Cookie
	cookies.SetJar(c.Jar, c.BaseURL, c.Cookie)
}

// SavePersistedCookies writes the current cookie header to disk.
func (c *Client) SavePersistedCookies() error {
	return session.Save(c.EnvPrefix, session.Data{Cookie: c.Cookie, BaseURL: c.BaseURL})
}

// ApplyCookieHeader replaces the client cookie and syncs the jar.
func (c *Client) ApplyCookieHeader(raw string) {
	if raw == "" {
		return
	}
	c.Cookie = cookies.MergeStrings(c.Cookie, raw)
	cookies.SetJar(c.Jar, c.BaseURL, c.Cookie)
}

// CookiesFilePath returns ~/.{slug}/cookies.json for this client.
func (c *Client) CookiesFilePath() string {
	return session.FilePath(c.EnvPrefix)
}

func (c *Client) Throttle() {
	if c.Pacer != nil {
		c.Pacer.Wait()
	}
}

func (c *Client) ApplyCookie(req *http.Request) {
	if c.Cookie != "" {
		req.Header.Set("cookie", c.Cookie)
	}
}

func (c *Client) SetAPIHeaders(req *http.Request) {
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "es-ES,es;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("user-agent", c.UserAgent)
	if c.BaseURL != "" {
		req.Header.Set("referer", c.BaseURL+"/")
	}
}

func (c *Client) SetDocumentHeaders(req *http.Request) {
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("accept-language", "es-ES,es;q=0.9,en;q=0.8")
	req.Header.Set("user-agent", c.UserAgent)
}

// GetJSON performs GET and decodes JSON into out.
func (c *Client) GetJSON(path string, out any) error {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return err
	}
	return c.DoJSON(req, out)
}

// GetRaw performs GET and returns the response body.
func (c *Client) GetRaw(url string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return body, resp.StatusCode, nil
}

// PostJSON performs POST with a JSON body and decodes the response.
func (c *Client) PostJSON(url string, payload, out any) error {
	c.Throttle()
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(b))
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	if c.BaseURL != "" {
		req.Header.Set("origin", c.BaseURL)
	}
	c.ApplyCookie(req)
	return c.DoJSON(req, out)
}

// DoJSON executes req and decodes JSON response.
func (c *Client) DoJSON(req *http.Request, out any) error {
	c.Throttle()
	c.SetAPIHeaders(req)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{Status: resp.StatusCode, Body: Truncate(string(body), 300)}
	}
	if out != nil && len(body) > 0 {
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("decode json: %w", err)
		}
	}
	return nil
}

// FetchHTML GETs a document URL.
func (c *Client) FetchHTML(url string) (string, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	c.SetDocumentHeaders(req)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	text := string(body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", &HTTPError{Status: resp.StatusCode, Body: Truncate(text, 300)}
	}
	return text, nil
}

// HTTPError is a non-2xx response.
type HTTPError struct {
	Status int
	Body   string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, e.Body)
}

// Truncate shortens s to at most n bytes.
func Truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

// Absolutize joins base and path.
func Absolutize(base, path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	if strings.HasPrefix(path, "//") {
		return "https:" + path
	}
	if strings.HasPrefix(path, "/") {
		return strings.TrimRight(base, "/") + path
	}
	return strings.TrimRight(base, "/") + "/" + path
}

// FirstNonEmpty returns the first non-empty string.
func FirstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
