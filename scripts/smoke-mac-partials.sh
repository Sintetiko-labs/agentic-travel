#!/usr/bin/env bash
# Mac Terminal smoke for WAF-heavy partial CLIs (chrome session + search).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
# shellcheck source=mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"

TIMEOUT="${SMOKE_PARTIALS_TIMEOUT:-180}"
CITY="${SMOKE_PARTIALS_CITY:-Madrid}"
CLIS=(melia iberostar)

classify_search() {
  local bin="$1"
  local out code
  set +e
  if command -v gtimeout >/dev/null 2>&1; then
    out=$(gtimeout "$TIMEOUT" "$bin" search --json "$CITY" --limit 3 2>&1)
  elif command -v timeout >/dev/null 2>&1; then
    out=$(timeout "$TIMEOUT" "$bin" search --json "$CITY" --limit 3 2>&1)
  else
    out=$("$bin" search --json "$CITY" --limit 3 2>&1)
  fi
  code=$?
  set -e
  if echo "$out" | grep -qi 'search not yet implemented'; then
    echo "stub|search not implemented"
    return
  fi
  if echo "$out" | grep -qiE 'akamai blocked|incapsula|cloudflare|access denied|HTTP 403'; then
    echo "waf|$(echo "$out" | head -1 | cut -c1-72)"
    return
  fi
  if [ "$code" -ne 0 ]; then
    echo "error|exit $code: $(echo "$out" | head -1 | cut -c1-60)"
    return
  fi
  if echo "$out" | python3 -c "import json,sys; d=json.load(sys.stdin); h=d.get('hotels') or []; sys.exit(0 if h and h[0].get('name') else 1)" 2>/dev/null; then
    n=$(echo "$out" | python3 -c "import json,sys; print(len(json.load(sys.stdin).get('hotels') or []))")
    echo "live|${n} hotels"
  else
    echo "partial|no hotel names in JSON"
  fi
}

printf "%-12s %-8s %-10s %s\n" "CLI" "BUILD" "SEARCH" "NOTES"
printf "%-12s %-8s %-10s %s\n" "---" "-----" "------" "-----"

for slug in "${CLIS[@]}"; do
  if ! mac_cli_build_cached "$slug" "$ROOT"; then
    printf "%-12s %-8s %-10s %s\n" "$slug" "FAIL" "skip" "build"
    continue
  fi
  bin="$(mac_cli_cached_bin "$slug")"
  printf "%-12s %-8s" "$slug" "OK"

  set +e
  if command -v gtimeout >/dev/null 2>&1; then
    gtimeout 240 "$bin" session chrome --replace --wait >/dev/null 2>&1
  elif command -v timeout >/dev/null 2>&1; then
    timeout 240 "$bin" session chrome --replace --wait >/dev/null 2>&1
  else
    "$bin" session chrome --replace --wait >/dev/null 2>&1
  fi
  set -e

  IFS='|' read -r status notes <<< "$(classify_search "$bin")"
  printf " %-10s %s\n" "$status" "$notes"
done

echo "SHA $(git -C "$ROOT" rev-parse HEAD 2>/dev/null || echo unknown)"
