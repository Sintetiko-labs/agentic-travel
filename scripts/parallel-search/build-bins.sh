#!/usr/bin/env bash
# Pre-build signed CLIs for parallel fan-out (avoid go run cold start).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
# shellcheck source=../mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"

export AGENTIC_TRAVEL_BIN_CACHE="${AGENTIC_TRAVEL_BINS:-${AGENTIC_TRAVEL_BIN_CACHE}}"

SLUGS=(ryanair vueling volotea binter travelodge hilton)

echo "Building CLIs → $AGENTIC_TRAVEL_BIN_CACHE" >&2
mkdir -p "$AGENTIC_TRAVEL_BIN_CACHE"

for slug in "${SLUGS[@]}"; do
  mac_cli_build_cached "$slug" "$ROOT" || echo "warn: build failed for $slug" >&2
done

echo "Done. export AGENTIC_TRAVEL_BINS=$AGENTIC_TRAVEL_BIN_CACHE" >&2
