package parse

import (
	"regexp"
	"strings"
)

var (
	hiltonHotelCardRE = regexp.MustCompile(`href="(https://www\.hilton\.com/en/hotels/[^"]+)"[^>]*>[\s\S]*?data-testid="listViewPropertyName"[^>]*>([^<]+)`)
	hiltonHotelLinkRE = regexp.MustCompile(`href="(https://www\.hilton\.com/en/hotels/([^"/]+))/?"`)
)

// HotelsFromHiltonLocations extracts hotel rows from Hilton destination listing pages.
func HotelsFromHiltonLocations(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range hiltonHotelCardRE.FindAllStringSubmatch(html, -1) {
		url, name := strings.TrimRight(m[1], "/")+"/", strings.TrimSpace(m[2])
		if seen[url] || name == "" {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID:   hiltonIDFromURL(url),
			Name: name,
			URL:  url,
		})
	}
	if len(out) > 0 {
		return out
	}
	for _, m := range hiltonHotelLinkRE.FindAllStringSubmatch(html, -1) {
		url, slug := strings.TrimRight(m[1], "/")+"/", m[2]
		if seen[url] {
			continue
		}
		seen[url] = true
		name := hiltonNameFromSlug(slug)
		out = append(out, HotelLD{
			ID:   slug,
			Name: name,
			URL:  url,
		})
	}
	return out
}

func hiltonIDFromURL(raw string) string {
	parts := strings.Split(strings.Trim(raw, "/"), "/")
	if len(parts) == 0 {
		return raw
	}
	return parts[len(parts)-1]
}

func hiltonNameFromSlug(slug string) string {
	parts := strings.SplitN(slug, "-", 2)
	if len(parts) == 2 {
		return titleFromSlug(parts[1])
	}
	return titleFromSlug(slug)
}
