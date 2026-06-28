// Package chrome captures Akamai session cookies from a Chrome CDP debugging port.
package chrome

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/cookies"
)

// Options configures a Chrome cookie capture.
type Options struct {
	EnvPrefix   string
	BaseURL     string
	StartURL    string
	ExtraURLs   []string
	Port        int
	Wait        bool
	WaitTimeout time.Duration
	Replace     bool
	SyncOnly    bool
}

// Result is a captured cookie header string.
type Result struct {
	Cookie  string
	Ready   bool
	HasAbck bool
	HasBmSz bool
}

// Capture opens or attaches to Chrome and returns cookies for baseURL.
func Capture(opts Options) (Result, error) {
	if opts.Port == 0 {
		opts.Port = 9222
	}
	if opts.WaitTimeout == 0 {
		if opts.Wait {
			opts.WaitTimeout = 3 * time.Minute
		} else {
			opts.WaitTimeout = 30 * time.Second
		}
	}
	if opts.StartURL == "" {
		opts.StartURL = opts.BaseURL
	}
	debugURL := fmt.Sprintf("http://127.0.0.1:%d", opts.Port)
	if err := ensureDebugging(debugURL, opts); err != nil {
		return Result{}, err
	}
	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), debugURL)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	if err := chromedp.Run(ctx, chromedp.Navigate(opts.StartURL)); err != nil {
		return Result{}, fmt.Errorf("navigate %s: %w", opts.StartURL, err)
	}
	deadline := time.Now().Add(opts.WaitTimeout)
	for {
		cookieHeader, err := readCookieHeader(ctx, cookieURLs(opts))
		if err != nil {
			return Result{}, err
		}
		report := akamai.AnalyzeCookies(cookieHeader)
		ready := captureReady(opts, cookieHeader)
		if ready {
			return Result{
				Cookie: cookieHeader, Ready: true,
				HasAbck: report.HasAbck, HasBmSz: report.HasBmSz,
			}, nil
		}
		if time.Now().After(deadline) {
			if opts.Wait && cookieHeader != "" {
				return Result{
					Cookie: cookieHeader, Ready: false,
					HasAbck: report.HasAbck, HasBmSz: report.HasBmSz,
				}, fmt.Errorf("timeout (%s) waiting for WAF cookies (_abck+bm_sz, cf_clearance, or Incapsula) on %s — browse the site in headed Chrome",
					opts.WaitTimeout, opts.BaseURL)
			}
			return Result{}, fmt.Errorf("timeout (%s) waiting for cookies on %s", opts.WaitTimeout, opts.BaseURL)
		}
		time.Sleep(2 * time.Second)
	}
}

func captureReady(opts Options, cookie string) bool {
	if cookie == "" {
		return false
	}
	if opts.SyncOnly || !opts.Wait {
		return true
	}
	return akamai.SessionReady(cookie)
}

func cookieURLs(opts Options) []string {
	seen := make(map[string]struct{})
	var urls []string
	add := func(u string) {
		u = strings.TrimRight(strings.TrimSpace(u), "/")
		if u == "" {
			return
		}
		if _, ok := seen[u]; ok {
			return
		}
		seen[u] = struct{}{}
		urls = append(urls, u, u+"/")
	}
	add(opts.BaseURL)
	add(opts.StartURL)
	for _, u := range opts.ExtraURLs {
		add(u)
	}
	return urls
}

func readCookieHeader(ctx context.Context, urls []string) (string, error) {
	var cks []*network.Cookie
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		cks, err = network.GetCookies().WithURLs(urls).Do(ctx)
		return err
	})); err != nil {
		return "", err
	}
	parts := make([]string, 0, len(cks))
	seen := make(map[string]struct{})
	for _, ck := range cks {
		if ck.Name == "" {
			continue
		}
		if _, ok := seen[ck.Name]; ok {
			continue
		}
		seen[ck.Name] = struct{}{}
		parts = append(parts, ck.Name+"="+ck.Value)
	}
	return strings.Join(parts, "; "), nil
}

func profileDir(envPrefix string) string {
	slug := strings.ToLower(strings.ReplaceAll(envPrefix, "_", "-"))
	return filepath.Join(os.Getenv("HOME"), "."+slug, "chrome-profile")
}

func ensureDebugging(debugURL string, opts Options) error {
	if resp, err := http.Get(debugURL + "/json/version"); err == nil {
		resp.Body.Close()
		return nil
	}
	if !opts.Replace {
		return fmt.Errorf("Chrome not listening on %s — launch headed Chrome with: chrome --remote-debugging-port=%d --user-data-dir=%q %s",
			debugURL, opts.Port, profileDir(opts.EnvPrefix), opts.StartURL)
	}
	_ = exec.Command("pkill", "-f", profileDir(opts.EnvPrefix)).Run()
	args := []string{
		fmt.Sprintf("--remote-debugging-port=%d", opts.Port),
		"--user-data-dir=" + profileDir(opts.EnvPrefix),
		"--no-first-run",
		"--no-default-browser-check",
		opts.StartURL,
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", append([]string{"-na", "Google Chrome", "--args"}, args...)...)
	default:
		cmd = exec.Command("google-chrome", args...)
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	deadline := time.Now().Add(30 * time.Second)
	for {
		if resp, err := http.Get(debugURL + "/json/version"); err == nil {
			resp.Body.Close()
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("Chrome did not start on port %d", opts.Port)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// ApplyToClient sets captured cookies on a travelkit HTTP client jar.
func ApplyToClient(jar cookies.Jar, baseURL, raw string) {
	cookies.SetJar(jar, baseURL, raw)
}
