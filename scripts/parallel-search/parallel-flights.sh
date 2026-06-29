#!/usr/bin/env bash
# Fan-out flight search across airline CLIs (Mac residential IP only).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-/tmp/agentic-travel-bins}"
ORCH="$BIN_DIR/parallel-flights"

if [ ! -x "$ORCH" ]; then
  echo "orchestrator missing; building bins…" >&2
  "$ROOT/scripts/parallel-search/build-bins.sh" ryanair vueling volotea aireuropa binter
fi

exec "$ORCH" "$@"
