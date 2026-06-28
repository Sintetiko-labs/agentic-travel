#!/usr/bin/env bash
# Verify all agentic-travel CLIs build and --help works.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
RESULTS="$ROOT/scripts/verify-results.txt"
: > "$RESULTS"

pass=0
fail=0

echo "Verifying CLIs in $ROOT ..."

for dir in "$ROOT"/*-cli; do
  [ -d "$dir" ] || continue
  slug=$(basename "$dir" -cli)
  bin="$dir/$slug"
  status="PASS"
  notes=""

  if ! (cd "$dir" && go mod tidy >/dev/null 2>&1 && go build -o "$slug" "./cmd/$slug" 2>/dev/null); then
    status="FAIL"
    notes="build"
    ((fail++)) || true
    echo "$slug,$status,$notes" >> "$RESULTS"
    echo "FAIL $slug (build)"
    continue
  fi

  if ! "$bin" help >/dev/null 2>&1; then
    status="FAIL"
    notes="help"
    ((fail++)) || true
    echo "$slug,$status,$notes" >> "$RESULTS"
    echo "FAIL $slug (help)"
    rm -f "$bin"
    continue
  fi

  ((pass++)) || true
  echo "$slug,$status," >> "$RESULTS"
  echo "PASS $slug"
  rm -f "$bin"
done

echo ""
echo "Summary: $pass passed, $fail failed"
echo "Results: $RESULTS"
[ "$fail" -eq 0 ]
