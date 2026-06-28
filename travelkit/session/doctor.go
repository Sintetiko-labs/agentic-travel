package session

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/network"
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
	Slug              string
	EnvPrefix         string
	BaseURL           string
	Cookie            string
	ProbeURL          string
	ProbeMethod       string
	ProbeBody         string
	ProbeContentType  string
	ProbeOrigin       string
	ProbeReferer      string
	ProbeHeaders      map[string]string
	SessionOptional   bool // probe success is enough without WAF cookies
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
	cookieOK := cookie != "" && akamai.SessionReady(cookie)
	siteCookieOK := cookie != "" && !akamai.NeedsAkamaiWAF(slug) &&
		(akamai.SiteSessionReady(slug, cookie) || akamai.HasSessionMaterial(cookie))
	if cookie == "" {
		if !opts.SessionOptional {
			res.Status = DoctorMissing
			res.Message = "no saved session — APIs behind Akamai/Incapsula need headed Chrome cookies"
			res.NextStep = slug + " session chrome --wait --timeout 3m"
			if opts.ProbeURL == "" {
				return res
			}
		}
	} else if !cookieOK && akamai.NeedsAkamaiWAF(slug) && !opts.SessionOptional {
		res.Status = DoctorIncomplete
		res.Message = "WAF cookies incomplete — need _abck+bm_sz (Akamai), cf_clearance (Cloudflare), or Incapsula pair"
		res.NextStep = slug + " session chrome --wait --timeout 3m"
		if opts.ProbeURL == "" {
			return res
		}
	}
	if opts.ProbeURL == "" {
		if cookieOK {
			res.Status = DoctorOK
			res.Message = "WAF cookies present — run search to verify API access"
			return res
		}
		res.Status = DoctorIncomplete
		return res
	}
	method := strings.ToUpper(strings.TrimSpace(opts.ProbeMethod))
	if method == "" {
		method = http.MethodGet
	}
	probeCookie := ""
	if cookieOK || siteCookieOK || opts.SessionOptional {
		probeCookie = cookie
	}
	status, probeBody, err := probeHTTP(probeRequest{
		method: method, url: opts.ProbeURL, cookie: probeCookie,
		body: opts.ProbeBody, contentType: opts.ProbeContentType,
		origin: opts.ProbeOrigin, referer: opts.ProbeReferer, baseURL: opts.BaseURL,
		extra: opts.ProbeHeaders,
	})
	res.ProbeHTTPStatus = status
	if err != nil {
		res.Status = DoctorAPIError
		res.Message = fmt.Sprintf("probe failed: %v", err)
		if !opts.SessionOptional {
			res.NextStep = slug + " session chrome --wait --timeout 3m"
		}
		return res
	}
	switch {
	case status >= 200 && status < 300 && akamai.IsWAFBlocked(status, probeBody):
		res.Status = DoctorBlocked
		if cookieOK {
			res.Message = fmt.Sprintf("probe returned WAF challenge (HTTP %d) — cookies may be stale", status)
		} else {
			res.Message = fmt.Sprintf("probe returned WAF challenge (HTTP %d) — need headed Chrome session", status)
		}
		res.NextStep = slug + " session chrome --wait --timeout 3m"
	case status >= 200 && status < 300:
		res.Status = DoctorOK
		if cookieOK || siteCookieOK {
			res.Message = fmt.Sprintf("session OK — cookies present, probe HTTP %d", status)
		} else if opts.SessionOptional {
			res.Message = fmt.Sprintf("API OK (HTTP %d) — session optional for this brand", status)
		} else {
			res.Message = fmt.Sprintf("probe HTTP %d but WAF cookies missing", status)
			res.Status = DoctorMissing
			res.NextStep = slug + " session chrome --wait --timeout 3m"
		}
	case status == 403 || status == 401:
		res.Status = DoctorBlocked
		if cookieOK {
			res.Message = fmt.Sprintf("API probe blocked (HTTP %d) — cookies may be stale", status)
		} else {
			res.Message = fmt.Sprintf("API probe blocked (HTTP %d) — need headed Chrome session", status)
		}
		res.NextStep = slug + " session chrome --wait --timeout 3m"
	case status == 404 || status == 405:
		if !cookieOK || akamai.IsAppNotFoundWithoutSession(status, probeBody) {
			res.Status = DoctorBlocked
			if cookieOK {
				res.Message = fmt.Sprintf("API probe blocked (HTTP %d) — cookies may be stale", status)
			} else {
				res.Message = fmt.Sprintf("API probe blocked (HTTP %d) — need headed Chrome session", status)
			}
			res.NextStep = slug + " session chrome --wait --timeout 3m"
		} else {
			res.Status = DoctorAPIError
			res.Message = fmt.Sprintf("probe endpoint not found (HTTP %d) — check API path in client", status)
		}
	default:
		res.Status = DoctorAPIError
		res.Message = fmt.Sprintf("unexpected probe status HTTP %d", status)
		if status >= 500 {
			res.NextStep = "retry later or " + slug + " session chrome --wait --timeout 3m"
		}
	}
	return res
}

type probeRequest struct {
	method, url, cookie, body, contentType, origin, referer, baseURL string
	extra map[string]string
}

func probeHTTP(p probeRequest) (int, string, error) {
	var body io.Reader
	if p.body != "" {
		body = strings.NewReader(p.body)
	}
	req, err := http.NewRequest(p.method, p.url, body)
	if err != nil {
		return 0, "", err
	}
	if p.cookie != "" {
		req.Header.Set("cookie", p.cookie)
	}
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	if p.contentType != "" {
		req.Header.Set("content-type", p.contentType)
	}
	origin := p.origin
	if origin == "" {
		origin = strings.TrimRight(p.baseURL, "/")
	}
	if origin != "" {
		req.Header.Set("origin", origin)
	}
	referer := p.referer
	if referer == "" && p.baseURL != "" {
		referer = strings.TrimRight(p.baseURL, "/") + "/"
	}
	if referer != "" {
		req.Header.Set("referer", referer)
	}
	for k, v := range p.extra {
		if k != "" && v != "" {
			req.Header.Set(k, v)
		}
	}
	resp, err := network.DirectClient(30 * time.Second).Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	return resp.StatusCode, string(raw), nil
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
