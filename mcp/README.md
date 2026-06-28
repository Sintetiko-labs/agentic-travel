# Travel MCP integration

Official travel APIs exposed to AI agents via [Model Context Protocol](https://modelcontextprotocol.io/) (MCP).

## Selected server: [duffel-mcp](https://github.com/bokangsibolla/duffel-mcp)

After comparing public MCP servers (Duffel, Amadeus, Kiwi, multi-API aggregators), **Duffel** is the recommended integration for `agentic-travel`:

| Server | Pros | Cons |
|--------|------|------|
| **duffel-mcp** (chosen) | Read-only by design; flights + stays; free test tokens; modern REST API; no OAuth dance | Not published reliably on npm yet — vendored via `install.sh` |
| [HaroldLeo/amadeus-mcp](https://github.com/HaroldLeo/amadeus-mcp) | Broad Amadeus coverage (flights, hotels, tours) | Amadeus Self-Service **shuts down July 2026**; OAuth2; 26+ tools (heavy) |
| [lev-corrupted/travel-mcp-server](https://github.com/lev-corrupted/travel-mcp-server) | Amadeus + AviationStack flight tracking | Same Amadeus sunset; extra AviationStack key |
| Kiwi / Skyscanner MCP | — | No maintained open-source MCP with API-key auth found |

Duffel complements the repo's **reverse-engineered airline CLIs** (Ryanair, Vueling, …): MCP returns **GDS/NDC aggregated offers** across carriers; CLIs return **airline-direct** fares and booking URLs.

## Tools exposed

| Tool | Use |
|------|-----|
| `search_flights` | IATA origin/destination, departure date, optional return, cabin, direct-only |
| `get_offer` | Full fare conditions, baggage, price for one offer ID |
| `search_stays` | Hotels near coordinates for check-in/out dates |

All tools are **read-only** in v0.1 (no accidental bookings).

## Quick install

```bash
cd /Users/fbelchi/github/agentic-travel
./mcp/install.sh
export DUFFEL_ACCESS_TOKEN=duffel_test_xxx   # from https://duffel.com dashboard (Test mode)
```

See [docs/MCP_SETUP.md](../docs/MCP_SETUP.md) for Cursor/Claude config and the MAD→London smoke script.

## Directory layout

```
mcp/
├── README.md              # this file
├── install.sh             # clone + build duffel-mcp into vendor/
├── package.json           # deps for call-search-flights.mjs
├── call-search-flights.mjs
└── vendor/duffel-mcp/     # gitignored — created by install.sh
```

## When to use MCP vs CLIs

| Scenario | Use |
|----------|-----|
| Multi-carrier MAD→London comparison | **Duffel MCP** (`search_flights`) |
| Ryanair-specific fare + booking URL | **`ryanair` CLI** |
| Akamai/WAF-protected brand sites | **Session CLIs** (`session chrome`) |
| Hotels UK (Travelodge, Hilton) | **hotel CLIs** or Duffel `search_stays` |

Full reliability notes: [docs/MCP_RELIABILITY.md](../docs/MCP_RELIABILITY.md).
