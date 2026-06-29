#!/usr/bin/env bash
# Quick verify: wave scripts exist, shellcheck-friendly syntax, merge smoke.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
fail=0

check_executable() {
  local f="$1"
  if [ ! -f "$f" ]; then
    echo "MISSING $f" >&2
    fail=1
    return
  fi
  if [ ! -x "$f" ]; then
    chmod +x "$f" || true
  fi
  bash -n "$f" || { echo "SYNTAX $f" >&2; fail=1; }
  echo "OK $f"
}

for s in \
  "$ROOT/scripts/wave-search-common.sh" \
  "$ROOT/scripts/wave-search-full.sh" \
  "$ROOT/scripts/wave-search-madrid-london.sh" \
  "$ROOT/scripts/mcp-travel-search-parallel.sh" \
  "$ROOT/scripts/parallel-search/build-bins.sh" \
  "$ROOT/scripts/parallel-search/parallel-flights.sh" \
  "$ROOT/scripts/parallel-search/parallel-hotels.sh"
do
  check_executable "$s"
done

# Merge smoke (no network)
TMP="$(mktemp -d)"
cat >"$TMP/ryanair.json" <<'EOF'
{"flights":[{"id":"FR1","origin":"MAD","destination":"STN","depart_at":"2026-07-05T10:00:00Z","arrive_at":"2026-07-05T12:00:00Z","stops":0,"price":"49.99","currency":"EUR"}],"source":"ryanair"}
EOF

manifest="$(node -e "
console.log(JSON.stringify({
  query: { from: 'MAD', to: 'STN', depart: '2026-07-05' },
  sources: [{ id: 'ryanair', status: 'ok', duration_ms: 1200, file: process.argv[1] }],
  wall_clock_ms: 1500
}));
" "$TMP/ryanair.json")"

out="$(node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest")"
echo "$out" | node -e "
const j = JSON.parse(require('fs').readFileSync(0,'utf8'));
if (!Array.isArray(j.flights) || j.flights.length !== 1) throw new Error('merge smoke failed');
if (!Array.isArray(j.timed_out)) throw new Error('timed_out missing');
if (!j.sources[0].duration_ms) throw new Error('duration_ms missing');
console.log('OK merge-wave-result.mjs smoke');
"

rm -rf "$TMP"
[ "$fail" -eq 0 ]
