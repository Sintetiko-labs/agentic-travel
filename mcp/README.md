# Travel MCP integration

Official travel APIs exposed to AI agents via [Model Context Protocol](https://modelcontextprotocol.io/) (MCP).

## Selected server: [duffel-mcp](https://github.com/bokangsibolla/duffel-mcp)

After comparing public MCP servers (Duffel, Amadeus, Kiwi, Gondola, multi-API aggregators), **loop 7** integrates:

| Server | Config | Auth | Best for |
|--------|--------|------|----------|
| **kiwi-com-flight-search** | `https://mcp.kiwi.com` | None | Flight metasearch, booking links |
| **gondola** | `https://mcp.gondola.ai/mcp` | None | Marriott/Hilton/Accor/… hotel search |
| **duffel** (vendored) | local Node | `DUFFEL_ACCESS_TOKEN` | Multi-carrier flights + stays |
| **cursor-ide-browser** | built-in | None | WAF partial CLIs |

| Server | Pros | Cons |
|--------|------|------|
| **duffel-mcp** (vendored) | Read-only; flights + stays; free test tokens | LCC gaps; requires install |
| **Kiwi MCP** (remote) | Official; no key; flexible dates | Search only; less LCC depth than CLIs |
| **Gondola MCP** (remote) | Major chain rates + points | No Spanish regional chains |
| [HaroldLeo/amadeus-mcp](https://github.com/HaroldLeo/amadeus-mcp) | Broad Amadeus coverage | Amadeus Self-Service **shuts down July 2026** |

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
├── call-mcp-http.mjs       # Streamable HTTP client (Kiwi, Gondola)
├── call-kiwi-search.mjs
├── call-gondola-search.mjs
├── merge-wave-result.mjs   # CombinedSearchResult merge
└── vendor/duffel-mcp/     # gitignored — created by install.sh
```

## When to use MCP vs CLIs

| Scenario | Use |
|----------|-----|
| Multi-carrier MAD→London comparison | **Kiwi MCP** or **Duffel MCP** (`search_flights`) |
| Ryanair-specific fare + booking URL | **`ryanair` CLI** |
| Akamai/WAF-protected brand sites | **`cursor-ide-browser`** or session CLIs |
| Hotels UK (Marriott, Hilton, Accor) | **Gondola MCP** or hotel CLIs / Duffel `search_stays` |

Full reliability notes: [docs/MCP_RELIABILITY.md](../docs/MCP_RELIABILITY.md).
