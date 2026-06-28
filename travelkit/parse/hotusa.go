package parse

import (
	"regexp"
	"strings"
)

var hotusaHotelLinkRE = regexp.MustCompile(`href="([^"]*/(?:hotel|hoteles)/[^"#?]+)"`)

// HotelsFromHotusaLinks extracts hotel detail links from Hotusa listing HTML.
func HotelsFromHotusaLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range hotusaHotelLinkRE.FindAllStringSubmatch(html, -1) {
		url := absolutize(baseURL, m[1])
		if seen[url] {
			continue
		}
		if strings.HasSuffix(url, "/es/hoteles") || strings.HasSuffix(url, "/es/hoteles/") {
			continue
		}
		seen[url] = true
		slug := slugFromHotelURL(url)
		out = append(out, HotelLD{
			ID:   slug,
			Name: titleFromSlug(slug),
			URL:  url,
		})
	}
	return out
}
