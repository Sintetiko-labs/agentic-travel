#!/usr/bin/env bash
# Benchmark sequential vs parallel MAD→STN flight search.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-/tmp/agentic-travel-bins}"
FROM="${FROM:-MAD}"
TO="${TO:-STN}"
DEPART="${DEPART:-2026-07-05}"
SLUGS=(ryanair vueling volotea aireuropa)

"$ROOT/scripts/parallel-search/build-bins.sh" "${SLUGS[@]}" >/dev/null

ms_now() { python3 -c 'import time; print(int(time.time()*1000))'; }

SEQ_MS=$(ms_now)
for slug in "${SLUGS[@]}"; do
  "$BIN_DIR/$slug" search --json --from "$FROM" --to "$TO" --depart "$DEPART" >/dev/null 2>&1 || true
done
SEQ_MS=$(($(ms_now) - SEQ_MS))

PAR_T0=$(ms_now)
PAR_JSON=$("$BIN_DIR/parallel-flights" --from "$FROM" --to "$TO" --depart "$DEPART" \
  --slugs "$(IFS=,; echo "${SLUGS[*]}")" 2>/dev/null || echo '{}')
PAR_MS=$(($(ms_now) - PAR_T0))

PAR_TOTAL=$(echo "$PAR_JSON" | python3 -c 'import json,sys; d=json.load(sys.stdin); print(d.get("total",0))' 2>/dev/null || echo 0)
SPEEDUP=$(python3 -c "s=$SEQ_MS; p=max($PAR_MS,1); print(f'{s/p:.2f}')")

cat <<EOF
{
  "route": "${FROM}->${TO}",
  "depart": "${DEPART}",
  "airlines": $(printf '%s\n' "${SLUGS[@]}" | python3 -c 'import json,sys; print(json.dumps([l.strip() for l in sys.stdin if l.strip()]))'),
  "sequential_ms": ${SEQ_MS},
  "parallel_ms": ${PAR_MS},
  "parallel_flights_total": ${PAR_TOTAL},
  "speedup_ratio": "${SPEEDUP}"
}
EOF
