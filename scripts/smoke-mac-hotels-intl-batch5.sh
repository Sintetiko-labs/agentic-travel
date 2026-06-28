#!/usr/bin/env bash
# Live smoke for international hotel batch 5 (London / Paris / Berlin).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
DOC="$ROOT/docs/SMOKE_MAC_HOTELS_INTL_BATCH5.md"
BIN_DIR="${TMPDIR:-/tmp}/hotels-intl-batch5-$$"
mkdir -p "$BIN_DIR"
CITIES=(London Paris Berlin)
CLIS=(citizenm hoxton wyndham ihg hyatt accor bestwestern radisson mamashelter 25hours easyhotel bbhotels numa sonder limehome)
TIMEOUT=180

classify_search() {
  local city="$1" bin="$2"
  local out code
  set +e
  if command -v gtimeout >/dev/null 2>&1; then
    out=$(gtimeout "$TIMEOUT" "$bin" search --json "$city" --limit 3 2>&1)
  elif command -v timeout >/dev/null 2>&1; then
    out=$(timeout "$TIMEOUT" "$bin" search --json "$city" --limit 3 2>&1)
  else
    out=$("$bin" search --json "$city" --limit 3 2>&1)
  fi
  code=$?
  set -e
  if echo "$out" | grep -qi 'search not yet implemented'; then
    echo stub
    return
  fi
  if [ "$code" -ne 0 ]; then
    echo partial
    return
  fi
  if echo "$out" | python3 -c "import json,sys; d=json.load(sys.stdin); sys.exit(0 if (d.get('hotels') or []) and d['hotels'][0].get('name') else 1)" 2>/dev/null; then
    echo live
  else
    echo partial
  fi
}

declare -A OVERALL
TMP_ROWS="$BIN_DIR/rows.txt"
: > "$TMP_ROWS"

for cli in "${CLIS[@]}"; do
  cli_dir="$ROOT/${cli}-cli"
  bin="$BIN_DIR/$cli"
  if ! (cd "$cli_dir" && go build -o "$bin" "./cmd/$cli" >/dev/null 2>&1); then
    OVERALL[$cli]=stub
    echo "$cli|build failed" >> "$TMP_ROWS"
    continue
  fi
  live=0 stub=0
  line="$cli"
  for city in "${CITIES[@]}"; do
    st=$(classify_search "$city" "$bin")
    line+="|$city:$st"
    case "$st" in live) live=$((live+1));; stub) stub=$((stub+1));; esac
  done
  echo "$line" >> "$TMP_ROWS"
  if [ "$live" -eq 3 ]; then OVERALL[$cli]=live
  elif [ "$stub" -eq 3 ]; then OVERALL[$cli]=stub
  else OVERALL[$cli]=partial
  fi
done

SHA=$(git -C "$ROOT" rev-parse HEAD 2>/dev/null || echo unknown)
DATE=$(date -u +"%Y-%m-%dT%H:%MZ")

{
  echo "# Smoke: Mac hotels international batch 5"
  echo
  echo "- Date (UTC): $DATE"
  echo "- Git SHA: \`$SHA\`"
  echo "- Cities: London, Paris, Berlin"
  echo
  echo "## Summary"
  echo
  echo "| CLI | Overall |"
  echo "|-----|---------|"
  for cli in "${CLIS[@]}"; do
    echo "| $cli | ${OVERALL[$cli]:-partial} |"
  done
  echo
  echo "## Per city"
  echo
  while IFS= read -r row; do
    name="${row%%|*}"
    rest="${row#*|}"
    echo "### $name"
    echo
    echo "| City | Status |"
    echo "|------|--------|"
    if [ "$rest" = "build failed" ]; then
      echo "| (all) | stub |"
    else
      IFS='|' read -ra parts <<< "$rest"
      for part in "${parts[@]}"; do
        [ -z "$part" ] && continue
        city="${part%%:*}"
        st="${part#*:}"
        echo "| $city | $st |"
      done
    fi
    echo
  done < "$TMP_ROWS"
} > "$DOC"

echo "Wrote $DOC"
echo "SHA $SHA"
printf '%-14s %s\n' CLI OVERALL
for cli in "${CLIS[@]}"; do
  printf '%-14s %s\n' "$cli" "${OVERALL[$cli]:-partial}"
done
