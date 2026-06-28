#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
source "$ROOT/scripts/mac-cli-common.sh"
start=$(date +%s); ok=0; fail=0; failed=()
for slug in "${LIVE_CLI_SLUGS[@]}"; do
  if mac_cli_build_cached "$slug" "$ROOT"; then ok=$((ok+1)); else fail=$((fail+1)); failed+=("$slug"); fi
done
elapsed=$(($(date +%s)-start))
echo "mac-build-all: $ok ok, $fail failed (${elapsed}s)" >&2
echo "cache: $AGENTIC_TRAVEL_BIN_CACHE" >&2
[ "$fail" -eq 0 ] || { echo "failed: ${failed[*]}" >&2; exit 1; }
