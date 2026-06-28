package parse

import (
	"regexp"
	"strings"
)

var vincciHotelLinkRE = regexp.MustCompile(`href="(/es/hoteles/[^"#?]+/?)"`)

// HotelsFromVincciLinks extracts hotel rows from Vincci homepage/listing links.
func HotelsFromVincciLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range vincciHotelLinkRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) < 4 {
			continue
		}
		city, slug := parts[2], parts[3]
		if city == "" || slug == "" {
			continue
		}
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID:      slug,
			Name:    titleFromSlug(slug),
			Address: city,
			URL:     url,
		})
	}
	return out
}
