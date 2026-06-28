# Browser MCP bridge — WAF reliability layer

**Loop 7** prototype: use **Cursor `cursor-ide-browser` MCP** (or Playwright MCP fallback) as a reliability layer for Akamai/Incapsula-blocked partial CLIs, instead of embedding `chromedp` in every Go binary.

## Problem

Seven **partial** CLIs (`melia`, `nh`, `iberostar`, `marriott`, `easyjet`, `aireuropa`, `iberiaexpress`) reverse-engineer brand BFFs but fail without a real browser session:

| CLI | WAF | Probe without session | Data path |
|-----|-----|----------------------|-----------|
| **melia** | Akamai | HTTP 404 / Next shell | POST `/services/search/hotels/v2/search` |
| **nh** | Akamai | HTTP 403 | GET `/nh/es/api/v1/hotels/search` |
| **marriott** | Akamai | HTTP 403 | HTML `/search/findHotels.mi` |
| **easyjet** | Akamai | HTTP 403 | GET `/ejavailability/api/v5/availability/query` |
| iberostar | Akamai | HTTP 404 | POST `/api/graphql` |
| aireuropa | Akamai | HTTP 403 | Amadeus `dapi` redirect |
| iberiaexpress | Incapsula | HTTP 200 challenge | GET `/api/availability/v1/flights` |

Today, `travelkit/chrome` + `session chrome --wait` harvests cookies via **chromedp** attached to a local CDP port (`scripts/smoke-mac-cdp.py` mirrors the same flow in Python). That works on a headed Mac but is brittle in agent sandboxes and duplicates browser logic per CLI.

## Idea

```
┌──────────────────┐     browser_navigate + network/DOM     ┌─────────────────────┐
│ Agent orchestrator│ ──────────────────────────────────────▶│ Brand site (real TLS)│
│ (Cursor chat)     │     browser_network_requests (JSON)    │ Akamai passes        │
└────────┬─────────┘                                      └──────────┬──────────┘
         │                                                              │
         │  bridge/browser-mcp/adapters/{slug}.mjs                      │
         ▼                                                              │
┌──────────────────┐     HotelSearchResult / FlightSearchResult          │
│ Normalized JSON  │ ◀──────────────────────────────────────────────────┘
│ (travelkit types)│
└──────────────────┘
```

**Agent calls browser MCP → captures API JSON or DOM → adapter maps to `travelkit/types` → same `--json` shape as CLI.**

Go CLIs remain the fast path when `{PREFIX}_COOKIE` is warm; the bridge is the **fallback** when `session doctor` reports `blocked`.

## Feasibility

