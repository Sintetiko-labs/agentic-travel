#!/usr/bin/env bash
# Smoke parent hotel APIs: Madrid + London per priority parent.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BUILD="$ROOT/scripts/mac-build-cli.sh"
PARENTS=(accor ihg hyatt marriott hilton wyndham bestwestern radisson)
CITIES=(Madrid London)
PASS=0; FAIL=0; SKIP=0
for slug in "${PARENTS[@]}"; do
  for city in "${CITIES[@]}"; do
    printf "== %s %s ... " "$slug" "$city"
    if out=$("$BUILD" "$slug" search --json --limit 3 "$city" 2>&1); then
      total=$(echo "$out" | python3 -c "import sys,json; print(json.load(sys.stdin).get('total',0))" 2>/dev/null || echo 0)
      if [ "${total:-0}" -gt 0 ] 2>/dev/null; then echo "PASS total=$total"; PASS=$((PASS+1)); else echo "FAIL"; FAIL=$((FAIL+1)); fi
    else
      if echo "$out" | grep -qiE 'akamai|session chrome|no hotels parsed'; then echo "BLOCKED"; SKIP=$((SKIP+1)); else echo "FAIL"; FAIL=$((FAIL+1)); fi
    fi
  done
done
echo "PASS=$PASS FAIL=$FAIL BLOCKED=$SKIP"
