package akamai

import "strings"

// CookieReport summarizes WAF-related cookies in a header string.
type CookieReport struct {
	HasAbck      bool `json:"has_abck"`
	HasBmSz      bool `json:"has_bm_sz"`
	HasCF        bool `json:"has_cf_clearance"`
	HasIncapsula bool `json:"has_incapsula"`
}

// AnalyzeCookies inspects a raw Cookie header for common WAF tokens.
func AnalyzeCookies(cookie string) CookieReport {
	lower := strings.ToLower(cookie)
	return CookieReport{
		HasAbck:      strings.Contains(lower, "_abck="),
		HasBmSz:      strings.Contains(lower, "bm_sz="),
		HasCF:        strings.Contains(lower, "cf_clearance="),
		HasIncapsula: strings.Contains(lower, "incap_ses") || strings.Contains(lower, "visid_incap"),
	}
}

// HasRequiredAkamaiCookies reports both Akamai bot-manager cookies are present.
func HasRequiredAkamaiCookies(cookie string) bool {
	r := AnalyzeCookies(cookie)
	return r.HasAbck && r.HasBmSz
}

// SessionReady reports whether cookie material is enough for WAF-protected APIs.
func SessionReady(cookie string) bool {
	if cookie == "" {
		return false
	}
	if HasRequiredAkamaiCookies(cookie) {
		return true
	}
	r := AnalyzeCookies(cookie)
	if r.HasCF {
		return true
	}
	return r.HasIncapsula && strings.Contains(strings.ToLower(cookie), "visid_incap")
}
