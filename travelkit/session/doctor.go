package session

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
)

// DoctorStatus classifies persisted session health.
type DoctorStatus string

const (
	DoctorOK          DoctorStatus = "ok"
	DoctorMissing     DoctorStatus = "missing_session"
	DoctorIncomplete  DoctorStatus = "incomplete_cookies"
	DoctorBlocked     DoctorStatus = "blocked"
	DoctorAPIError    DoctorStatus = "api_error"
)

// DoctorOptions configures a session health probe.
type DoctorOptions struct {
	Slug        string
	EnvPrefix   string
	BaseURL     string
	Cookie      string
	ProbeURL    string
	ProbeMethod string
}

// DoctorResult is the structured outcome of session doctor.
type DoctorResult struct {
	Status            DoctorStatus       `json:"status"`
	SessionFile       string             `json:"session_file"`
	SessionFileExists bool               `json:"session_file_exists"`
	SessionAge        string             `json:"session_age,omitempty"`
	SessionAgeSeconds int64              `json:"session_age_seconds,omitempty"`
	Cookies           akamai.CookieReport `json:"cookies"`
	ProbeHTTPStatus   int                `json:"probe_http_status,omitempty"`
	Message           string             `json:"message"`
	NextStep          string             `json:"next_step,omitempty"`
}

// Doctor inspects persisted cookies and optionally probes a brand API.
func Doctor(opts DoctorOptions) DoctorResult {
	path := FilePath(opts.EnvPrefix)
	exists, age := sessionFileInfo(path)
	cookie := strings.TrimSpace(opts.Cookie)
	if cookie == "" {
		if d, err := Load(opts.EnvPrefix); err == nil {
			cookie = d.Cookie
		}
	}
	report := akamai.AnalyzeCookies(cookie)
	res := DoctorResult{
		SessionFile:       path,
		SessionFileExists: exists,
		Cookies:           report,
	}
	if exists {
		res.SessionAge = formatAge(age)
		res.SessionAgeSeconds = int64(age.Seconds())
	}
	slug := strings.ToLower(opts.Slug)
	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(opts.EnvPrefix, "_", "-"))
	}
	if cookie == "" {
		res.Status = DoctorMissing
		res.Message = "no saved session — APIs behind Akamai/Incapsula need headed Chrome cookies"
		res.NextStep = slug + " session chrome --wait --timeout 3m"
		return res
	}
	if !akamai.SessionReady(cookie) {
		res.Status = DoctorIncomplete
		res.Message = "WAF cookies incomplete — need _abck+bm_sz (Akamai), cf_clearance (Cloudflare), or Incapsula pair"
		res.NextStep = slug + " session chrome --wait --timeout 3m"
		return res
	}
	if opts.ProbeURL == "" {
		res.Status = DoctorOK
		res.Message = "WAF cookies present — run search to verify API access"
		return res
	}
	method := strings.ToUpper(strings.TrimSpace(opts.ProbeMethod))
	if method == "" {
		method = http.MethodGet
	}
	status, err := probeHTTP(method, opts.ProbeURL, cookie)
	res.ProbeHTTPStatus = status
	if err != nil {
		res.Status = DoctorAPIError
		res.Message = fmt.Sprintf("probe failed: %v", err)
		res.NextStep = slug + " session chrome --wait --timeout 3m"
		return res
	}
	if status == 403 || status == 401 {
		if akamai.IsDenied(status, "") {
			res.Status = DoctorBlocked
			res.Message = fmt.Sprintf("API probe blocked (HTTP %d) — cookies may be stale", status)
			res.NextStep = slug + " session chrome --wait --timeout 3m"
			return res
		}
	}
	if status >= 200 && status < 500 {
		res.Status = DoctorOK
		res.Message = fmt.Sprintf("session OK — probe HTTP %d", status)
		return res
	}
	res.Status = DoctorAPIError
	res.Message = fmt.Sprintf("unexpected probe status HTTP %d", status)
	return res
}

func probeHTTP(method, url, cookie string) (int, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("cookie", cookie)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
	return resp.StatusCode, nil
}

func sessionFileInfo(path string) (bool, time.Duration) {
	info, err := os.Stat(path)
	if err != nil {
		return false, 0
	}
	return true, time.Since(info.ModTime())
}

func formatAge(d time.Duration) string {
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		m := int(d.Minutes())
		if m == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", m)
	default:
		h := int(d.Hours())
		if h == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", h)
	}
}
