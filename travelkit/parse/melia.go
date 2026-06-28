package parse

import (
	"regexp"
	"strings"
)

var meliaHotelLinkRE = regexp.MustCompile(`href="(/es/hoteles/[^"#?]+)"[^>]*>([^<]+)</a>`)

// HotelsFromMeliaDirectory extracts hotel rows from the Meliá hotel directory page.
func HotelsFromMeliaDirectory(html, baseURL, query string) []HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range meliaHotelLinkRE.FindAllStringSubmatch(html, -1) {
		if len(m) < 3 {
			continue
		}
		path := strings.TrimSpace(m[1])
		name := strings.TrimSpace(m[2])
		if name == "" || seen[name] {
			continue
		}
		if q != "" && !strings.Contains(strings.ToLower(name), q) && !strings.Contains(strings.ToLower(path), q) {
			continue
		}
		seen[name] = true
		out = append(out, HotelLD{
			Name: name,
			URL:  absolutize(baseURL, path),
			ID:   slugFromPath(path),
		})
	}
	if len(out) > 0 {
		return out
	}
	rows := HotelsFromJSONLD(html, baseURL)
	if q == "" {
		return rows
	}
	filtered := make([]HotelLD, 0, len(rows))
	for _, h := range rows {
		if strings.Contains(strings.ToLower(h.Name), q) || strings.Contains(strings.ToLower(h.Address), q) {
			filtered = append(filtered, h)
		}
	}
	return filtered
}

func slugFromPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
