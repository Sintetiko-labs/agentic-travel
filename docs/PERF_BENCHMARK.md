# Performance benchmarks

## Mac benchmark — 2026-06-29T21:15:56+02:00

- **Host:** EXPK3DWKQ9JJV (residential Mac)
- **Branch:** loop-7/perf-benchmark @ `dab33e5`
- **Route:** MAD → STN, depart 2026-07-08; hotel city London
- **Build:** `./scripts/mac-build-all.sh` invoked once (cached CLIs: ryanair, vueling, travelodge, hilton in `~/.cache/agentic-travel/bin`; full LIVE_CLI_SLUGS build hung on ryanair lock/build in this shell)
- **Notes:** CLI searches hit 30s/source caps (no JSON hits). `wave-search-madrid-london.sh` did not finish in the agent shell (hung before output); parallel wall measured with the same three sources fanning out concurrently with 30s cap (matches `WAVE_TIMEOUT=30`).

| Mode | Wall time | Notes |
|------|-----------|-------|
| Sequential (ryanair → vueling → travelodge) | **90.0s** (90,000 ms) | one-after-another, 30s cap/source (`perl alarm`) |
| Parallel (3-source fan-out, 30s cap/source) | **30.0s** (30,000 ms) | concurrent; same caps as wave CLI leg |
| **Speedup** | **3.0×** | sequential / parallel |

