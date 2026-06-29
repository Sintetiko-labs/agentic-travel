# Learnings from @ecommartinez — scraping stack for agentic-travel

**Source:** [Thread by @ecommartinez](https://x.com/ecommartinez/status/2071360257906168143) (Jun 28, 2026) — unrolled via [Thread Reader](https://threadreaderapp.com/thread/2071360257906168143.html).

**Context:** Alejo Martínez (@ecommartinez) curates open-source tooling for AI builders. This thread lists **10 GitHub repos for web scraping** — not travel-specific, but directly relevant to our **partial CLIs**, **browser MCP bridge**, and **parallel search** goals on Mac Mini.

---

## Post summary

The thread promotes a **scraping toolkit** for agents that need structured data from arbitrary websites without paid APIs:

| # | Tool | Role |
|---|------|------|
| 1 | [Firecrawl](https://github.com/firecrawl/firecrawl) | Full-site crawl + JS render → clean structured output for LLMs |
| 2 | [Crawl4AI](https://github.com/unclecode/crawl4ai) | Fast URL → markdown/LLM-ready text; no API keys |
| 3 | [browser-use](https://github.com/browser-use/browser-use) | AI agent drives a real browser (click, scroll, login, forms) |
| 4 | [Crawlee](https://github.com/apify/crawlee) | Production scraping: proxies, retries, fingerprint spoofing, queues |
| 5 | [Scrapy](https://github.com/scrapy/scrapy) | Industrial-scale crawl framework (millions of pages) |
| 6 | [MarkItDown](https://github.com/microsoft/markitdown) | PDF/Office/HTML/image → markdown for LLM ingestion |
| 7 | [Scrapling](https://github.com/D4Vinci/Scrapling) | Adaptive selectors; anti-bot evasion |
| 8 | [scrcpy](https://github.com/Genymobile/scrcpy) | Remote Android control — scrape mobile-only apps |
| 9 | [AutoScraper](https://github.com/alirezamika/autoscraper) | Learn pattern from one example → generalize |
| 10 | [curl-impersonate](https://github.com/lwthiker/curl-impersonate) | curl with real browser TLS/JA3 fingerprints |

**Implicit thesis:** agents need **layered extraction** — from cheap HTTP impersonation → headless browser → full agentic browser — plus **parallel fan-out** when hunting data across many sources. The thread does **not** mention MCP, commerce APIs, or orchestration patterns; it is a **generic scraping catalog**.

**Related threads from same author (context, not in this URL):**

- **Apodex 1.0** (Jun 25): up to **150 parallel verifier agents** for deep research — separate verification from generation.
- **EverOS 1.0** (Jun 23): portable agent memory across Claude Code, Codex, OpenClaw.
- **SkyClaw** (May 28): 5-tool chained workflows without context loss.

These reinforce **parallel sub-agents** and **persistent memory** as complementary patterns to scraping.

---

## Techniques & architectures mentioned

| Technique | Where in thread | Relevance to agentic-travel |
|-----------|-----------------|------------------------------|
| **JS-rendered scraping** | Firecrawl, Crawl4AI | Hotel/airline SPAs behind Next.js shells (Meliá, NH) |
| **Agentic browser control** | browser-use | Overlaps our `cursor-ide-browser` bridge for WAF partials |
| **Anti-bot / fingerprint** | Crawlee, curl-impersonate, Scrapling | Akamai `_abck`, Incapsula on `melia`, `nh`, `easyjet` |
| **Proxy rotation & retries** | Crawlee | Resilience when residential IP rate-limits |
| **Mobile-only inventory** | scrcpy | Niche: airline/hotel apps without web BFF |
| **Pattern learning** | AutoScraper | Rapid adapter prototyping when BFF URL changes |
| **LLM-ready normalization** | MarkItDown, Crawl4AI | Post-process HTML/PDF rate sheets into `travelkit` types |
| **Parallel agents** | Apodex (adjacent thread) | Same wave as our Go orchestrator + MCP fan-out |

**Not mentioned but we use:** MCP (Duffel, Kiwi, Gondola), official aggregator APIs, domain-specific reverse-engineered BFF CLIs, Go `travelkit` types.

---

## Mapping to our stack

```
                    ┌─────────────────────────────────────────┐
                    │         User query (agent)               │
                    └────────────────────┬────────────────────┘
                                         │
         ┌───────────────────────────────┼───────────────────────────────┐
         │ MCP wave (official APIs)      │ CLI wave (brand BFFs)          │ Scraping layer (ecommartinez)
         ▼                               ▼                                ▼
  Kiwi / Duffel / Gondola          parallel-flights.sh              curl-impersonate (cheap TLS)
  mcp-travel-search.sh             parallel-hotels.sh               Crawlee (retry/proxy)
  wave-search-madrid-london.sh     Go orchestrator (8–10 workers)   browser-use / browser MCP (WAF)
                                   travelkit JSON out               Firecrawl/Crawl4AI (fallback HTML)
```

| Our component | ecommartinez analogue | Relationship |
|---------------|----------------------|--------------|
| **Duffel/Kiwi/Gondola MCP** | — | **We are ahead** — official structured APIs beat generic scraping for discovery |
| **parallel-flights/hotels.sh** | Apodex parallel agents | **Aligned** — same-wave fan-out, 30s timeout, partial results |
| **browser MCP bridge** | browser-use (#3) | **Overlapping** — we capture XHR JSON; browser-use adds autonomous form-fill |
| **Go CLIs + uTLS** | curl-impersonate (#10) | **Complementary** — impersonate could replace/customize TLS in cold-start path |
| **session chrome / chromedp** | browser-use + Crawlee fingerprint | **We have this** — headed cookie harvest on Mac |
| **Partial WAF brands** | Crawlee + Scrapling | **Gap** — no proxy pool or adaptive selectors in CLI layer yet |
| **wave-search** | — | **Unique** — MCP + CLI same clock; not in thread |

---

## What to adopt

### 1. curl-impersonate as CLI cold-start layer (high priority)

Before launching full browser MCP for Akamai brands, try **Chrome-fingerprint HTTP** for BFF endpoints documented in [BROWSER_MCP_BRIDGE.md](BROWSER_MCP_BRIDGE.md). Cheaper than 15–45s browser settle; may warm enough for `melia`/`nh` search APIs.

**Action:** Spike `travelkit/impersonate` wrapper; call from `session doctor` when `blocked` but before `session chrome`.

### 2. Crawlee retry/proxy patterns for partial CLIs (medium)

Crawlee's **automatic retries, session rotation, and queue semantics** map to our per-source 30s timeout + partial return model. We do not need full Crawlee in Go — port the **policy**: exponential backoff within timeout, mark source `degraded` not `failed`.

**Action:** Add `SourceStatus` retry hint in orchestrator output; document proxy env (`HTTP_PROXY`) for rate-limited slugs.

### 3. browser-use evaluation for WAF bridge (medium)

[browser-use](https://github.com/browser-use/browser-use) is the closest open-source peer to our **cursor-ide-browser** bridge. Worth a **Mac Mini headless PoC** for `melia` + `easyjet` when Cursor MCP is unavailable (CI, remote agent host).

**Action:** `bridge/browser-use-poc/` — same adapters as `bridge/browser-mcp/adapters/`; compare latency vs `cursor-ide-browser`.

### 4. Crawl4AI for non-BFF fallback pages (low)

When Marriott HTML scrape or chain directory pages are the only path, **Crawl4AI** converts DOM → markdown faster than custom parsers. Use only when network JSON capture fails.

**Action:** Optional post-processor in browser bridge adapters for DOM-only slugs.

### 5. Parallel MCP fan-out (high — extends loop 7)

Thread implies "scrape everything in parallel"; we parallelize **CLIs** but often call **one MCP at a time**. Run **Kiwi + Duffel + Gondola** in the same wave as `parallel-flights.sh` / `parallel-hotels.sh`.

**Action:** Extend `mcp-travel-search.sh` → `mcp-travel-search-parallel.sh` with `wait` merge into `CombinedSearchResult`.

### 6. Verifier pattern from Apodex (low — quality)

For price comparison across 10+ sources, add a **lightweight verifier step**: top-3 offers re-fetched via `read`/`availability` before ranking. Prevents stale cache / hallucinated merges.

**Action:** Optional `--verify-top 3` flag on wave script.

---

## What we already do better

| Area | ecommartinez thread | agentic-travel |
|------|---------------------|----------------|
| **Discovery** | Generic scrape any site | **Official MCP** (Kiwi, Duffel, Gondola) — structured, ToS-aware, booking links |
| **Spanish LCC / chains** | Would need custom scrapers per site | **Dedicated CLIs** (Ryanair, Vueling, Meliá, NH…) with BFF knowledge |
| **Orchestration** | Tool list only | **Documented protocol** — parallel scripts, <15s target, anti-patterns in AGENTS.md |
| **Output contract** | Unstructured markdown | **`travelkit` types** — `flights[]`, `hotels[]`, never null |
| **Hybrid wave** | Not covered | **MCP background + CLI same clock** (`wave-search-madrid-london.sh`) |
| **WAF strategy** | browser-use generically | **Tiered**: warm cookies → browser MCP XHR capture → DOM fallback |
| **Commerce path** | Scraping | **booking_url** from API or CLI — closer to agentic commerce than HTML extraction |

We should **not** replace MCP/CLI with Firecrawl/Scrapy for core travel search — latency and fragility are worse than our stack. Scraping tools are **fallback and enrichment** only.

---

## Concrete next steps — faster parallel search on Mac Mini

Target: **<15s** aggregated MAD→LON (or any multi-brand query) on M-series Mac Mini.

### Immediate (this sprint)

1. **Pre-build CLIs** — `./scripts/parallel-search/build-bins.sh` → `AGENTIC_TRAVEL_BINS=/tmp/agentic-travel-bins` (avoid `go run` cold start per fan-out).
2. **Same-wave MCP** — background all three aggregators:
   ```bash
   MCP_FROM=MAD MCP_TO=STN MCP_DEPART=2026-07-05 ./scripts/mcp-travel-search.sh > /tmp/mcp.json &
   ./scripts/parallel-flights.sh --from MAD --to STN --depart 2026-07-05 > /tmp/cli.json &
   wait && jq -s 'add' /tmp/mcp.json /tmp/cli.json
   ```
3. **Cap workers at `NumCPU()`** — 8–10 on Mac Mini; do not exceed (memory + Akamai rate limits).
4. **30s per-source timeout** — return partials; surface `timed_out: ["vueling"]` in JSON for agent retry logic.
5. **Cookie pre-warm cron** — `melia session chrome --wait` + `nh session chrome --wait` every 6h on Mac Mini so parallel hotel sweep skips WAF on first request.

### Short-term (loop 7+)

6. **curl-impersonate spike** on `nh` and `easyjet` BFF URLs — measure p50 latency vs browser bridge.
7. **`mcp-travel-search-parallel.sh`** — fan-out Kiwi + Duffel (if token) + Gondola with shared merge. ✅ Implemented in `scripts/mcp-travel-search-parallel.sh` + `scripts/wave-search-full.sh`.
8. **browser-use PoC** on Mac Mini headless Chromium for agent hosts without Cursor IDE.
9. **Verifier pass** — `read --json` on top 3 MCP offers only (not full N-brand sweep).
10. **Benchmark script** — `scripts/bench-wave.sh` records wall clock + per-source latency to `docs/bench-mac-mini.json`.

### Not recommended

- **Scrapy / Firecrawl** as primary flight/hotel search — overkill, slower than BFF CLIs.
- **scrcpy** unless a target brand is app-only (none in current `groups.json`).
- **150-agent Apodex-style** verification for every search — reserve for high-stakes price audit, not default sweep.

---

## References

- [AGENTS.md](../AGENTS.md) — PARALLEL SEARCH PROTOCOL
- [FAST_SEARCH.md](FAST_SEARCH.md) — decision tree & wave commands
- [BROWSER_MCP_BRIDGE.md](BROWSER_MCP_BRIDGE.md) — WAF partial strategy
- [MCP_VS_CLI.md](MCP_VS_CLI.md) — when MCP vs CLI
- Thread: https://x.com/ecommartinez/status/2071360257906168143

---

*Document created on branch `loop-7/learnings-ecommartinez` as part of loop 7 research.*
