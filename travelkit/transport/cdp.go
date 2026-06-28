package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fbelchi/travelkit/cookies"
)

const DefaultCDPPort = 9222

func CDPAvailable(port int) bool {
	if port <= 0 {
		port = DefaultCDPPort
	}
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/json/version", port))
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == 200
}

func CDPPortFromEnv() int {
	if v := strings.TrimSpace(os.Getenv("TRAVEL_CHROME_PORT")); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			return p
		}
	}
	return DefaultCDPPort
}

type cdpFetchResult struct {
	Status  int               `json:"status"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

func FetchViaChrome(port int, req *http.Request, fallback http.RoundTripper) (*http.Response, error) {
	if port <= 0 {
		port = DefaultCDPPort
	}
	if !CDPAvailable(port) {
		if fallback != nil {
			return fallback.RoundTrip(req)
		}
		return nil, fmt.Errorf("chrome not listening on port %d", port)
	}
	debugURL := fmt.Sprintf("http://127.0.0.1:%d", port)
	allocCtx, cancel := chromedp.NewRemoteAllocator(req.Context(), debugURL)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	cookieHeader := req.Header.Get("Cookie")
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
	headersJSON, _ := json.Marshal(headerMap)

	var bodyArg string
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(b))
		bodyArg = string(b)
	}
	bodyJSON, _ := json.Marshal(bodyArg)

	var result cdpFetchResult
	actions := []chromedp.Action{network.Enable()}
	if cookieHeader != "" {
		u := req.URL
		host := u.Hostname()
		for _, pair := range cookies.SplitHeader(cookieHeader) {
			name, val, ok := strings.Cut(pair, "=")
			if !ok || strings.TrimSpace(name) == "" {
				continue
			}
			cookieName := strings.TrimSpace(name)
			cookieVal := cookies.SanitizeValue(strings.TrimSpace(val))
			actions = append(actions, chromedp.ActionFunc(func(ctx context.Context) error {
				return network.SetCookie(cookieName, cookieVal).
					WithDomain(host).WithPath("/").WithSecure(u.Scheme == "https").Do(ctx)
			}))
		}
	}
	script := fmt.Sprintf(`(async () => {
  const headers = %s;
  const opts = {method: %q, headers, credentials: 'include'};
  const body = %s;
  if (body) opts.body = body;
  const r = await fetch(%q, opts);
  const text = await r.text();
  const hdrs = {};
  r.headers.forEach((v, k) => { hdrs[k] = v; });
  return {status: r.status, body: text, headers: hdrs};
})()`, string(headersJSON), req.Method, string(bodyJSON), req.URL.String())
	actions = append(actions, chromedp.Evaluate(script, &result))

	runCtx, runCancel := context.WithTimeout(ctx, 45*time.Second)
	defer runCancel()
	if err := chromedp.Run(runCtx, actions...); err != nil {
		if fallback != nil {
			return fallback.RoundTrip(req)
		}
		return nil, fmt.Errorf("cdp fetch: %w", err)
	}

	resp := &http.Response{
		Status:     fmt.Sprintf("%d %s", result.Status, http.StatusText(result.Status)),
		StatusCode: result.Status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(result.Body)),
		Request:    req,
	}
	for k, v := range result.Headers {
		resp.Header.Set(k, v)
	}
	return resp, nil
}

var _ sync.Mutex
