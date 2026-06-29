#!/usr/bin/env bash
# Parallel MCP fan-out: Duffel (if token) + Kiwi + Gondola HTTP.
# Merges into CombinedSearchResult on stdout.
#
# Env: MCP_FROM, MCP_TO, MCP_DEPART, MCP_CITY, MCP_CHECK_IN, MCP_CHECK_OUT
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
# shellcheck source=wave-search-common.sh
source "$ROOT/scripts/wave-search-common.sh"

FROM="${MCP_FROM:-MAD}"
TO="${MCP_TO:-STN}"
DEPART="${MCP_DEPART:-2026-07-05}"
CITY="${MCP_CITY:-London}"
CHECK_IN="${MCP_CHECK_IN:-$DEPART}"
CHECK_OUT="${MCP_CHECK_OUT:-$(wave_default_checkout "$DEPART")}"

trap wave_cleanup_tmp EXIT

wave_init_tmp
WAVE_WALL_START="$(wave_now_ms)"

wave_run_duffel_mcp "$FROM" "$TO" "$DEPART"
wave_run_kiwi_mcp "$FROM" "$TO" "$DEPART"
wave_run_gondola_mcp "$CITY" "$CHECK_IN" "$CHECK_OUT"

wave_wait_all
wave_add_mcp_fallbacks_for_errors "$FROM" "$TO" "$DEPART" "$CITY" "$CHECK_IN" "$CHECK_OUT"

manifest="$(wave_build_manifest "$(
  node -e "console.log(JSON.stringify({from:process.argv[1],to:process.argv[2],city:process.argv[3],depart:process.argv[4],check_in:process.argv[5],check_out:process.argv[6]}))" \
    "$FROM" "$TO" "$CITY" "$DEPART" "$CHECK_IN" "$CHECK_OUT"
)")"

node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest"
