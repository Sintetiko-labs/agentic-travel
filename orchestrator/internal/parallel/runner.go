package parallel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	tktypes "github.com/fbelchi/travelkit/types"
	"golang.org/x/sync/errgroup"
)

const (
	DefaultBinDir  = "/tmp/agentic-travel-bins"
	DefaultTimeout = 30 * time.Second
)

// RunFlights fans out airline CLIs and merges JSON results.
func RunFlights(from, to, depart, ret string, slugs []string, binDir string, timeout time.Duration, workers int) (*tktypes.CombinedSearchResult, error) {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	if binDir == "" {
		binDir = DefaultBinDir
	}
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	start := time.Now()
	query := tktypes.CombinedQuery{
		Kind:   "flights",
		From:   strings.ToUpper(from),
		To:     strings.ToUpper(to),
		Depart: depart,
		Return: ret,
	}

	results := make([]sourceOutcome, len(slugs))
	sem := make(chan struct{}, workers)
	g, ctx := errgroup.WithContext(context.Background())

	for i, slug := range slugs {
		i, slug := i, slug
		g.Go(func() error {
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = runFlightCLI(ctx, binDir, slug, query, timeout)
			return nil
		})
	}
	_ = g.Wait()

	return mergeFlightOutcomes(query, results, workers, int(timeout.Seconds()), start), nil
}

// RunHotels fans out hotel CLIs and merges JSON results.
func RunHotels(city string, slugs []string, binDir string, timeout time.Duration, workers int) (*tktypes.CombinedSearchResult, error) {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	if binDir == "" {
		binDir = DefaultBinDir
	}
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	start := time.Now()
	query := tktypes.CombinedQuery{
		Kind: "hotels",
		City: city,
	}

	results := make([]sourceOutcome, len(slugs))
	sem := make(chan struct{}, workers)
	g, ctx := errgroup.WithContext(context.Background())

	for i, slug := range slugs {
		i, slug := i, slug
		g.Go(func() error {
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = runHotelCLI(ctx, binDir, slug, city, timeout)
			return nil
		})
	}
	_ = g.Wait()

	return mergeHotelOutcomes(query, results, workers, int(timeout.Seconds()), start), nil
}

type sourceOutcome struct {
	slug       string
	ok         bool
	durationMs int64
	total      int
	err        string
	flights    []tktypes.FlightHit
	hotels     []tktypes.HotelHit
}

func runFlightCLI(ctx context.Context, binDir, slug string, q tktypes.CombinedQuery, timeout time.Duration) sourceOutcome {
	t0 := time.Now()
	out := sourceOutcome{slug: slug}

	args := []string{
		"search", "--json",
		"--from", q.From,
		"--to", q.To,
		"--depart", q.Depart,
	}
	if q.Return != "" {
		args = append(args, "--return", q.Return)
	}

	raw, err := execCLI(ctx, binDir, slug, args, timeout)
	out.durationMs = time.Since(t0).Milliseconds()
	if err != nil {
		out.err = err.Error()
		return out
	}

	var res tktypes.FlightSearchResult
	if err := json.Unmarshal(raw, &res); err != nil {
		out.err = fmt.Sprintf("parse json: %v", err)
		return out
	}
	out.ok = true
	out.total = res.Total
	out.flights = res.Flights
	return out
}

func runHotelCLI(ctx context.Context, binDir, slug, city string, timeout time.Duration) sourceOutcome {
	t0 := time.Now()
	out := sourceOutcome{slug: slug}

	args := []string{"search", "--json", city}
	raw, err := execCLI(ctx, binDir, slug, args, timeout)
	out.durationMs = time.Since(t0).Milliseconds()
	if err != nil {
		out.err = err.Error()
		return out
	}

	var res tktypes.HotelSearchResult
	if err := json.Unmarshal(raw, &res); err != nil {
		out.err = fmt.Sprintf("parse json: %v", err)
		return out
	}
	out.ok = true
	out.total = res.Total
	out.hotels = res.Hotels
	return out
}

