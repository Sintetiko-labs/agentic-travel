// Package base provides a reusable HTTP client skeleton for travel CLIs.
package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/chrome"
	"github.com/fbelchi/travelkit/cookies"
	"github.com/fbelchi/travelkit/network"
	"github.com/fbelchi/travelkit/ratelimit"
	"github.com/fbelchi/travelkit/session"
	"github.com/fbelchi/travelkit/transport"
)

const (
	DefaultUA             = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	defaultRequestTimeout = 25 * time.Second
	requestTimeoutEnvKey  = "TRAVELKIT_REQUEST_TIMEOUT"
)

func RequestTimeout() time.Duration {
	if v := strings.TrimSpace(os.Getenv(requestTimeoutEnvKey)); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d > 0 {
			return d
		}
	}
	return defaultRequestTimeout
}


// Client is the shared HTTP transport used by scaffolded travel CLIs.
type Client struct {
	HTTP       *http.Client
	stdlibHTTP *http.Client
	Jar        cookies.Jar
	BaseURL    string
	UserAgent  string
	Cookie     string
	Pacer      *ratelimit.Pacer
	EnvPrefix  string
}

// New builds a client with Chrome-like TLS and optional env-based rate limiting.
func New(baseURL, envPrefix string) *Client {
	network.EnsureResidential()
	jar := cookies.NewJar()
	prefix := strings.ToUpper(strings.ReplaceAll(envPrefix, "-", "_"))
	if prefix == "" {
		prefix = "TRAVEL"
	}
	timeout := RequestTimeout()
	stdlib := &http.Client{Timeout: timeout, Jar: jar, Transport: network.DirectTransport()}
	hc := stdlib
	if os.Getenv(prefix+"_STD_HTTP") != "1" {
		if tr, err := transport.SharedRoundTripper(); err == nil {
			hc = &http.Client{Timeout: timeout, Jar: jar, Transport: tr}
		}
	}
	c := &Client{
		HTTP:       hc,
		stdlibHTTP: stdlib,
		Jar:        jar,
		BaseURL:    strings.TrimRight(baseURL, "/"),
		UserAgent:  DefaultUA,
		Cookie:     strings.TrimSpace(os.Getenv(prefix + "_COOKIE")),
		Pacer:      ratelimit.NewPacerFromEnv(prefix),
		EnvPrefix:  prefix,
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
	body, status, err := c.getRawWith(c.HTTP, url)
	if err == nil || c.stdlibHTTP == nil || c.stdlibHTTP == c.HTTP {
		return body, status, err
	}
	return c.getRawWith(c.stdlibHTTP, url)
}

func (c *Client) getRawWith(hc *http.Client, url string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	c.ApplyCookie(req)
	resp, err := hc.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	body, status := c.maybeChromeFallback(req, resp.StatusCode, body)
	return body, status, nil
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
	body, status := c.maybeChromeFallback(req, resp.StatusCode, body)
	if status < 200 || status >= 300 {
		return &HTTPError{Status: status, Body: Truncate(string(body), 300)}
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
	text, err := c.fetchHTMLWith(c.HTTP, url)
	if err == nil {
		return text, nil
	}
	if c.stdlibHTTP == nil || c.stdlibHTTP == c.HTTP {
		return "", err
	}
	return c.fetchHTMLWith(c.stdlibHTTP, url)
}

func (c *Client) fetchHTMLWith(hc *http.Client, url string) (string, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	c.SetDocumentHeaders(req)
	c.ApplyCookie(req)
	resp, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	text := string(body)
	textBody, status := c.maybeChromeFallback(req, resp.StatusCode, body)
	text = string(textBody)
	if status < 200 || status >= 300 {
		return "", &HTTPError{Status: status, Body: Truncate(text, 300)}
	}
	return text, nil
}

// ChromePort returns the CDP debugging port for this client.
func (c *Client) ChromePort() int {
	port := chrome.CDPPortFromEnv()
	if p := strings.TrimSpace(os.Getenv(c.EnvPrefix + "_CHROME_PORT")); p != "" {
		if n, err := parsePort(p); err == nil {
			port = n
		}
	}
	return port
}

func parsePort(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil || n <= 0 {
		return 0, fmt.Errorf("invalid port %q", s)
	}
	return n, nil
}

// ChromeFetchEnabled reports whether CDP fetch fallback can run.
func (c *Client) ChromeFetchEnabled() bool {
	if c.Cookie == "" {
		return false
	}
	return chrome.CDPAvailable(c.ChromePort())
}

// FetchViaChrome executes fetch() in headed Chrome with saved cookies.
func (c *Client) FetchViaChrome(url, method string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, 0, err
	}
	if len(body) > 0 {
		req.Body = io.NopCloser(bytes.NewReader(body))
		req.ContentLength = int64(len(body))
	}
	c.SetAPIHeaders(req)
	if len(body) > 0 {
		req.Header.Set("content-type", "application/json")
	}
	if c.BaseURL != "" {
		req.Header.Set("origin", c.BaseURL)
	}
	return c.FetchViaChromeReq(req)
}

// FetchViaChromeReq runs fetch() in Chrome using headers and body from req.
func (c *Client) FetchViaChromeReq(req *http.Request) ([]byte, int, error) {
	if !c.ChromeFetchEnabled() {
		return nil, 0, fmt.Errorf("chrome fetch unavailable — start Chrome with --remote-debugging-port=%d", c.ChromePort())
	}
	c.ApplyCookie(req)

	var bodyArg string
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, 0, err
		}
		req.Body = io.NopCloser(bytes.NewReader(b))
		bodyArg = string(b)
	}

	headerMap := map[string]string{}
	for k, vals := range req.Header {
		if len(vals) == 0 {
			continue
		}
		lower := strings.ToLower(k)
		if lower == "cookie" || lower == "host" || lower == "content-length" {
			continue
		}
		headerMap[k] = vals[0]
	}

	result, err := chrome.Fetch(context.Background(), chrome.FetchOpts{
		Port:    c.ChromePort(),
		URL:     req.URL.String(),
		Method:  req.Method,
		Body:    bodyArg,
		Headers: headerMap,
		Cookie:  c.Cookie,
	})
	if err != nil {
		return nil, 0, err
	}
	return []byte(result.Body), result.Status, nil
}


// maybeChromeFallback retries the request via Chrome CDP when utls returns 403.
func (c *Client) maybeChromeFallback(req *http.Request, status int, body []byte) ([]byte, int) {
	if status != http.StatusForbidden || !c.ChromeFetchEnabled() {
		return body, status
	}
	chromeBody, chromeStatus, err := c.FetchViaChromeReq(req)
	if err != nil {
		return body, status
	}
	return chromeBody, chromeStatus
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
