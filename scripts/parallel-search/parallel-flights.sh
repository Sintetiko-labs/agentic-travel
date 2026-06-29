#!/usr/bin/env bash
# Parallel LCC flight CLIs — 30s/source timeout, partial results.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
# shellcheck source=../mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"
# shellcheck source=../wave-search-common.sh
source "$ROOT/scripts/wave-search-common.sh"

FROM="" TO="" DEPART=""
SLUGS=(ryanair vueling)

while [ $# -gt 0 ]; do
  case "$1" in
    --from) FROM="$2"; shift 2 ;;
    --to) TO="$2"; shift 2 ;;
    --depart) DEPART="$2"; shift 2 ;;
    --slugs) IFS=',' read -r -a SLUGS <<< "$2"; shift 2 ;;
    *) echo "unknown arg: $1" >&2; exit 2 ;;
  esac
done

[ -n "$FROM" ] && [ -n "$TO" ] && [ -n "$DEPART" ] || {
  echo "usage: parallel-flights.sh --from ORIGIN --to DEST --depart YYYY-MM-DD" >&2
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
  wave_run_bg "$slug" "$bin" search --json --from "$FROM" --to "$TO" --depart "$DEPART"
done

wave_wait_all

manifest="$(wave_build_manifest "$(
  cat <<EOF
{"from":"$FROM","to":"$TO","depart":"$DEPART"}
EOF
)")"

node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest"
