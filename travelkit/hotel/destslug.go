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

func AccorDestinationPath(query string) string {
	switch normCity(query) {
	case "london":
		return "/gb/united-kingdom/london.hotels.html"
	case "paris":
		return "/fr/france/paris.hotels.html"
	case "berlin":
		return "/de/germany/berlin.hotels.html"
	default:
		s := slugify(query)
		if s == "" {
			return ""
		}
		return "/hotels/" + s + ".hotels.html"
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
