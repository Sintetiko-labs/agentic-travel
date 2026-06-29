// Command parallel-flights fans out airline CLIs concurrently and emits merged JSON.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fbelchi/agentic-travel/orchestrator/internal/parallel"
)

func main() {
	fs := flag.NewFlagSet("parallel-flights", flag.ExitOnError)
	from := fs.String("from", "", "origin IATA (required)")
	to := fs.String("to", "", "destination IATA (required)")
	depart := fs.String("depart", "", "departure date YYYY-MM-DD (required)")
	ret := fs.String("return", "", "return date YYYY-MM-DD (optional)")
	binDir := fs.String("bin-dir", parallel.DefaultBinDir, "pre-built CLI binaries directory")
	timeout := fs.Duration("timeout", parallel.DefaultTimeout, "per-CLI timeout")
	workers := fs.Int("workers", 0, "max concurrent workers (0 = CPU cores)")
	slugsFlag := fs.String("slugs", strings.Join(parallel.DefaultFlightSlugs, ","), "comma-separated airline slugs")
	_ = fs.Parse(os.Args[1:])

	if *from == "" || *to == "" || *depart == "" {
		fmt.Fprintln(os.Stderr, "usage: parallel-flights --from MAD --to STN --depart 2026-07-05")
		os.Exit(2)
	}

	slugs := splitSlugs(*slugsFlag)
	slugs = parallel.FilterFlightSlugs(slugs, *from, *to)

	res, err := parallel.RunFlights(*from, *to, *depart, *ret, slugs, *binDir, *timeout, *workers)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(res); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func splitSlugs(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
