#!/usr/bin/env bash
# Fan-out hotel search across brand CLIs (Mac residential IP only).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-/tmp/agentic-travel-bins}"
ORCH="$BIN_DIR/parallel-hotels"

if [ ! -x "$ORCH" ]; then
  echo "orchestrator missing; building bins…" >&2
  "$ROOT/scripts/parallel-search/build-bins.sh" travelodge hilton barcelo marriott melia nh
fi

exec "$ORCH" "$@"
