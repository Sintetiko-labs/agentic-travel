#!/usr/bin/env bash
# Generic parallel wave: MCP (Duffel/Kiwi/Gondola) + CLI (ryanair, vueling, travelodge, hilton).
# Mac residential only — run from Terminal.app.
#
# Usage:
#   ./scripts/wave-search-full.sh --from MAD --to STN --city London --depart 2026-07-05
#   WAVE_OUT=wave-result.json ./scripts/wave-search-full.sh --from MAD --to STN --city London --depart 2026-07-05
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
# shellcheck source=mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"
# shellcheck source=wave-search-common.sh
source "$ROOT/scripts/wave-search-common.sh"

export AGENTIC_TRAVEL_BIN_CACHE="${AGENTIC_TRAVEL_BINS:-${AGENTIC_TRAVEL_BIN_CACHE}}"

FROM="" TO="" CITY="" DEPART=""
CHECK_IN="" CHECK_OUT=""
HOTEL_LIMIT="${WAVE_HOTEL_LIMIT:-10}"
FLIGHT_SLUGS=(ryanair vueling)
HOTEL_SLUGS=(travelodge hilton)

while [ $# -gt 0 ]; do
  case "$1" in
    --from) FROM="$2"; shift 2 ;;
    --to) TO="$2"; shift 2 ;;
    --city) CITY="$2"; shift 2 ;;
    --depart) DEPART="$2"; shift 2 ;;
    --check-in) CHECK_IN="$2"; shift 2 ;;
    --check-out) CHECK_OUT="$2"; shift 2 ;;
    --hotel-limit) HOTEL_LIMIT="$2"; shift 2 ;;
    --help|-h)
      echo "usage: wave-search-full.sh --from ORIGIN --to DEST --city CITY --depart YYYY-MM-DD" >&2
      exit 0
      ;;
    *) echo "unknown arg: $1" >&2; exit 2 ;;
  esac
done

FROM="${FROM:-${WAVE_FROM:-}}"
TO="${TO:-${WAVE_TO:-}}"
CITY="${CITY:-${WAVE_CITY:-}}"
DEPART="${DEPART:-${WAVE_DEPART:-}}"

[ -n "$FROM" ] && [ -n "$TO" ] && [ -n "$CITY" ] && [ -n "$DEPART" ] || {
  echo "usage: wave-search-full.sh --from ORIGIN --to DEST --city CITY --depart YYYY-MM-DD" >&2
  exit 2
}

CHECK_IN="${CHECK_IN:-$DEPART}"
CHECK_OUT="${CHECK_OUT:-$(wave_default_checkout "$DEPART")}"
OUT="${WAVE_OUT:-wave-result.json}"

trap wave_cleanup_tmp EXIT

echo "wave search: $FROM→$TO flights + $CITY hotels (depart $DEPART)" >&2

wave_init_tmp
WAVE_WALL_START="$(wave_now_ms)"

# MCP wave (parallel)
wave_run_duffel_mcp "$FROM" "$TO" "$DEPART"
wave_run_kiwi_mcp "$FROM" "$TO" "$DEPART"
wave_run_gondola_mcp "$CITY" "$CHECK_IN" "$CHECK_OUT"

# CLI wave (parallel, cached bins)
for slug in "${FLIGHT_SLUGS[@]}"; do
  mac_cli_build_cached "$slug" "$ROOT" || true
  bin="$(mac_cli_cached_bin "$slug")"
  if [ ! -x "$bin" ]; then
    wave_register_skipped "$slug" "binary missing — run ./scripts/parallel-search/build-bins.sh"
    continue
  fi
  wave_run_bg "$slug" "$bin" search --json --from "$FROM" --to "$TO" --depart "$DEPART"
done

for slug in "${HOTEL_SLUGS[@]}"; do
  mac_cli_build_cached "$slug" "$ROOT" || true
  bin="$(mac_cli_cached_bin "$slug")"
  if [ ! -x "$bin" ]; then
    wave_register_skipped "$slug" "binary missing — run ./scripts/parallel-search/build-bins.sh"
    continue
  fi
  wave_run_bg "$slug" "$bin" search --json --limit "$HOTEL_LIMIT" "$CITY"
done

wave_wait_all
wave_add_mcp_fallbacks_for_errors "$FROM" "$TO" "$DEPART" "$CITY" "$CHECK_IN" "$CHECK_OUT"

manifest="$(wave_build_manifest "$(
  node -e "console.log(JSON.stringify({from:process.argv[1],to:process.argv[2],city:process.argv[3],depart:process.argv[4],check_in:process.argv[5],check_out:process.argv[6]}))" \
    "$FROM" "$TO" "$CITY" "$DEPART" "$CHECK_IN" "$CHECK_OUT"
)")"

node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest" | tee "$OUT"
echo "wrote $OUT" >&2
