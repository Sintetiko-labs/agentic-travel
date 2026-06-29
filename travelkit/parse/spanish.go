package parse

import (
	"regexp"
	"strings"
)

var (
	esCityHotelLinkRE = regexp.MustCompile(`href="(/es/hoteles/[a-z0-9-]+/[a-z0-9-]+/?)"`)
	esHotelPathRE     = regexp.MustCompile(`href="(/es/hoteles?/[a-z0-9][a-z0-9-]*/?)"`)
	esParadorLinkRE   = regexp.MustCompile(`href="(/es/parador(?:es)?/[a-z0-9-]+/?)"`)
	hotelNameCardRE   = regexp.MustCompile(`data-hotel-name="([^"]+)"[^>]*(?:data-provider-url|data-href)="([^"]+)"`)
	hotelNameCardRE2  = regexp.MustCompile(`data-hotel-name="([^"]+)"`)
	iberostarHotelRE  = regexp.MustCompile(`href="(/es/hoteles/[^"#?]+)"[^>]*>([^<]{3,80})</a>`)
	nhHotelLinkRE     = regexp.MustCompile(`href="(/es/hotel/[a-z0-9-]+/?)"`)
	bahiaHotelRE      = regexp.MustCompile(`href="(/es/hoteles?/[^"#?]+/?)"`)
	onlyYouHotelRE    = regexp.MustCompile(`href="(/es/hoteles?/[^"#?]+/?)"[^>]*>([^<]{3,80})</a>`)
)

// HotelsFromEsCityHotelLinks extracts /es/hoteles/{city}/{hotel} rows (Vincci-style).
func HotelsFromEsCityHotelLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range esCityHotelLinkRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) < 4 {
			continue
		}
		city, slug := parts[2], parts[3]
		if city == "" || slug == "" || isEsListingSegment(city) {
			continue
		}
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID: slug, Name: titleFromSlug(slug), Address: city, URL: url,
		})
	}
	return out
}

// HotelsFromEsHotelPathLinks extracts /es/hotel(s)/{slug} detail links.
func HotelsFromEsHotelPathLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range esHotelPathRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) < 3 {
			continue
		}
		slug := parts[len(parts)-1]
		if slug == "" || isEsListingSegment(slug) {
			continue
		}
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		city := ""
		if len(parts) >= 4 {
			city = parts[2]
		}
		out = append(out, HotelLD{
			ID: slug, Name: titleFromSlug(slug), Address: city, URL: url,
		})
	}
	return out
}

// HotelsFromHotelNameCards extracts Palladium-style data-hotel-name cards.
func HotelsFromHotelNameCards(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range hotelNameCardRE.FindAllStringSubmatch(html, -1) {
		name, rawURL := strings.TrimSpace(m[1]), m[2]
		url := absolutize(baseURL, rawURL)
		if name == "" || seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID: pathID(url), Name: name, URL: url,
		})
	}
	if len(out) > 0 {
		return out
	}
	for _, m := range hotelNameCardRE2.FindAllStringSubmatch(html, -1) {
		name := strings.TrimSpace(m[1])
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		out = append(out, HotelLD{
			ID: slug, Name: name, URL: absolutize(baseURL, "/es/hoteles/"+slug),
		})
	}
	return out
}

// HotelsFromParadoresLinks extracts Paradores detail links.
func HotelsFromParadoresLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range esParadorLinkRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		slug := pathID(path)
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID: slug, Name: titleFromSlug(slug), URL: url,
		})
	}
	if len(out) == 0 {
		out = HotelsFromEsHotelPathLinks(html, baseURL)
	}
	return out
}

// HotelsFromIberostarDirectory extracts hotel rows from Iberostar directory HTML.
func HotelsFromIberostarDirectory(html, baseURL, query string) []HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range iberostarHotelRE.FindAllStringSubmatch(html, -1) {
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
			Name: name, URL: absolutize(baseURL, path), ID: slugFromPath(path),
		})
	}
	if len(out) > 0 {
		return out
	}
	return filterJSONLDByQuery(HotelsFromJSONLD(html, baseURL), query)
}

