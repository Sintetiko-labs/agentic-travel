package hotel

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/destination"
	"github.com/fbelchi/travelkit/parse"
	tktypes "github.com/fbelchi/travelkit/types"
)

// EsHotelPaths returns common Spanish hotel listing paths for a destination query.
func EsHotelPaths(query string) []string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(query), " ", "-"))
	paths := []string{"/es/hoteles", "/es/hoteles/", "/es/hotel", "/es/"}
	if slug != "" {
		for _, term := range destination.Expand(query) {
			t := strings.ToLower(strings.ReplaceAll(term, " ", "-"))
			paths = append([]string{
				"/es/hoteles/" + t,
				"/es/hoteles/" + t + "/",
				"/es/destinos/" + t,
				"/es/destinos/" + t + "/",
				"/es/hotel/" + t,
			}, paths...)
		}
	}
	return dedupePaths(paths)
}

// EsParadoresPaths returns Paradores listing paths.
func EsParadoresPaths(query string) []string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(query), " ", "-"))
	paths := []string{"/es/paradores", "/es/paradores/"}
	if slug != "" {
		for _, term := range destination.Expand(query) {
			t := strings.ToLower(strings.ReplaceAll(term, " ", "-"))
			paths = append([]string{"/es/paradores/" + t, "/es/paradores/" + t + "/"}, paths...)
		}
	}
	return dedupePaths(paths)
}

// IberostarDirectoryPaths returns Iberostar directory fallback paths.
func IberostarDirectoryPaths(query string) []string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(query), " ", "-"))
	paths := []string{"/es/hoteles", "/es/hoteles/espana"}
	if slug != "" {
		paths = append([]string{"/es/hoteles/espana/" + slug}, paths...)
	}
	return dedupePaths(paths)
}

// NHDirectoryPaths returns NH hotel directory fallback paths.
func NHDirectoryPaths(query string) []string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(query), " ", "-"))
	paths := []string{"/es/hoteles", "/es/hoteles/"}
	if slug != "" {
		for _, term := range destination.Expand(query) {
			t := strings.ToLower(strings.ReplaceAll(term, " ", "-"))
			paths = append([]string{"/es/hoteles/" + t, "/es/hoteles/" + t + "/"}, paths...)
		}
	}
	return dedupePaths(paths)
}

// SpanishHTMLSearch fetches listing pages and parses hotel rows.
func SpanishHTMLSearch(
	fetch func(path string) (string, error),
	baseURL string,
	paths []string,
	parseFn func(html, baseURL string) []parse.HotelLD,
	query string,
) ([]parse.HotelLD, error) {
	var rows []parse.HotelLD
	seen := map[string]bool{}
	var lastErr error
	for _, p := range paths {
		html, err := fetch(p)
		if err != nil {
			lastErr = err
			continue
		}
		for _, h := range parseFn(html, baseURL) {
			key := h.URL
			if key == "" {
				key = h.ID
			}
			if key == "" || seen[key] {
				continue
			}
			seen[key] = true
			rows = append(rows, h)
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	if lastErr != nil {
		return nil, lastErr
	}
	return rows, nil
}

// FilterSpanishQuery filters rows by destination query and aliases.
func FilterSpanishQuery(rows []parse.HotelLD, query string) []parse.HotelLD {
	filtered := FilterHotelLD(rows, query)
	if len(filtered) > 0 {
		return filtered
	}
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if destination.MatchQuery(query, h.Name, h.URL, h.Address, h.ID) {
			out = append(out, h)
		}
	}
	return out
}

// SpanishSearchOpts configures a Spanish hotel HTML search.
type SpanishSearchOpts struct {
	Query, Brand, BaseURL, Source, DefaultBrand string
	Paths                                       []string
	Parse                                       func(html, baseURL string) []parse.HotelLD
	Page, PageSize                              int
}

// SearchSpanishHTML runs listing scrape + filter + pagination.
func SearchSpanishHTML(fetch func(path string) (string, error), opts SpanishSearchOpts) (*tktypes.HotelSearchResult, error) {
	rows, err := SpanishHTMLSearch(fetch, opts.BaseURL, opts.Paths, opts.Parse, opts.Query)
	if err != nil {
		return nil, fmt.Errorf("search %q: %w", opts.Query, err)
	}
	filtered := FilterSpanishQuery(rows, opts.Query)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels — try Madrid or Barcelona", opts.Query)
	}
	brand := opts.Brand
	if brand == "" {
		brand = opts.DefaultBrand
	}
	return LDToResult(filtered, opts.Query, opts.Page, opts.PageSize, brand, opts.BaseURL, opts.Source), nil
}

func dedupePaths(paths []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(paths))
	for _, p := range paths {
		if p == "" || seen[p] {
			continue
		}
		seen[p] = true
		out = append(out, p)
	}
	return out
}
