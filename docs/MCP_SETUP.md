<<<<<<< HEAD
# MCP setup (Duffel travel server)

Connect Cursor (or Claude Desktop) to official flight/hotel search via [duffel-mcp](https://github.com/bokangsibolla/duffel-mcp). Uses **Duffel's REST API** — no Akamai, no browser session. Safe on Mac residential or any network.

## Prerequisites
=======
# MCP setup — agentic-travel

Configure Cursor MCP servers for **aggregate travel search** (Duffel, Amadeus) and **WAF brand automation** (`cursor-ide-browser`). See [MCP_LOCAL_AUDIT.md](./MCP_LOCAL_AUDIT.md) for what was available locally before this setup.

---

## 1. Workspace config (committed)

Project file: [`.cursor/mcp.json`](../.cursor/mcp.json)

```json
{
  "mcpServers": {
    "cursor-ide-browser": {},
    "playwright": {
      "command": "npx",
      "args": ["-y", "@playwright/mcp@latest"],
      "description": "Browser automation fallback when cursor-ide-browser is unavailable"
    }
  }
}
```

| Server | Purpose |
|--------|---------|
| **`cursor-ide-browser`** | Cursor built-in browser MCP — enable for WAF partial CLIs (`melia`, `nh`, `marriott`, `easyjet`, …). Empty `{}` opts in the built-in server for this workspace (same pattern as `suapea`, `firstmate`, `agentic-commerce` project caches). |
| **`playwright`** | Fallback when built-in browser MCP fails to register ([Cursor #3878](https://github.com/cursor/cursor/issues/3878)). Tool names align with `cursor-ide-browser`. |

**After editing:** restart Cursor or reload the window. Verify under **Settings → Tools & MCP** — `cursor-ide-browser` should list 16 tools (`browser_navigate`, `browser_snapshot`, `browser_cdp`, …).

Copy [`.cursor/mcp.json.example`](../.cursor/mcp.json.example) for Duffel-only config; merge entries into `.cursor/mcp.json` when API keys are ready (section 3).

---

## 2. Browser MCP — WAF brands (melia / nh / marriott)

Use when `{slug} session doctor --json` reports **`blocked`** and headed `session chrome` is impractical in the agent host.

### Session pattern (all three)

```
browser_navigate → wait → browser_snapshot → extract results (page or network)
```

| Step | Tool | Notes |
|------|------|-------|
| 1 | `browser_navigate` | Brand start URL (see registry) |
| 2 | `browser_lock` | After tab exists; unlock when done |
| 3 | Wait | Short CDP poll or 6–12s for Akamai settle — avoid long blind waits |
| 4 | `browser_snapshot` | Required before click/type; get element `ref`s |
| 5 | Interact | `browser_type` / `browser_click` if search form needed |
| 6 | Extract | **Primary:** `browser_cdp` with `Network.enable` + filter XHR JSON. **Fallback:** parse `browser_snapshot` DOM / JSON-LD |

### Per-brand URLs and filters

| Slug | Navigate | Network filter | Adapter |
|------|----------|----------------|---------|
| **melia** | `https://www.melia.com/es/hoteles` | `/services/search/hotels/v2/search` | `bridge/browser-mcp/adapters/melia.mjs` |
| **nh** | `https://www.nh-hotels.com/es/hoteles/espana/madrid` | `/nh/es/api/v1/hotels/search` | `bridge/browser-mcp/adapters/nh.mjs` |
| **marriott** | `findHotels.mi?destinationAddress.city=London&…` | `findHotels` or DOM | `bridge/browser-mcp/adapters/marriott.mjs` |

Full playbook: [bridge/browser-mcp/prompts/madrid-london.md](../bridge/browser-mcp/prompts/madrid-london.md) · architecture: [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md).

### Example: Meliá Madrid

```
browser_navigate(url="https://www.melia.com/es/hoteles")
# wait ~8s for Akamai
browser_snapshot()
browser_type → "Madrid" → submit
browser_cdp(method="Network.enable")
# capture POST .../services/search/hotels/v2/search response body
```

Normalize:

```bash
node bridge/browser-mcp/adapters/melia.mjs --file /tmp/melia-bff.json --query Madrid
```

### Fast path when cookies are warm

If `~/.melia/cookies.json` has valid `_abck` + `bm_sz`, prefer CLI:

```bash
melia search --json Madrid
```

Browser MCP is the **reliability fallback**, not the default when cookies work.

---

## 3. Duffel MCP (flights + stays) — when API key available

### Prerequisites
>>>>>>> origin/loop-7/mcp-setup

- **Node.js 20+**
- **Duffel test token** (free): [duffel.com](https://duffel.com) → Dashboard → **Test mode** → Create access token (`duffel_test_…`)

<<<<<<< HEAD
## 1. Install the MCP server (vendored)

From the repo root:
=======
### Install vendored server
>>>>>>> origin/loop-7/mcp-setup

```bash
chmod +x mcp/install.sh scripts/mcp-travel-search.sh
./mcp/install.sh
```

<<<<<<< HEAD
This clones `bokangsibolla/duffel-mcp` into `mcp/vendor/duffel-mcp/` and runs `npm ci && npm run build`.

## 2. Export API token
=======
Clones `bokangsibolla/duffel-mcp` into `mcp/vendor/duffel-mcp/` and runs `npm ci && npm run build`.

### Add to `.cursor/mcp.json`

Merge into existing `mcpServers` (keep `cursor-ide-browser`):

```json
"duffel": {
  "command": "node",
  "args": ["${workspaceFolder}/mcp/vendor/duffel-mcp/dist/index.js"],
  "env": {
    "DUFFEL_ACCESS_TOKEN": "${env:DUFFEL_ACCESS_TOKEN}"
  }
}
```

Or copy from [`.cursor/mcp.json.example`](../.cursor/mcp.json.example) and set an absolute path to `dist/index.js`.
>>>>>>> origin/loop-7/mcp-setup

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_your_token_here
```

<<<<<<< HEAD
Add to `~/.zshrc` or a local `.env` (never commit tokens).

## 3. Cursor MCP config (Sintetiko-labs / agentic-travel)

Copy the example and replace the token placeholder:

```bash
mkdir -p .cursor
cp .cursor/mcp.json.example .cursor/mcp.json
# Edit .cursor/mcp.json — set DUFFEL_ACCESS_TOKEN
```

Example structure (see [.cursor/mcp.json.example](../.cursor/mcp.json.example)):

```json
{
  "mcpServers": {
    "duffel": {
      "command": "node",
      "args": ["/Users/fbelchi/github/agentic-travel/mcp/vendor/duffel-mcp/dist/index.js"],
      "env": {
        "DUFFEL_ACCESS_TOKEN": "duffel_test_xxx"
      }
    }
  }
}
```

Restart Cursor after saving `.cursor/mcp.json`. The agent can then call `search_flights`, `get_offer`, and `search_stays`.

**Path note:** Use an absolute path to `dist/index.js` on your machine. The example uses the standard clone path under `/Users/fbelchi/github/agentic-travel`.

## 4. Shell smoke: MAD → London
=======
Never commit real tokens. Add `DUFFEL_ACCESS_TOKEN` to `~/.zshrc` or a local `.env`.

### Smoke: MAD → London
>>>>>>> origin/loop-7/mcp-setup

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_xxx
./scripts/mcp-travel-search.sh
```

<<<<<<< HEAD
Defaults: `MAD` → `STN` (Ryanair hub), depart `2026-07-05`. Override:

```bash
MCP_FROM=MAD MCP_TO=LHR MCP_DEPART=2026-08-01 ./scripts/mcp-travel-search.sh
```

Direct MCP client (same backend):

```bash
node mcp/call-search-flights.mjs --from MAD --to STN --depart 2026-07-05
```

## 5. Compare with ryanair-cli

Run both on the same route/date (see [MCP_RELIABILITY.md](./MCP_RELIABILITY.md)):

```bash
# Aggregator (Duffel MCP — multi-carrier)
./scripts/mcp-travel-search.sh

# Airline-direct (Ryanair only)
./scripts/mac-build-cli.sh ryanair search --json --from MAD --to STN --depart 2026-07-05
```
=======
Tools exposed: `search_flights`, `get_offer`, `search_stays` (read-only v0.1).

Compare with airline CLI: [MCP_RELIABILITY.md](./MCP_RELIABILITY.md).

---

## 4. Amadeus MCP (optional) — when API keys available

Amadeus Self-Service **sunsets July 2026**; prefer Duffel for new work. Use Amadeus when you already have licensed keys or need GDS coverage Duffel lacks.

### Candidate server

[HaroldLeo/amadeus-mcp](https://github.com/HaroldLeo/amadeus-mcp) — OAuth2, broad flight/hotel surface (26+ tools).

### Add to `.cursor/mcp.json` (after clone + build)

```json
"amadeus": {
  "command": "node",
  "args": ["${workspaceFolder}/mcp/vendor/amadeus-mcp/dist/index.js"],
  "env": {
    "AMADEUS_CLIENT_ID": "${env:AMADEUS_CLIENT_ID}",
    "AMADEUS_CLIENT_SECRET": "${env:AMADEUS_CLIENT_SECRET}"
  }
}
```

```bash
export AMADEUS_CLIENT_ID=...
export AMADEUS_CLIENT_SECRET=...
```

Restart Cursor. Agent uses Amadeus for **city-level hotel/flight discovery**; fall back to brand CLIs for Spanish chains and LCCs ([MCP_VS_CLI.md](./MCP_VS_CLI.md)).

---

## 5. Routing summary

| Intent | MCP | CLI |
|--------|-----|-----|
| Multi-carrier MAD→LON | **Duffel** `search_flights` | `ryanair`, `vueling` for LCC-direct |
| City hotels (generic) | **Duffel** `search_stays` or **Amadeus** | — |
| Named Spanish chain | — | **`melia`**, `barcelo`, `nh`, … |
| WAF partial + doctor `blocked` | **`cursor-ide-browser`** bridge | `session chrome` on headed Mac |

---
>>>>>>> origin/loop-7/mcp-setup

## Alternatives considered

| MCP | Status |
|-----|--------|
<<<<<<< HEAD
| **duffel-mcp** | **Integrated** — read-only, flights + stays, test tokens |
| amadeus-mcp | Amadeus Self-Service sunset Jul 2026; OAuth2; heavier tool surface |
| travel-mcp-server | Amadeus + AviationStack; same sunset risk |
| Kiwi / Skyscanner | No maintained OSS MCP with API-key auth found |

=======
| **duffel-mcp** | **Integrated** (vendored) — read-only, flights + stays |
| **cursor-ide-browser** | **Enabled** in `.cursor/mcp.json` — WAF bridge |
| **playwright** | **Fallback** in `.cursor/mcp.json` |
| amadeus-mcp | Documented; OAuth2; sunset risk |
| Kiwi / Skyscanner | No maintained OSS MCP with API-key auth found |

---

>>>>>>> origin/loop-7/mcp-setup
## Troubleshooting

| Issue | Fix |
|-------|-----|
<<<<<<< HEAD
| `DUFFEL_ACCESS_TOKEN is required` | Export token; must start with `duffel_test_` or `duffel_live_` |
| `Cannot find module …/dist/index.js` | Run `./mcp/install.sh` |
| MCP tools not visible in Cursor | Restart Cursor; check `.cursor/mcp.json` path is absolute |
| Empty flight results | Try `STN`, `LGW`, or `LHR` instead of city code `LON`; confirm future date |
| Node version error | `node -v` must be ≥ 20 |

=======
| `cursor-ide-browser` tools missing | Restart Cursor; confirm `.cursor/mcp.json` has `"cursor-ide-browser": {}`; try Playwright fallback |
| `DUFFEL_ACCESS_TOKEN is required` | Export token; must start with `duffel_test_` or `duffel_live_` |
| `Cannot find module …/dist/index.js` | Run `./mcp/install.sh` |
| Empty flight results (Duffel) | Try `STN`, `LGW`, `LHR` instead of `LON`; confirm future date |
| Browser MCP Akamai interstitial | Wait 10s, re-navigate; see [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md) |
| Node version error | `node -v` must be ≥ 20 |

---

>>>>>>> origin/loop-7/mcp-setup
## Security

- Prefer **`duffel_test_`** tokens; test searches are free and cannot book live inventory.
- duffel-mcp v0.1 is **read-only** — no `create_order` exposed.
<<<<<<< HEAD
- Do not commit `.cursor/mcp.json` with real tokens (example file uses placeholder).
=======
- Do not commit `.cursor/mcp.json` with real tokens (use `${env:…}` or local overrides).
- Browser MCP + CLIs: **search-only**, local/residential IP, same ToS posture as reverse-engineered CLIs.

---

## Related

- [MCP_LOCAL_AUDIT.md](./MCP_LOCAL_AUDIT.md) — local inventory
- [MCP_VS_CLI.md](./MCP_VS_CLI.md) — hybrid architecture
- [AGENTS.md](../AGENTS.md) — agent routing rules
>>>>>>> origin/loop-7/mcp-setup
