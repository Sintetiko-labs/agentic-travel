#!/usr/bin/env bash
# Parallel hotel chain CLIs — 30s/source timeout, partial results.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
# shellcheck source=../mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"
# shellcheck source=../wave-search-common.sh
source "$ROOT/scripts/wave-search-common.sh"

CITY="" LIMIT=10
SLUGS=(travelodge hilton)

while [ $# -gt 0 ]; do
  case "$1" in
    --city) CITY="$2"; shift 2 ;;
    --limit) LIMIT="$2"; shift 2 ;;
    --slugs) IFS=',' read -r -a SLUGS <<< "$2"; shift 2 ;;
    *) echo "unknown arg: $1" >&2; exit 2 ;;
  esac
done

[ -n "$CITY" ] || {
  echo "usage: parallel-hotels.sh --city CITY [--limit N]" >&2
  exit 2
}

wave_init_tmp
WAVE_WALL_START="$(wave_now_ms)"

for slug in "${SLUGS[@]}"; do
  mac_cli_build_cached "$slug" "$ROOT" || true
  bin="$(mac_cli_cached_bin "$slug")"
  if [ ! -x "$bin" ]; then
    wave_register_skipped "$slug" "binary missing — run build-bins.sh"
    continue
  fi
  wave_run_bg "$slug" "$bin" search --json --limit "$LIMIT" "$CITY"
done

wave_wait_all

manifest="$(wave_build_manifest "$(
  cat <<EOF
{"city":"$CITY","limit":$LIMIT}
EOF
)")"

node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest"
