#!/usr/bin/env bash
# Madrid → London hybrid wave (MCP + CLI parallel fan-out).
# Writes wave-result.json in cwd (override with WAVE_OUT).
#
# Usage:
#   WAVE_DEPART=2026-07-05 ./scripts/wave-search-madrid-london.sh
#   WAVE_OUT=/tmp/mad-lon.json ./scripts/wave-search-madrid-london.sh
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"

export WAVE_FROM="${WAVE_FROM:-MAD}"
export WAVE_TO="${WAVE_TO:-STN}"
export WAVE_CITY="${WAVE_HOTELS:-${WAVE_CITY:-London}}"
export WAVE_DEPART="${WAVE_DEPART:-2026-07-05}"
export WAVE_OUT="${WAVE_OUT:-wave-result.json}"

exec "$ROOT/scripts/wave-search-full.sh" \
  --from "$WAVE_FROM" \
  --to "$WAVE_TO" \
  --city "$WAVE_CITY" \
  --depart "$WAVE_DEPART"
