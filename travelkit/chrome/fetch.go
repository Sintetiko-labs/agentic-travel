package chrome

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

// FetchOpts configures an in-page fetch() executed over CDP.
type FetchOpts struct {
	Port    int
	URL     string
	Method  string
	Body    string
	Headers map[string]string
	Cookie  string
}

// FetchResult is the HTTP status and body from Chrome fetch().
type FetchResult struct {
	Status int
	Body   string
}

// CDPPortFromEnv returns the default Chrome debugging port.
func CDPPortFromEnv() int {
	for _, key := range []string{"CHROME_DEBUG_PORT", "CHROME_PORT", "CHROMEDP_PORT"} {
		if p := os.Getenv(key); p != "" {
			if n, err := strconv.Atoi(p); err == nil && n > 0 {
				return n
			}
		}
	}
	return 9222
}

// CDPAvailable reports whether Chrome exposes a CDP HTTP endpoint on port.
func CDPAvailable(port int) bool {
	if port <= 0 {
		port = CDPPortFromEnv()
	}
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/json/version", port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// Fetch runs fetch() inside a Chrome tab attached to the debugging port.
func Fetch(ctx context.Context, opts FetchOpts) (FetchResult, error) {
	if opts.Port <= 0 {
		opts.Port = CDPPortFromEnv()
	}
	if opts.Method == "" {
		opts.Method = http.MethodGet
	}
	debugURL := fmt.Sprintf("http://127.0.0.1:%d", opts.Port)
	allocCtx, cancel := chromedp.NewRemoteAllocator(ctx, debugURL)
	defer cancel()
	tabCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	if _, ok := ctx.Deadline(); !ok {
		var cancelDeadline context.CancelFunc
		tabCtx, cancelDeadline = context.WithTimeout(tabCtx, 90*time.Second)
		defer cancelDeadline()
	}

	headersJSON, err := json.Marshal(opts.Headers)
	if err != nil {
		return FetchResult{}, err
	}
	script := fmt.Sprintf(`(async () => {
  const headers = %s;
  if (%q) headers['cookie'] = %q;
  const init = { method: %q, headers, credentials: 'include' };
  if (%q) { init.body = %q; }
  const resp = await fetch(%q, init);
  const body = await resp.text();
  return JSON.stringify({ status: resp.status, body });
})()`, string(headersJSON), opts.Cookie != "", opts.Cookie, opts.Method, opts.Body != "", opts.Body, opts.URL)

	var raw string
	if err := chromedp.Run(tabCtx, chromedp.Evaluate(script, &raw, chromedp.EvalAsValue)); err != nil {
		return FetchResult{}, fmt.Errorf("chrome fetch %s: %w", opts.URL, err)
	}
	var parsed struct {
		Status int    `json:"status"`
		Body   string `json:"body"`
	}
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return FetchResult{}, fmt.Errorf("chrome fetch decode: %w", err)
	}
	return FetchResult{Status: parsed.Status, Body: parsed.Body}, nil
}
