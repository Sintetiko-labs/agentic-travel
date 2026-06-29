package parse

import (
	"regexp"
	"strings"
)

var (
	accorHotelLinkRE = regexp.MustCompile(`href="((?:https?://[^"]*all\.accor\.com)?/[^"]*?/hotels/[^"?#]+)"`)
	accorHotelJSONRE = regexp.MustCompile(`"hotelCode"\s*:\s*"([A-Z0-9]{3,8})"[\s\S]{0,600}?"name"\s*:\s*"([^"\\]+)"`)
)

// HotelsFromAccorSearch extracts hotel rows from Accor destination or SSR search HTML.
func HotelsFromAccorSearch(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range accorHotelJSONRE.FindAllStringSubmatch(html, -1) {
		id, name := m[1], strings.TrimSpace(m[2])
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  absolutize(baseURL, "/hotels/"+strings.ToLower(id)+".html"),
		})
	}
	for _, m := range accorHotelLinkRE.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		low := strings.ToLower(u)
		if strings.Contains(low, "/ssr/") || strings.Contains(low, "/login") {
			continue
		}
		id := pathID(u)
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   id,
			Name: titleFromSlug(id),
			URL:  u,
		})
	}
	return out
}
