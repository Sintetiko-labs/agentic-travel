package chrome

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type FetchOpts struct {
	Port    int
	URL     string
	Method  string
	Body    string
	Headers map[string]string
	Cookie  string
}

type FetchResult struct {
	Status int
	Body   string
}

func CDPPortFromEnv() int {
	for _, k := range []string{"CHROME_PORT", "CHROME_DEBUG_PORT"} {
		if p := strings.TrimSpace(os.Getenv(k)); p != "" {
			if n, err := strconv.Atoi(p); err == nil && n > 0 {
				return n
			}
		}
	}
	return 9222
}

func CDPAvailable(port int) bool {
	if port <= 0 {
		port = CDPPortFromEnv()
	}
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/json/version", port))
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == 200
}

func Fetch(ctx context.Context, opts FetchOpts) (FetchResult, error) {
	if opts.Port <= 0 {
		opts.Port = CDPPortFromEnv()
	}
	if !CDPAvailable(opts.Port) {
		return FetchResult{}, fmt.Errorf("Chrome not listening on port %d", opts.Port)
	}
	method := strings.ToUpper(strings.TrimSpace(opts.Method))
	if method == "" {
		method = http.MethodGet
	}
	headerJSON, _ := json.Marshal(opts.Headers)
	script := fmt.Sprintf(`(async () => {
		const headers = %s;
		if (%q) headers['cookie'] = %q;
		const init = { method: %q, headers, credentials: 'include' };
		if (%q) init.body = %q;
		const resp = await fetch(%q, init);
		const text = await resp.text();
		return JSON.stringify({ status: resp.status, body: text });
	})()`, string(headerJSON), opts.Cookie != "", opts.Cookie, method, opts.Body != "", opts.Body, opts.URL)
	debugURL := fmt.Sprintf("http://127.0.0.1:%d", opts.Port)
	allocCtx, cancel := chromedp.NewRemoteAllocator(ctx, debugURL)
	defer cancel()
	cdpCtx, cancel2 := chromedp.NewContext(allocCtx)
	defer cancel2()
	runCtx, cancelRun := context.WithTimeout(cdpCtx, 60*time.Second)
	defer cancelRun()
	var raw string
	if err := chromedp.Run(runCtx, chromedp.Evaluate(script, &raw)); err != nil {
		return FetchResult{}, err
	}
	var parsed struct {
		Status int    `json:"status"`
		Body   string `json:"body"`
	}
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return FetchResult{}, fmt.Errorf("decode fetch result: %w", err)
	}
	return FetchResult{Status: parsed.Status, Body: parsed.Body}, nil
}