// HotelsFromNHDirectory extracts NH hotel detail links from listing HTML.
func HotelsFromNHDirectory(html, baseURL, query string) []HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range nhHotelLinkRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		slug := pathID(path)
		if seen[slug] {
			continue
		}
		if q != "" && !strings.Contains(strings.ToLower(slug), q) && !strings.Contains(strings.ToLower(path), q) {
			continue
		}
		seen[slug] = true
		out = append(out, HotelLD{
			ID: slug, Name: titleFromSlug(slug), URL: absolutize(baseURL, path+"/"),
		})
	}
	if len(out) > 0 {
		return out
	}
	return filterJSONLDByQuery(HotelsFromJSONLD(html, baseURL), query)
}

// HotelsFromBahiaLinks extracts Bahía Príncipe resort links.
func HotelsFromBahiaLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range bahiaHotelRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		slug := pathID(path)
		if isEsListingSegment(slug) {
			continue
		}
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		out = append(out, HotelLD{
			ID: slug, Name: titleFromSlug(slug), URL: url,
		})
	}
	return out
}

// HotelsFromOnlyYouLinks extracts Only YOU / Room Mate style named anchor links.
func HotelsFromOnlyYouLinks(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range onlyYouHotelRE.FindAllStringSubmatch(html, -1) {
		path := strings.TrimSuffix(m[1], "/")
		name := strings.TrimSpace(m[2])
		slug := pathID(path)
		if isEsListingSegment(slug) {
			continue
		}
		url := absolutize(baseURL, path+"/")
		if seen[url] {
			continue
		}
		seen[url] = true
		if name == "" {
			name = titleFromSlug(slug)
		}
		out = append(out, HotelLD{
			ID: slug, Name: name, URL: url,
		})
	}
	if len(out) > 0 {
		return out
	}
	return HotelsFromEsCityHotelLinks(html, baseURL)
}

// HotelsFromSpanishBrandLinks tries city links, path links, cards, then JSON-LD.
func HotelsFromSpanishBrandLinks(html, baseURL, host string) []HotelLD {
	for _, fn := range []func(string, string) []HotelLD{
		HotelsFromEsCityHotelLinks,
		HotelsFromEsHotelPathLinks,
		HotelsFromHotelNameCards,
	} {
		if rows := fn(html, baseURL); len(rows) > 0 {
			return rows
		}
	}
	if rows := HotelsFromBrandHomeLinks(html, baseURL, host); len(rows) > 0 {
		return rows
	}
	return HotelsFromJSONLD(html, baseURL)
}

func filterJSONLDByQuery(rows []HotelLD, query string) []HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return rows
	}
	out := make([]HotelLD, 0, len(rows))
	for _, h := range rows {
		blob := strings.ToLower(strings.Join([]string{h.Name, h.Address, h.URL}, " "))
		if strings.Contains(blob, q) {
			out = append(out, h)
		}
	}
	return out
}

func isEsListingSegment(seg string) bool {
	switch strings.ToLower(seg) {
	case "hoteles", "hotel", "espana", "españa", "spain", "destinos", "ofertas", "reservas":
		return true
	default:
		return false
	}
}

func hostFromBase(baseURL string) string {
	h := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	return strings.TrimSuffix(h, "/")
}

// HotelsFromGlobalesLinks parses Globales listing HTML.
func HotelsFromGlobalesLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.globales.com")
}

func HotelsFromGrupotelLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.grupotel.com")
}

func HotelsFromHipotelsLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.hipotels.com")
}

func HotelsFromSenatorLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.senator.es")
}

func HotelsFromMedplayaLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.medplaya.com")
}

func HotelsFromZenitLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.zenithoteles.com")
}

func HotelsFromAbbaLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.abbahoteles.com")
}

func HotelsFromPorthotelsLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.porthotels.es")
}

func HotelsFromOnaLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.onahotels.com")
}

func HotelsFromBeliveLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.belivehotels.com")
}

func HotelsFromEveniaLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.eveniahotels.com")
}

func HotelsFromIlunionLinks(html, baseURL string) []HotelLD {
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.ilunionhotels.com")
}

func HotelsFromPetitPalaceLinks(html, baseURL string) []HotelLD {
	rows := HotelsFromEsCityHotelLinks(html, baseURL)
	if len(rows) > 0 {
		return rows
	}
	return HotelsFromSpanishBrandLinks(html, baseURL, "www.petitpalace.com")
}

func HotelsFromRoomMateLinks(html, baseURL string) []HotelLD {
	return HotelsFromOnlyYouLinks(html, baseURL)
}
