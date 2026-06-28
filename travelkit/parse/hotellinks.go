package parse

import (
	"regexp"
	"strings"
)

var (
	cataloniaHotelLinkRE = regexp.MustCompile(`href="(https://www\.cataloniahotels\.com/es/hotel/[^"]+)"`)
	palladiumCardRE      = regexp.MustCompile(`data-hotel-name="([^"]+)"[^>]*data-provider-url="([^"]+)"`)
	lopesanHotelLinkRE   = regexp.MustCompile(`href="(https://www\.lopesan\.com/es/hoteles/[^"]+)"`)
	princessHotelH3RE    = regexp.MustCompile(`<h3[^>]*>\s*(?:<a[^>]*>)?\s*([^<]+?)\s*(?:</a>)?\s*</h3>`)
)

// HotelsFromCataloniaLinks extracts hotel rows from Catalonia homepage/listing HTML.
func HotelsFromCataloniaLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range cataloniaHotelLinkRE.FindAllStringSubmatch(html, -1) {
		url := m[1]
		if seen[url] {
			continue
		}
		seen[url] = true
		slug := strings.TrimPrefix(url, baseURL+"/es/hotel/")
		name := titleFromSlug(slug)
		out = append(out, HotelLD{
			ID:   slug,
			Name: name,
			URL:  url,
		})
	}
	return out
}

// HotelsFromPalladiumCards extracts hotel rows from Palladium listing HTML.
func HotelsFromPalladiumCards(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range palladiumCardRE.FindAllStringSubmatch(html, -1) {
		slug, url := m[1], m[2]
		if seen[url] {
			continue
		}
		seen[url] = true
		name := titleFromSlug(strings.ReplaceAll(slug, "-", " "))
		out = append(out, HotelLD{
			ID:   slug,
			Name: name,
			URL:  absolutize(baseURL, url),
		})
	}
	return out
}

// HotelsFromLopesanLinks extracts hotel detail links from Lopesan listing HTML.
func HotelsFromLopesanLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range lopesanHotelLinkRE.FindAllStringSubmatch(html, -1) {
		url := m[1]
		if strings.Contains(url, "/ofertas") || strings.Contains(url, "/destinos") {
			continue
		}
		parts := strings.Split(strings.Trim(url, "/"), "/")
		if len(parts) < 6 {
			continue
		}
		slug := parts[len(parts)-1]
		if seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID:   slug,
			Name: titleFromSlug(slug),
			URL:  url,
		})
	}
	return out
}

// HotelsFromPrincessHeadings extracts hotel names from Princess destination pages.
func HotelsFromPrincessHeadings(html, baseURL, pageURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range princessHotelH3RE.FindAllStringSubmatch(html, -1) {
		name := strings.TrimSpace(m[1])
		low := strings.ToLower(name)
		if name == "" || strings.Contains(low, "nuestros hoteles") {
			continue
		}
		if !strings.Contains(low, "princess") && !strings.Contains(low, "hotel") {
			continue
		}
		slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		if seen[slug] {
			continue
		}
		seen[slug] = true
		out = append(out, HotelLD{
			ID:   slug,
			Name: name,
			URL:  pageURL,
		})
	}
	return out
}

func titleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		if len(p) <= 3 && strings.ToUpper(p) == p {
			parts[i] = p
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
	}
	return strings.Join(parts, " ")
}
