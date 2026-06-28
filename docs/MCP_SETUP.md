# MCP setup — agentic-travel (loop 7)

Unified Cursor MCP configuration for **Duffel** (vendored local), **Kiwi.com** and **Gondola** (hosted remotes), and **browser MCP** for WAF-blocked partial CLIs.

Related: [MCP_TRAVEL_INVENTORY.md](./MCP_TRAVEL_INVENTORY.md) · [MCP_VS_CLI.md](./MCP_VS_CLI.md) · [MCP_LOCAL_AUDIT.md](./MCP_LOCAL_AUDIT.md) · [MCP_RELIABILITY.md](./MCP_RELIABILITY.md) · [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md)

---

## Quick start

```bash
chmod +x mcp/install.sh scripts/mcp-travel-search.sh
./mcp/install.sh
cp .cursor/mcp.json.example .cursor/mcp.json
# Edit paths/token in .cursor/mcp.json — never commit real tokens
```

Restart Cursor after saving `.cursor/mcp.json`.

---

## 1. Duffel MCP (vendored, flights + stays)

Uses [duffel-mcp](https://github.com/bokangsibolla/duffel-mcp) via Duffel REST API — no Akamai, no browser session.

**Prerequisites:** Node.js 20+, free test token from [duffel.com](https://duffel.com) (`duffel_test_…`).

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_your_token_here
./mcp/install.sh
```

Tools: `search_flights`, `get_offer`, `search_stays` (read-only v0.1).

**Shell smoke (MAD → London):**

```bash
./scripts/mcp-travel-search.sh
node mcp/call-search-flights.mjs --from MAD --to STN --depart 2026-07-05
```

Use an **absolute** path to `mcp/vendor/duffel-mcp/dist/index.js` in `.cursor/mcp.json` (see [.cursor/mcp.json.example](../.cursor/mcp.json.example)).

---

## 2. Kiwi.com flight search (remote, no API key)

Official hosted MCP: `https://mcp.kiwi.com` — meta flight discovery, booking links per result.

```json
"kiwi": { "url": "https://mcp.kiwi.com" }
```

Search-only; for airline-direct LCC fares use brand CLIs (`ryanair`, `vueling`, …). Details: [MCP_TRAVEL_INVENTORY.md](./MCP_TRAVEL_INVENTORY.md).

---

## 3. Gondola hotel search (remote, no API key)

Multi-chain hotel availability: `https://mcp.gondola.ai/mcp`

```json
"gondola": { "url": "https://mcp.gondola.ai/mcp" }
```

Spanish regional chains — use brand CLIs. See [MCP_TRAVEL_INVENTORY.md](./MCP_TRAVEL_INVENTORY.md).

---

## 4. Browser MCP (WAF partials)

Enable **`cursor-ide-browser`** in `.cursor/mcp.json` when `session doctor` returns **`blocked`**:

```json
"cursor-ide-browser": {}
```

See [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md) and `bridge/browser-mcp/`.

---

## Troubleshooting

| Issue | Fix |
|-------|-----|
| `cursor-ide-browser` tools missing | Restart Cursor; try Playwright fallback in example |
| `DUFFEL_ACCESS_TOKEN is required` | Export `duffel_test_` or `duffel_live_` token |
| `Cannot find module …/dist/index.js` | Run `./mcp/install.sh` |
| Remote MCP unreachable | Check Kiwi/Gondola URLs in inventory |
| Node version | `node -v` ≥ 20 |

---

## Security

- Do not commit `.cursor/mcp.json` with live tokens.
- duffel-mcp v0.1 is read-only.
