package akamai

import "strings"

// IsDenied reports whether body looks like an Akamai/edge WAF block page.
func IsDenied(status int, body string) bool {
	if status == 403 || status == 401 {
		low := strings.ToLower(body)
		if strings.Contains(low, "access denied") ||
			strings.Contains(low, "edgesuite.net") ||
			strings.Contains(low, "messageblocked") ||
			strings.Contains(low, "akamai") && strings.Contains(low, "denied") {
			return true
		}
	}
	return false
}

// NeedsSessionHint returns guidance when Akamai blocks a request.
func NeedsSessionHint(envPrefix string) string {
	p := strings.ToUpper(envPrefix)
	return p + "_COOKIE required — run: " + strings.ToLower(envPrefix) + " session chrome"
}
