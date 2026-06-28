# Fast search — agent decision tree

Agent-facing guide for **sub-15s** multi-source travel search on **Mac Mini M-series**. Use orchestrator scripts — never sequential CLI loops.

> See also [orchestrator/README.md](../orchestrator/README.md), [AGENTS.md](../AGENTS.md).

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
    │                 NEVER run CLIs in a for-loop
    │
    └─ Aggregate discovery + confirm?
            └─ YES → hybrid wave (same clock):
                      1. Start Duffel MCP async (background)
                      2. In parallel: parallel-flights.sh + parallel-hotels.sh
                      3. Merge; rank top 3 MCP offers
                      4. Optional: CLI read/availability on top 3 only
```

### Routing table

| Signal | Tool | Command |
|--------|------|---------|
| 1 brand named | CLI only | `ryanair search --json --from MAD --to STN --depart …` |
| Multi-airline / LCC sweep | Parallel CLIs | `./scripts/parallel-flights.sh --from MAD --to LON --depart …` |
| Multi-chain hotel city | Parallel CLIs | `./scripts/parallel-hotels.sh --city London` |
| City-pair discovery | MCP first | `./scripts/mcp-travel-search.sh` |
| Flights **and** hotels (MAD→LON) | Hybrid wave | `./scripts/wave-search-madrid-london.sh` |
| Top-3 price confirm | CLI spot-check | `{slug} read --json <id>` on MCP shortlist only |

---

## Parallel search rules (agents)

1. **NEVER** run brand CLIs sequentially for multi-brand queries — latency stacks (4×30s = 2 min).
2. **ALWAYS** use `./scripts/parallel-flights.sh` or `./scripts/parallel-hotels.sh` for N-brand sweeps.
3. **Same wave:** launch MCP Duffel in **background** while parallel CLIs run; do not `wait` on MCP before starting CLIs.
4. **Concurrency:** max **8–10** on Mac Mini M-series (`workers = NumCPU()`, typically 8–10).
5. **Timeout:** **30s per source** (`--timeout 30`); return **partial** results.
6. **Output:** `CombinedSearchResult` JSON on stdout (`travelkit/types`); `flights[]` / `hotels[]` never `null`.

### One-time setup

```bash
./scripts/parallel-search/build-bins.sh   # pre-build signed CLIs → /tmp/agentic-travel-bins
export DUFFEL_ACCESS_TOKEN=duffel_test_…  # optional, for MCP wave
```

---

## Example: "Madrid London flights + hotel July"

**User intent:** MAD→London flights in July + London hotels for the stay.

**One command** (target **<15s** wall clock; partials OK):

```bash
cd /path/to/agentic-travel

export DUFFEL_ACCESS_TOKEN="${DUFFEL_ACCESS_TOKEN:-}"
export WAVE_DEPART=2026-07-05
export WAVE_FROM=MAD
export WAVE_HOTELS=London
export WAVE_OUT=/tmp/mad-lon-july.json

./scripts/wave-search-madrid-london.sh
```

**What runs in parallel (single wave):**

| Source | Role |
|--------|------|
| Duffel MCP | GDS/NDC aggregate MAD→STN |
| Ryanair CLI | MAD→STN LCC |
| Vueling CLI | MAD→LGW direct or MAD-BCN+BCN-LGW |
| Travelodge CLI | London hotels |
| Hilton CLI | London hotels |

**Alternative (Go orchestrator, same schema):**

```bash
cd orchestrator && WAVE_DEPART=2026-07-05 go run ./
```

**Split commands** when flights and hotels are queried separately:

```bash
./scripts/parallel-flights.sh --from MAD --to STN --depart 2026-07-05
./scripts/parallel-hotels.sh --city London --limit 10
```

**MCP + CLI same wave (manual):**

```bash
MCP_FROM=MAD MCP_TO=STN MCP_DEPART=2026-07-05 ./scripts/mcp-travel-search.sh > /tmp/mcp.json &
./scripts/parallel-flights.sh --from MAD --to STN --depart 2026-07-05 > /tmp/cli.json &
wait
```

---

## Anti-patterns

| Don't | Do instead |
|-------|------------|
| `for s in ryanair vueling; do $s search …; done` | `./scripts/parallel-flights.sh` |
| Wait for MCP, then start CLIs | Start MCP + CLIs in one wave |
| Retry all slugs on one timeout | Return partial; check per-source status |
| 15+ concurrent CLIs on Mac Mini | Cap at 8–10 (orchestrator semaphore) |
| Invent prices when all sources empty | Report empty + which sources timed out |

---

## Related

- [AGENTS.md](../AGENTS.md) — PARALLEL SEARCH PROTOCOL
- [MCP_VS_CLI.md](MCP_VS_CLI.md) — when MCP vs CLI
- [MCP_SETUP.md](MCP_SETUP.md) — Duffel token + install
- [orchestrator/README.md](../orchestrator/README.md) — Go fan-out internals
