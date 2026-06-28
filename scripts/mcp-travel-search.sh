#!/usr/bin/env bash
# Thin wrapper: MAD → London flight search via Duffel MCP (search_flights).
# Requires: ./mcp/install.sh, DUFFEL_ACCESS_TOKEN, Node 20+.
# Run from Terminal.app (same as other agentic-travel scripts).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"

FROM="${MCP_FROM:-MAD}"
TO="${MCP_TO:-STN}"
DEPART="${MCP_DEPART:-2026-07-05}"
RETURN="${MCP_RETURN:-}"
ADULTS="${MCP_ADULTS:-1}"
DIRECT="${MCP_DIRECT:-}"

SERVER="$ROOT/mcp/vendor/duffel-mcp/dist/index.js"
if [ ! -f "$SERVER" ]; then
  echo "Duffel MCP not installed. Run: $ROOT/mcp/install.sh" >&2
  exit 1
fi

if [ -z "${DUFFEL_ACCESS_TOKEN:-}" ]; then
  echo "Set DUFFEL_ACCESS_TOKEN (duffel_test_… from https://duffel.com, Test mode)." >&2
  exit 1
fi

if [ ! -d "$ROOT/mcp/node_modules/@modelcontextprotocol/sdk" ]; then
  echo "MCP client deps missing. Run: (cd $ROOT/mcp && npm ci)" >&2
  exit 1
fi

args=(--from "$FROM" --to "$TO" --depart "$DEPART" --adults "$ADULTS")
[ -n "$RETURN" ] && args+=(--return "$RETURN")
[ -n "$DIRECT" ] && args+=(--direct)

echo "MCP search_flights: $FROM → $TO depart $DEPART" >&2
exec node "$ROOT/mcp/call-search-flights.mjs" "${args[@]}"
