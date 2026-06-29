package hotel

import (
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/parse"
)

func normCity(query string) string {
	return strings.ToLower(strings.TrimSpace(query))
}

func slugify(query string) string {
	s := strings.ToLower(strings.TrimSpace(query))
	s = strings.NewReplacer("'", "", ".", "", ",", "", "&", "and").Replace(s)
	return strings.ReplaceAll(s, " ", "-")
}

func WyndhamDestinationSlug(query string) string {
	switch normCity(query) {
	case "london":
		return "london-england-united-kingdom"
	case "madrid":
		return "madrid-community-of-madrid-spain"
	case "paris":
		return "paris-ile-de-france-france"
	case "berlin":
		return "berlin-germany"
	default:
		return slugify(query)
	}
}

func RadissonDestinationSlug(query string) string {
	switch normCity(query) {
	case "london":
		return "london-uk"
	case "madrid":
		return "madrid-spain"
	case "paris":
		return "paris-france"
	case "berlin":
		return "berlin-germany"
	default:
		return slugify(query) + "-hotels"
	}
}

func HyattSearchSlug(query string) string {
	switch normCity(query) {
	case "london":
		return "london-uk"
	case "madrid":
		return "madrid-spain"
	case "paris":
		return "paris-france"
	case "berlin":
		return "berlin-germany"
	default:
		return slugify(query)
	}
}

func IHGQDest(query string) string {
	return strings.TrimSpace(query)
}

func IHGSearchPath(query string) string {
	locale := "gb"
	switch normCity(query) {
	case "madrid", "barcelona", "valencia", "seville", "sevilla", "malaga", "bilbao":
		locale = "es"
	case "paris", "lyon", "marseille":
		locale = "fr"
	case "berlin", "munich", "münchen", "frankfurt":
		locale = "de"
	}
	return "/hotels/" + locale + "/en/find-hotels/hotel/list?qDest=" + url.QueryEscape(IHGQDest(query))
}

func AccorDestinationPath(query string) string {
	switch normCity(query) {
	case "london":
		return "/a/en/destination/city/hotels-london-v2352.html"
	case "madrid":
		return "/ssr/app/accor/rates/index.en.shtml?destination=Madrid"
	case "paris":
		return "/a/en/destination/city/hotels-paris-v2240.html"
	case "berlin":
		return "/de/germany/berlin.hotels.html"
	default:
		return "/ssr/app/accor/rates/index.en.shtml?destination=" + url.QueryEscape(strings.TrimSpace(query))
	}
}

func AccorSearchFallbackPath(query string) string {
	q := url.QueryEscape(strings.TrimSpace(query))
	if q == "" {
		return ""
	}
	return "/ssr/app/accor/rates/index.en.shtml?destination=" + q
}

func BestWesternPlace(query string) string {
	switch normCity(query) {
	case "london":
		return "London, United Kingdom"
	case "madrid":
		return "Madrid, Spain"
	case "paris":
		return "Paris, France"
	case "berlin":
		return "Berlin, Germany"
	default:
		return strings.TrimSpace(query)
	}
}

func NumaCitySlug(query string) string { return slugify(query) }
func LimehomeDestinationSlug(query string) string { return slugify(query) }

func BBHotelsCityPath(query string) string {
	switch normCity(query) {
	case "london":
		return "/en/gb/london"
	case "paris":
		return "/en/fr/paris"
	case "berlin":
		return "/en/de/berlin"
	default:
		return ""
	}
}

func MatchHoxtonSlug(query string, h parse.HotelLD) bool {
	q := normCity(query)
	if q == "" {
		return true
	}
	blob := strings.ToLower(strings.Join([]string{h.Name, h.URL, h.ID, h.Address}, " "))
	switch q {
	case "london":
		return strings.Contains(blob, "london") || strings.Contains(blob, "shoreditch") || strings.Contains(blob, "southwark") || strings.Contains(blob, "holborn")
	case "paris":
		return strings.Contains(blob, "paris")
	case "berlin":
		return strings.Contains(blob, "berlin")
	default:
		return strings.Contains(blob, q) || strings.Contains(blob, slugify(query))
	}
}
