// Command parallel-hotels fans out hotel CLIs concurrently and emits merged JSON.
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
	fs := flag.NewFlagSet("parallel-hotels", flag.ExitOnError)
	city := fs.String("city", "", "destination city (required)")
	binDir := fs.String("bin-dir", parallel.DefaultBinDir, "pre-built CLI binaries directory")
	timeout := fs.Duration("timeout", parallel.DefaultTimeout, "per-CLI timeout")
	workers := fs.Int("workers", 0, "max concurrent workers (0 = CPU cores)")
	slugsFlag := fs.String("slugs", strings.Join(parallel.DefaultHotelSlugs, ","), "comma-separated hotel slugs")
	_ = fs.Parse(os.Args[1:])

	if *city == "" {
		fmt.Fprintln(os.Stderr, "usage: parallel-hotels --city London")
		os.Exit(2)
	}

	slugs := splitSlugs(*slugsFlag)
	res, err := parallel.RunHotels(*city, slugs, *binDir, *timeout, *workers)
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
