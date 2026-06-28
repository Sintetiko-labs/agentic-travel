# MCP vs CLI reliability (MAD → London)

Compare **Duffel MCP** (official aggregated API) with **ryanair-cli** (reverse-engineered Ryanair API) on the same origin–destination–date.

## Route under test

| Field | Default |
|-------|---------|
| Origin | `MAD` (Madrid) |
| Destination | `STN` (London Stansted — primary Ryanair London airport) |
| Depart | `2026-07-05` |
| Passengers | 1 adult |

Also try `LHR`, `LGW`, `LTN` via Duffel for full London coverage; Ryanair CLI only returns Ryanair-operated routes.

## How to run comparison

From repo root, **Terminal.app**, residential IP (Ryanair WAF; Duffel is unaffected):

```bash
export DUFFEL_ACCESS_TOKEN=duffel_test_xxx
DATE=2026-07-05

# 1) Duffel MCP (multi-carrier aggregator)
node mcp/call-search-flights.mjs --from MAD --to STN --depart "$DATE" \
  | tee /tmp/mcp-mad-stn.json

# 2) Ryanair CLI (airline-direct)
./scripts/mac-build-cli.sh ryanair search --json --from MAD --to STN --depart "$DATE" \
  | tee /tmp/ryanair-mad-stn.json
```

Or use the wrapper:

```bash
./scripts/mcp-travel-search.sh
./scripts/mac-build-cli.sh ryanair search --json --from MAD --to STN --depart 2026-07-05
```

## Expected differences

| Dimension | Duffel MCP | ryanair-cli |
|-----------|------------|-------------|
| **Coverage** | Many carriers (Iberia, BA, Ryanair, …) | Ryanair only |
| **Auth** | `DUFFEL_ACCESS_TOKEN` (official) | Optional `RYANAIR_COOKIE` / Chrome session |
| **WAF / Akamai** | None (official API) | Yes — residential IP + cookies for live search |
| **Price source** | NDC/GDS offers via Duffel | Ryanair `farfnd` + booking API |
| **Booking URL** | Offer ID → Duffel flow | Direct `ryanair.com` deep link |
| **Empty results** | Rare if route exists on any airline | Empty if Ryanair doesn't fly route (e.g. MAD→BCN) |
| **Rate limits** | Duffel API quotas | ~1 req/min recommended |

## Interpretation

- **MCP** answers: "What flights exist MAD→London on this date across airlines?"
- **ryanair-cli** answers: "What does Ryanair charge right now, with a booking link?"

For the Madrid→London scenario (loop 6), use **both**: MCP for itinerary planning, Ryanair/Vueling/Iberia CLIs for airline-specific fares and session-gated live prices.

## Record results (template)

Fill after running with a valid `DUFFEL_ACCESS_TOKEN`:

| Source | Offers | Cheapest | Notes |
|--------|--------|----------|-------|
| Duffel MCP | _n_ | _EUR …_ | Includes non-Ryanair carriers |
| ryanair-cli | _n_ | _EUR …_ | Ryanair-only; may need `session chrome` if 403 |

_Last run: not executed in CI (requires user API token)._
