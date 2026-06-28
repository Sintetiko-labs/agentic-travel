// Package session persists browser cookies for travel CLIs.
package session

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/chrome"
	"github.com/fbelchi/travelkit/config"
)

const cookiesFile = "cookies.json"

// Data is persisted cookie material.
type Data struct {
	Cookie   string    `json:"cookie,omitempty"`
	Captured time.Time `json:"captured_at,omitempty"`
	BaseURL  string    `json:"base_url,omitempty"`
}

// FilePath returns the on-disk cookies path for envPrefix.
func FilePath(envPrefix string) string {
	return filepath.Join(config.Dir(envPrefix), cookiesFile)
}

// Load reads persisted cookies for envPrefix.
func Load(envPrefix string) (Data, error) {
	var d Data
	err := config.Load(envPrefix, cookiesFile, &d)
	return d, err
}

// Save writes cookies for envPrefix.
func Save(envPrefix string, d Data) error {
	if d.Captured.IsZero() {
		d.Captured = time.Now().UTC()
	}
	return config.Save(envPrefix, cookiesFile, d)
}

// ChromeOptions configures a session chrome capture.
type ChromeOptions struct {
	EnvPrefix   string
	BaseURL     string
	StartURL    string
	Port        int
	Wait        bool
	WaitTimeout time.Duration
	Replace     bool
	SyncOnly    bool
}

// ChromeResult is the outcome of session chrome.
type ChromeResult struct {
	Cookie  string
	Ready   bool
	HasAbck bool
	HasBmSz bool
}

// CaptureChrome connects to Chrome CDP and captures cookies for baseURL.
func CaptureChrome(opts ChromeOptions) (ChromeResult, error) {
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
	res, err := chrome.Capture(chrome.Options{
		EnvPrefix:   opts.EnvPrefix,
		BaseURL:     opts.BaseURL,
		StartURL:    opts.StartURL,
		Port:        opts.Port,
		WaitTimeout: opts.WaitTimeout,
		Replace:     opts.Replace,
	})
	if err != nil {
		return ChromeResult{}, err
	}
	lower := strings.ToLower(res.Cookie)
	hasAbck := strings.Contains(lower, "_abck=")
	hasBmSz := strings.Contains(lower, "bm_sz=")
	ready := hasAbck || hasBmSz || strings.Contains(lower, "cf_clearance=") ||
		strings.Contains(lower, "incap_ses") || strings.Contains(lower, "visid_incap")
	if !opts.Wait || ready || res.Cookie != "" {
		return ChromeResult{Cookie: res.Cookie, Ready: ready, HasAbck: hasAbck, HasBmSz: hasBmSz}, nil
	}
	return ChromeResult{}, fmt.Errorf("timeout waiting for session cookies on %s — %s",
		opts.BaseURL, akamai.NeedsSessionHint(opts.EnvPrefix))
}