func execCLI(ctx context.Context, binDir, slug string, args []string, timeout time.Duration) ([]byte, error) {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	bin := filepath.Join(binDir, slug)
	if _, err := os.Stat(bin); err != nil {
		return nil, fmt.Errorf("%s: binary missing at %s (run scripts/parallel-search/build-bins.sh)", slug, bin)
	}

	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		if runCtx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("%s: timeout after %s", slug, timeout)
		}
		return nil, fmt.Errorf("%s: %s", slug, msg)
	}
	return stdout.Bytes(), nil
}

func mergeFlightOutcomes(query tktypes.CombinedQuery, outcomes []sourceOutcome, workers, timeoutSec int, start time.Time) *tktypes.CombinedSearchResult {
	combined := &tktypes.CombinedSearchResult{
		Query:      query,
		SearchedAt: start.UTC(),
		DurationMs: time.Since(start).Milliseconds(),
		Workers:    workers,
		TimeoutSec: timeoutSec,
	}

	var flights []tktypes.FlightHit
	for _, o := range outcomes {
		combined.Sources = append(combined.Sources, tktypes.SourceResult{
			Slug:       o.slug,
			OK:         o.ok,
			DurationMs: o.durationMs,
			Total:      o.total,
			Error:      o.err,
		})
		if o.ok {
			flights = append(flights, o.flights...)
		}
	}
	sort.Slice(flights, func(i, j int) bool {
		return priceKey(flights[i].Price) < priceKey(flights[j].Price)
	})
	combined.Flights = flights
	combined.Total = len(flights)
	return combined
}

func mergeHotelOutcomes(query tktypes.CombinedQuery, outcomes []sourceOutcome, workers, timeoutSec int, start time.Time) *tktypes.CombinedSearchResult {
	combined := &tktypes.CombinedSearchResult{
		Query:      query,
		SearchedAt: start.UTC(),
		DurationMs: time.Since(start).Milliseconds(),
		Workers:    workers,
		TimeoutSec: timeoutSec,
	}

	var hotels []tktypes.HotelHit
	for _, o := range outcomes {
		combined.Sources = append(combined.Sources, tktypes.SourceResult{
			Slug:       o.slug,
			OK:         o.ok,
			DurationMs: o.durationMs,
			Total:      o.total,
			Error:      o.err,
		})
		if o.ok {
			hotels = append(hotels, o.hotels...)
		}
	}
	combined.Hotels = hotels
	combined.Total = len(hotels)
	return combined
}

func priceKey(p string) string {
	return strings.TrimSpace(p)
}

// BinterAirports is the set of airports Binter serves.
var BinterAirports = map[string]struct{}{
	"TFN": {}, "TFS": {}, "LPA": {}, "FUE": {}, "ACE": {}, "VDE": {}, "SPC": {}, "GMZ": {},
	"MAD": {}, "AGP": {}, "SVQ": {}, "LEI": {}, "OVD": {}, "VLL": {}, "VGO": {},
}

// FilterFlightSlugs drops binter when neither endpoint is in its network.
func FilterFlightSlugs(slugs []string, from, to string) []string {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)
	_, fromOK := BinterAirports[from]
	_, toOK := BinterAirports[to]
	if fromOK && toOK {
		return slugs
	}
	out := make([]string, 0, len(slugs))
	for _, s := range slugs {
		if s == "binter" {
			continue
		}
		out = append(out, s)
	}
	return out
}

// DefaultFlightSlugs is the MAD→London LCC + legacy carrier set.
var DefaultFlightSlugs = []string{"ryanair", "vueling", "volotea", "aireuropa", "binter"}

// DefaultHotelSlugs is the UK/London hotel brand set.
var DefaultHotelSlugs = []string{"travelodge", "hilton", "barcelo", "marriott", "melia", "nh"}