| Dimension | Assessment |
|-----------|------------|
| **WAF bypass** | **High** — real Chromium TLS + JS challenge execution beats uTLS/cookie replay for Akamai `_abck` rotation |
| **Schema parity** | **High** — adapters reuse the same BFF URLs and field mapping as Go clients; network capture returns identical JSON |
| **Agent integration** | **Medium** — requires `cursor-ide-browser` enabled in Cursor; [known registration bugs](https://github.com/cursor/cursor/issues/3878) → use `@playwright/mcp` as fallback |
| **Automation** | **Medium** — hotels need search UI interaction or direct XHR replay in-page; airlines often expose JSON on first navigation |
| **Latency** | **Low–Medium** — 15–45s per brand (WAF settle + search) vs &lt;2s with warm cookies |
| **Headless CI** | **Low** — interactive agent host only; not a replacement for `verify-clis.sh` |
| **Maintenance** | **Medium** — registry documents URL patterns; DOM fallback breaks on redesign |

**Verdict:** Feasible and **more reliable than uTLS-only** for the four priority Akamai brands below. Best as an **agent-side reliability layer**, not a Go library replacement.

### Priority partials (this prototype)

| Slug | Benefit | Extraction strategy |
|------|---------|---------------------|
| **melia** | BFF 404 without session; directory fallback is thin | Navigate → trigger search → capture POST BFF response |
| **nh** | Clean JSON API behind Akamai | Navigate `/es` → capture GET search API |
| **marriott** | HTML scrape needs residential session | Navigate `findHotels.mi` → parse JSON-LD / listing DOM |
| **easyjet** | ejavailability 403 without cookies | Navigate search UI MAD→LTN → capture availability API |

`iberostar`, `aireuropa`, `iberiaexpress` follow the same pattern; registry stubs included for later.

## MCP tools (cursor-ide-browser)

Tool descriptors are bundled with Cursor (not in repo `mcps/`). Standard surface:

| Tool | Role in bridge |
|------|----------------|
| `browser_navigate` | Open brand start URL or deep-linked search |
| `browser_wait_for` | Wait for results text / time after WAF |
| `browser_snapshot` | Accessibility tree for click/type (search forms) |
| `browser_click` / `browser_type` | Fill destination, dates, submit search |
| `browser_network_requests` | **Primary** — filter XHR/fetch JSON by URL substring |
| `browser_console_messages` | Debug Akamai / CSP failures |
| `browser_take_screenshot` | Evidence when parse fails |

**Workflow rule:** always `browser_snapshot` before click/type (element `ref` from snapshot).

If built-in browser MCP is unavailable, configure Playwright MCP (`npx @playwright/mcp`) — same tool names with optional `includePayloads` on network capture.

## Sample flow: Madrid → London (2026-07-15)

Scenario: UK hotels in London + easyJet MAD→London for loop-6/7 integration.

### 1. easyJet flights (MAD → LTN)

```
browser_navigate(url="https://www.easyjet.com/es/cheap-flights/spain/madrid-to-london")
browser_wait_for(time=5)
browser_snapshot()
# If form visible: fill origin MAD, dest LTN/STN/LGW, date 2026-07-15, click Buscar
browser_network_requests()
# Filter: url contains "ejavailability/api/v5/availability/query"
```

Pipe response body through:

```bash
node bridge/browser-mcp/adapters/easyjet.mjs --stdin \
  --origin MAD --dest LTN --depart 2026-07-15 \
  > /tmp/easyjet-mad-ltn.json
```

Expected shape: `FlightSearchResult` with `source: "ejavailability-browser-mcp"`, `flights[]` non-null.

### 2. Meliá hotels (Madrid — origin city)

```
browser_navigate(url="https://www.melia.com/es/hoteles")
browser_wait_for(time=8)
browser_type(element="search", ref="<ref>", text="Madrid")
browser_click(element="search button", ref="<ref>")
browser_network_requests()
# Filter: "/services/search/hotels/v2/search"
```

```bash
node bridge/browser-mcp/adapters/melia.mjs --stdin --query Madrid \
  > /tmp/melia-madrid.json
```

### 3. NH hotels (Madrid)

```
browser_navigate(url="https://www.nh-hotels.com/es/hoteles/espana/madrid")
browser_wait_for(time=6)
browser_network_requests()
# Filter: "/nh/es/api/v1/hotels/search"
```

```bash
node bridge/browser-mcp/adapters/nh.mjs --stdin --query Madrid \
  > /tmp/nh-madrid.json
```

### 4. Marriott hotels (London — destination)

```
browser_navigate(url="https://www.marriott.com/search/findHotels.mi?searchType=InCity&destinationAddress.city=London&destinationAddress.country=GB&fromDate=2026-07-15&toDate=2026-07-16")
browser_wait_for(time=10)
browser_snapshot()
# Prefer network JSON if exposed; else DOM / JSON-LD via adapter --html
```

```bash
node bridge/browser-mcp/adapters/marriott.mjs --html page.html --query London \
  > /tmp/marriott-london.json
```

### 5. Merge for agent

All outputs share [`travelkit/types`](../travelkit/types/types.go) fields. Agent compares `price`, `booking_url` / `hotel_url`, and tags `source` with `-browser-mcp` suffix.

## Comparison: chromedp vs browser MCP

| | `travelkit/chrome` (chromedp) | Browser MCP bridge |
|--|------------------------------|-------------------|
| Host | Go binary / `session chrome` | Cursor agent session |
| Cookie harvest | Yes (`_abck`, `bm_sz`) | Implicit in browser profile |
| API replay | Go HTTP client + cookies | In-browser fetch or network log |
| Sandbox agents | Often blocked (no Chrome) | Uses IDE browser pane |
| Output | Saves `~/.{slug}/cookies.json` | Direct JSON to chat / file |
| Reuse in CI | `smoke-mac-cdp.py` | Manual / agent only |

**Hybrid:** run browser MCP once → optionally persist cookies from `document.cookie` into `~/.melia/cookies.json` for subsequent fast CLI calls (future `bridge/browser-mcp/export-cookies.mjs`).

## Directory layout

```
bridge/browser-mcp/
├── README.md
├── registry.json          # per-slug URLs, network filters, DOM hints
├── adapters/
│   ├── melia.mjs
│   ├── nh.mjs
│   ├── marriott.mjs
│   └── easyjet.mjs
├── examples/              # golden outputs (travelkit shape)
│   ├── melia-madrid.json
│   ├── nh-madrid.json
│   ├── marriott-london.json
│   └── easyjet-mad-ltn.json
└── prompts/
    └── madrid-london.md   # copy-paste agent playbook
```

## Routing decision (agent)

```
IF slug IN (melia, nh, marriott, easyjet, …) AND session doctor = blocked
THEN bridge/browser-mcp playbook
ELSE IF {PREFIX}_COOKIE warm
THEN {slug} search --json (fast path)
ELSE session chrome --wait (headed Mac) OR browser MCP bridge
```

## Risks and mitigations

| Risk | Mitigation |
|------|------------|
| `cursor-ide-browser` not registered | Playwright MCP in `.cursor/mcp.json`; document in MCP_SETUP |
| Network tool omits response bodies | Use in-page `fetch()` via snapshot + evaluate, or Playwright `includePayloads` |
| Marriott DOM-only | Reuse `travelkit/parse` logic ported to JS (JSON-LD) |
| Rate / bot detection | Human-paced `browser_wait_for`; residential IP on agent host |
| ToS | Search-only, same as CLI; personal/research use |

## Next steps (loop 7+)

1. Wire `registry.json` into orchestrator skill / `AGENTS.md` routing table.
2. Add `export-cookies.mjs` to feed `travelkit/cookies` from browser session.
3. Port `iberostar` GraphQL + `iberiaexpress` Incapsula adapters.
4. Optional: thin Go `bridge` subcommand that shells to saved network JSON (test fixtures).

## Related

- [QA_PARTIALS.md](./QA_PARTIALS.md) — doctor status for partial CLIs
- [MCP_VS_CLI.md](./MCP_VS_CLI.md) — official MCP vs brand CLI architecture
- [scripts/smoke-mac-cdp.py](../scripts/smoke-mac-cdp.py) — CDP session smoke (chromedp equivalent)
