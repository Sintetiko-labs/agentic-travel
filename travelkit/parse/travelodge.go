package parse

import (
	"regexp"
	"strings"
)

var travelodgeSitemapLocRE = regexp.MustCompile(`<loc>(https://www\.travelodge\.co\.uk/[^<]+)</loc>`)

// HotelsFromTravelodgeSitemap extracts hotel rows from travelodge.co.uk sitemap-fusion.xml.
func HotelsFromTravelodgeSitemap(xml, query string) []HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range travelodgeSitemapLocRE.FindAllStringSubmatch(xml, -1) {
		url := m[1]
		if !strings.HasSuffix(url, "/index.html") {
			continue
		}
		parts := strings.Split(strings.Trim(strings.TrimPrefix(url, "https://www.travelodge.co.uk/"), "/"), "/")
		if len(parts) < 3 || parts[0] != "uk" || parts[len(parts)-1] != "index.html" {
			continue
		}
		region := parts[1]
		slug := parts[len(parts)-2]
		if slug == "index.html" || strings.Contains(slug, "hotels-with-") {
			continue
		}
		if !matchesTravelodgeQuery(region, slug, url, q) {
			continue
		}
		if seen[url] {
			continue
		}
		seen[url] = true
		id := strings.Join(parts[1:len(parts)-1], "/")
		name := titleFromSlug(strings.ReplaceAll(slug, "'", "'"))
		if !strings.HasPrefix(strings.ToLower(name), "travelodge") {
			name = "Travelodge " + name
		}
		out = append(out, HotelLD{
			ID:      id,
			Name:    name,
			Address: titleFromSlug(strings.ReplaceAll(region, "-", " ")),
			URL:     url,
		})
	}
	return out
}

func matchesTravelodgeQuery(region, slug, url, q string) bool {
	if q == "" {
		return true
	}
	low := strings.ToLower(region + " " + slug + " " + url)
	return strings.Contains(low, q) ||
		strings.Contains(low, strings.ReplaceAll(q, " ", "-"))
}
