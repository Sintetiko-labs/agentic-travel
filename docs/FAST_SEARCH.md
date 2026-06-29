# Fast search — agent decision tree

Agent-facing guide for **sub-15s** multi-source travel search on **Mac Mini M-series**. Use wave orchestrators — never sequential CLI loops.

> See also [AGENTS.md](../AGENTS.md) · [LEARNINGS_ECOMMARTINEZ.md](LEARNINGS_ECOMMARTINEZ.md).

---

## Decision tree

```
User query
    │
    ├─ Names ONE brand (e.g. "Ryanair MAD→STN")?
    │       └─ YES → single CLI: `{slug} search --json …`
    │
    ├─ N brands / "cheapest" / no airline named?
    │       └─ YES → parallel orchestrator:
    │                 ./scripts/parallel-flights.sh
    │                 ./scripts/parallel-hotels.sh
    │
    └─ Flights + hotels / MCP + CLI / multi-source?
            └─ YES → wave (same clock):
                      ./scripts/wave-search-full.sh
                      or ./scripts/wave-search-madrid-london.sh (MAD→LON defaults)
```

### Routing table

| Signal | Tool | Command |
|--------|------|---------|
| 1 brand named | CLI only | `ryanair search --json --from MAD --to STN --depart …` |
| Multi-airline / LCC sweep | Parallel CLIs | `./scripts/parallel-flights.sh --from MAD --to STN --depart …` |
| Multi-chain hotel city | Parallel CLIs | `./scripts/parallel-hotels.sh --city London` |
| MCP aggregators only | Parallel MCP | `./scripts/mcp-travel-search-parallel.sh` |
| Flights **and** hotels (any route) | Generic wave | `./scripts/wave-search-full.sh --from … --to … --city … --depart …` |
| MAD→LON preset | Hybrid wave | `./scripts/wave-search-madrid-london.sh` |

---

## Parallel search rules (agents)

1. **NEVER** run brand CLIs sequentially for multi-brand queries.
2. **ALWAYS** use `wave-search-full.sh` or `wave-search-madrid-london.sh` when combining MCP + CLI or flights + hotels.
3. **Same wave:** Duffel + Kiwi + Gondola + CLIs start together; 30s timeout per source.
4. **Output:** `wave-result.json` with `sources[].duration_ms`, `timed_out[]`, `flights[]`, `hotels[]` (never null).
5. **Mac residential only** — Terminal.app on Mac Mini, not CI/datacenter.

### One-time setup

```bash
./scripts/parallel-search/build-bins.sh
export DUFFEL_ACCESS_TOKEN=duffel_test_…   # optional
(cd mcp && npm ci)                          # for Kiwi/Gondola HTTP MCP
```

---

## Example: Madrid → London (flights + hotels)

```bash
cd /path/to/agentic-travel
./scripts/parallel-search/build-bins.sh

WAVE_DEPART=2026-07-05 WAVE_OUT=wave-result.json ./scripts/wave-search-madrid-london.sh
```

**Parallel sources (single wave):**

| Source | Role |
|--------|------|
| Duffel MCP | GDS/NDC aggregate (if `DUFFEL_ACCESS_TOKEN`) |
| Kiwi MCP | Metasearch via HTTP (`mcp.kiwi.com`) |
| Gondola MCP | Chain hotels via HTTP |
| Ryanair / Vueling CLI | LCC MAD→STN |
| Travelodge / Hilton CLI | London hotels |

Generic route:

```bash
./scripts/wave-search-full.sh --from MAD --to STN --city London --depart 2026-07-05
```

---

## Anti-patterns

| Don't | Do instead |
|-------|------------|
| Sequential MCP then CLI | `wave-search-full.sh` |
| `for s in ryanair vueling; do …` | `parallel-flights.sh` |
| Ignore `timed_out[]` | Retry named sources or use `mcp_agent_fallback[]` |
