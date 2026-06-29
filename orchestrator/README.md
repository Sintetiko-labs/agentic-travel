# Parallel search orchestrator

Fan-out airline and hotel brand CLIs concurrently on the Mac Mini. Merges per-CLI JSON into `CombinedSearchResult` on stdout.

```bash
./scripts/parallel-search/build-bins.sh
./scripts/parallel-search/parallel-flights.sh --from MAD --to STN --depart 2026-07-05
./scripts/parallel-search/parallel-hotels.sh --city London
./scripts/parallel-search/benchmark-flights.sh
```

See `travelkit/types/combined.go` for the merged JSON schema.

Wave search (`wave.go`) remains available for Madrid→London MCP+CLI hybrid runs.
