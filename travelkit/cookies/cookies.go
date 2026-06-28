package cookies

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Jar is the cookie jar type used by travel HTTP clients.
type Jar = *cookiejar.Jar

// NewJar creates a cookie jar for travel sessions.
func NewJar() Jar {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	return jar
}

// MergeStrings merges multiple Cookie header strings, last value wins per name.
func MergeStrings(parts ...string) string {
	m := map[string]string{}
	order := []string{}
	for _, raw := range parts {
		for _, pair := range SplitHeader(raw) {
			name, val, ok := strings.Cut(pair, "=")
			if !ok || name == "" {
				continue
			}
			name = strings.TrimSpace(name)
			if _, seen := m[name]; !seen {
				order = append(order, name)
			}
			m[name] = strings.TrimSpace(val)
		}
	}
	out := make([]string, 0, len(order))
	for _, name := range order {
		out = append(out, name+"="+m[name])
	}
	return strings.Join(out, "; ")
}

// SplitHeader splits a Cookie header into name=value pairs.
func SplitHeader(raw string) []string {
	if raw == "" {
		return nil
	}
	var parts []string
	for _, chunk := range strings.Split(raw, ";") {
		if s := strings.TrimSpace(chunk); s != "" {
			parts = append(parts, s)
		}
	}
	return parts
}

// SetJar applies a raw Cookie header string to a jar for baseURL.
func SetJar(jar Jar, baseURL, raw string) {
	if jar == nil || raw == "" {
		return
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	var cookies []*http.Cookie
	for _, pair := range SplitHeader(raw) {
		name, val, ok := strings.Cut(pair, "=")
		if !ok || name == "" {
			continue
		}
		cookies = append(cookies, &http.Cookie{
			Name:  strings.TrimSpace(name),
			Value: SanitizeValue(strings.TrimSpace(val)),
		})
	}
	jar.SetCookies(u, cookies)
}

// JarString serializes jar cookies for baseURL as a Cookie header.
func JarString(jar Jar, baseURL string) string {
	if jar == nil {
		return ""
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	cookies := jar.Cookies(u)
	if len(cookies) == 0 {
		return ""
	}
	parts := make([]string, 0, len(cookies))
	for _, c := range cookies {
		parts = append(parts, c.Name+"="+c.Value)
	}
	return strings.Join(parts, "; ")
}

// ParseSetCookieHeader extracts the first name=value from a Set-Cookie header.
func ParseSetCookieHeader(v string) (name, value string, ok bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return "", "", false
	}
	first, _, _ := strings.Cut(v, ";")
	name, val, ok := strings.Cut(strings.TrimSpace(first), "=")
	if !ok || name == "" {
		return "", "", false
	}
	return strings.TrimSpace(name), strings.TrimSpace(val), true
}

// SanitizeValue strips quotes and control chars from cookie values.
func SanitizeValue(v string) string {
	v = strings.TrimSpace(v)
	if len(v) >= 2 && v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]
	}
	return v
}

// SanitizeHeader re-sanitizes each cookie pair in a Cookie header string.
func SanitizeHeader(raw string) string {
	parts := SplitHeader(raw)
	out := make([]string, 0, len(parts))
	for _, pair := range parts {
		name, val, ok := strings.Cut(pair, "=")
		if !ok || name == "" {
			continue
		}
		out = append(out, strings.TrimSpace(name)+"="+SanitizeValue(val))
	}
	return strings.Join(out, "; ")
}
