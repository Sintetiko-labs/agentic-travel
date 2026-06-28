package akamai

import (
	"strings"
	"unicode"
)

// IsIncapsulaChallenge reports whether body is an Imperva Incapsula JS challenge page.
func IsIncapsulaChallenge(body string) bool {
	low := strings.ToLower(body)
	return strings.Contains(low, "_incapsula_resource") ||
		(strings.Contains(low, "incapsula") && strings.Contains(low, "<script"))
}

// LooksLikeJSON reports whether body appears to be JSON (object or array).
func LooksLikeJSON(body string) bool {
	s := strings.TrimSpace(body)
	if s == "" {
		return false
	}
	return s[0] == '{' || s[0] == '['
}

// IsWAFBlocked reports Akamai denial, Incapsula challenge, or HTML challenge on 2xx.
func IsWAFBlocked(status int, body string) bool {
	if IsDenied(status, body) {
		return true
	}
	if status >= 200 && status < 300 && IsIncapsulaChallenge(body) {
		return true
	}
	if status >= 200 && status < 300 && !LooksLikeJSON(body) {
		low := strings.ToLower(body)
		if strings.Contains(low, "<html") && (strings.Contains(low, "access denied") || strings.Contains(low, "_incapsula")) {
			return true
		}
	}
	return false
}

// ProbeOK reports whether a probe response indicates API access (not WAF/challenge).
func ProbeOK(status int, body string) bool {
	if status < 200 || status >= 300 {
		return false
	}
	if IsWAFBlocked(status, body) {
		return false
	}
	trim := strings.TrimSpace(body)
	if trim == "" {
		return true
	}
	if LooksLikeJSON(trim) {
		return true
	}
	// Some dapi endpoints return plain text like "404 page not found".
	for _, r := range trim {
		if !unicode.IsPrint(r) && r != '\n' && r != '\r' && r != '\t' {
			return false
		}
	}
	return !strings.Contains(strings.ToLower(trim), "<html")
}

// IsAppNotFoundWithoutSession reports 404 responses that usually mean the edge was
// bypassed without a valid session cookie (Next.js shell, empty body, Go 404 text).
func IsAppNotFoundWithoutSession(status int, body string) bool {
	if status != 404 {
		return false
	}
	trimmed := strings.TrimSpace(body)
	if trimmed == "" || trimmed == "404 page not found" {
		return true
	}
	return strings.Contains(body, "__next_error__")
}
