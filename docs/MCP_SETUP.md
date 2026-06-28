# MCP setup (Duffel travel server)

Connect Cursor (or Claude Desktop) to official flight/hotel search via [duffel-mcp](https://github.com/bokangsibolla/duffel-mcp). Uses **Duffel's REST API** — no Akamai, no browser session. Safe on Mac residential or any network.

## Prerequisites

- **Node.js 20+**
- **Duffel test token** (free): [duffel.com](https://duffel.com) → Dashboard → **Test mode** → Create access token (`duffel_test_…`)

## 1. Install the MCP server (vendored)

From the repo root:

```bash
chmod +x mcp/install.sh scripts/mcp-travel-search.sh
./mcp/install.sh
```

This clones `bokangsibolla/duffel-mcp` into `mcp/vendor/duffel-mcp/` and runs `npm ci && npm run build`.

## 2. Export API token

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_your_token_here
```

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

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_xxx
./scripts/mcp-travel-search.sh
```

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

## Alternatives considered

| MCP | Status |
|-----|--------|
| **duffel-mcp** | **Integrated** — read-only, flights + stays, test tokens |
| amadeus-mcp | Amadeus Self-Service sunset Jul 2026; OAuth2; heavier tool surface |
| travel-mcp-server | Amadeus + AviationStack; same sunset risk |
| Kiwi / Skyscanner | No maintained OSS MCP with API-key auth found |

## Troubleshooting

| Issue | Fix |
|-------|-----|
| `DUFFEL_ACCESS_TOKEN is required` | Export token; must start with `duffel_test_` or `duffel_live_` |
| `Cannot find module …/dist/index.js` | Run `./mcp/install.sh` |
| MCP tools not visible in Cursor | Restart Cursor; check `.cursor/mcp.json` path is absolute |
| Empty flight results | Try `STN`, `LGW`, or `LHR` instead of city code `LON`; confirm future date |
| Node version error | `node -v` must be ≥ 20 |

## Security

- Prefer **`duffel_test_`** tokens; test searches are free and cannot book live inventory.
- duffel-mcp v0.1 is **read-only** — no `create_order` exposed.
- Do not commit `.cursor/mcp.json` with real tokens (example file uses placeholder).
