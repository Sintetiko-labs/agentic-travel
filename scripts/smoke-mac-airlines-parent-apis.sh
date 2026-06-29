#!/usr/bin/env bash
# Mac live smoke: parent-group airline CLIs (MAD→LHR one-way).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

FROM=MAD
TO=LHR
DEPART=2026-07-15
CLIS=(lufthansagroup airfranceklm britishairways turkish norwegian jet2 tui emirates qatar etihad wizzair iberia)

printf "%-18s %-8s %s\n" "CLI" "STATUS" "NOTES"
printf "%-18s %-8s %s\n" "---" "------" "-----"

for slug in "${CLIS[@]}"; do
  bin="/tmp/${slug}"
  notes=""
  status="ERROR"

  if ! (cd "${slug}-cli" && go build -o "$bin" "./cmd/${slug}" 2>/dev/null); then
    notes="build failed"
    printf "%-18s %-8s %s\n" "$slug" "ERROR" "$notes"
    continue
  fi

  out="$("$bin" search --json --from "$FROM" --to "$TO" --depart "$DEPART" 2>&1)" || rc=$?
  rc=${rc:-0}
  combined="$out"

  if echo "$combined" | grep -qi "not yet implemented"; then
    status="ERROR"
    notes="stub search"
  elif echo "$combined" | grep -qiE "akamai blocked|incapsula|cloudflare blocked|access denied"; then
    status="WAF"
    notes="$(echo "$combined" | head -1 | cut -c1-80)"
  elif [ "$rc" -ne 0 ]; then
    status="ERROR"
    notes="$(echo "$combined" | head -1 | cut -c1-80)"
  else
    flights=$(echo "$combined" | python3 -c "import json,sys; d=json.load(sys.stdin); print(len(d.get('flights') or []))" 2>/dev/null || echo "")
    if [ "$flights" = "" ]; then
      status="ERROR"
      notes="invalid JSON"
    elif [ "$flights" -gt 0 ]; then
      status="LIVE"
      notes="${flights} flights"
    else
      status="EMPTY"
      notes="0 flights (API OK)"
    fi
  fi
  unset rc
  printf "%-18s %-8s %s\n" "$slug" "$status" "$notes"
done
