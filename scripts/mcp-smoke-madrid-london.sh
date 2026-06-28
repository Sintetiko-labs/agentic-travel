#!/usr/bin/env bash
# MCP smoke playbook: Madrid → London (flights + hotels)
#
# This script documents the MCP tool sequence for Cursor agents. It does NOT
# invoke remote MCP servers from bash — use CallMcpTool / agent MCP integration.
#
# Prerequisites:
#   - .cursor/mcp.json configured (Kiwi, Gondola, Duffel, cursor-ide-browser)
#   - DUFFEL_ACCESS_TOKEN exported for Duffel tools
#   - ./mcp/install.sh completed for vendored duffel-mcp
#
# Usage:
#   ./scripts/mcp-smoke-madrid-london.sh          # print agent instructions
#   ./scripts/mcp-smoke-madrid-london.sh --duffel # also run local Duffel client
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"

FROM="${SMOKE_FROM:-MAD}"
TO="${SMOKE_TO:-STN}"
DEPART="${SMOKE_DEPART:-2026-07-15}"
RETURN="${SMOKE_RETURN:-}"
CHECK_IN="${SMOKE_CHECK_IN:-2026-07-15}"
CHECK_OUT="${SMOKE_CHECK_OUT:-2026-07-18}"
CITY="${SMOKE_CITY:-London}"
GUESTS="${SMOKE_GUESTS:-2}"

print_playbook() {
  cat <<EOF
================================================================================
MCP smoke: ${FROM} → ${TO} flights + ${CITY} hotels
================================================================================

Run these MCP tool calls from the agent (Cursor CallMcpTool). Adjust dates if stale.

── 1. Flights — Kiwi.com (no auth) ──────────────────────────────────────────
Server:  kiwi-com-flight-search
Tool:    search-flight
Args:    origin=${FROM}, destination=${TO}, departureDate=${DEPART}$([ -n "$RETURN" ] && echo ", returnDate=${RETURN}" || echo ""), adults=1, cabin=economy
Expect:  ranked offers with carrier, price, duration, kiwi.com booking links
Note:    Try STN, LTN, LGW, LHR if city code fails; ±3 day flex supported by Kiwi

── 2. Flights — Duffel (requires DUFFEL_ACCESS_TOKEN) ───────────────────────
Server:  duffel
Tool:    search_flights
Args:    slices=[{origin=${FROM}, destination=${TO}, departure_date=${DEPART}}], passengers=[{type=adult}], cabin_class=economy
Expect:  offer_id, segments[], total_amount, booking handoff via get_offer
Fallback CLI: ryanair search --json --from ${FROM} --to ${TO} --depart ${DEPART}

── 3. Hotels — Gondola (no auth; major chains) ─────────────────────────────
Server:  gondola
Tool:    search_hotels
Args:    location="${CITY}", check_in=${CHECK_IN}, check_out=${CHECK_OUT}, guests=${GUESTS}
Expect:  Marriott/Hilton/Hyatt/IHG/Accor/Wyndham availability, member rates, booking links
Note:    Prefer over marriott-cli/hilton-cli when Akamai blocks CLI (doctor: blocked)

── 4. Hotels — Duffel stays (optional) ─────────────────────────────────────
Server:  duffel
Tool:    search_stays
Args:    location near ${CITY}, check_in=${CHECK_IN}, check_out=${CHECK_OUT}, guests=${GUESTS}
Expect:  aggregated stay offers (broader than Gondola chain focus)

── 5. WAF brands — cursor-ide-browser (when CLI blocked) ───────────────────
Server:  cursor-ide-browser
Tools:   browser_navigate → browser_lock → wait → browser_snapshot → browser_cdp
Melía:   https://www.melia.com/es/hoteles — filter POST .../services/search/hotels/v2/search
NH:      https://www.nh-hotels.com/es/hoteles/espana/madrid
Marriott: findHotels.mi?destinationAddress.city=${CITY}
EasyJet: https://www.easyjet.com — ${FROM}→${TO} on ${DEPART}
Fallback server: playwright (same tool names)

── 6. Named Spanish brands — CLI only ───────────────────────────────────────
When user names a chain, skip aggregate MCP:
  melia search --json Madrid
  nh search --json London
  barcelo search --json Barcelona
See scripts/groups.json for slug → binary mapping.

── Success criteria ─────────────────────────────────────────────────────────
[ ] Kiwi returns ≥1 ${FROM}→${TO} offer for ${DEPART}
[ ] Gondola returns ≥1 ${CITY} hotel for ${CHECK_IN}–${CHECK_OUT}
[ ] Duffel search_flights succeeds if DUFFEL_ACCESS_TOKEN is set
[ ] Agent normalizes to travelkit shape: flights[], hotels[] (never null)

Docs: docs/MCP_SETUP.md · docs/MCP_TRAVEL_INVENTORY.md · AGENTS.md
================================================================================
EOF
}

print_playbook

if [[ "${1:-}" == "--duffel" ]]; then
  echo "" >&2
  echo "Running local Duffel client (step 2 only)…" >&2
  MCP_FROM="$FROM" MCP_TO="$TO" MCP_DEPART="$DEPART" MCP_RETURN="$RETURN" \
    "$ROOT/scripts/mcp-travel-search.sh"
fi
