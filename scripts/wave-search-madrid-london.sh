#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BUILD="$ROOT/scripts/mac-build-cli.sh"
MERGE="$ROOT/scripts/wave-merge.py"
TMP="${WAVE_TMP:-$(mktemp -d "${TMPDIR:-/tmp}/wave-search.XXXXXX")}"
trap 'rm -rf "$TMP"' EXIT
FROM="${WAVE_FROM:-MAD}"
DEPART="${WAVE_DEPART:-2026-07-15}"
CHECK_IN="${WAVE_CHECK_IN:-2026-07-15}"
CHECK_OUT="${WAVE_CHECK_OUT:-2026-07-18}"
OUT="${WAVE_OUT:-$ROOT/wave-result.json}"
ms_now() { python3 -c 'import time; print(int(time.time()*1000))'; }
run_timed() {
  local name="$1"; shift
  local out="$TMP/${name}.json" err="$TMP/${name}.err"
  local t0 t1 ms ok=1
  t0="$(ms_now)"
  if "$@" >"$out" 2>"$err"; then ok=1; else ok=0; fi
  t1="$(ms_now)"; ms=$((t1 - t0))
  printf '%s:%s:%s:%s\n' "$name" "$out" "$ms" "$ok" >>"$TMP/specs.txt"
}
run_vueling() {
  local direct_out="$TMP/vueling_direct.json" err="$TMP/vueling.err"
  local t0 t1 ms
  t0="$(ms_now)"
  if "$BUILD" vueling search --json --from "$FROM" --to LGW --depart "$DEPART" >"$direct_out" 2>"$err"; then
    if python3 -c "import json,sys; d=json.load(open(sys.argv[1])); sys.exit(0 if d.get('total',0)>0 else 1)" "$direct_out" 2>/dev/null; then
      t1="$(ms_now)"; ms=$((t1 - t0))
      printf 'vueling:%s:%s:1\n' "$direct_out" "$ms" >>"$TMP/specs.txt"
      return 0
    fi
  fi
  local leg1="$TMP/vueling_leg1.json" leg2="$TMP/vueling_leg2.json"
  local t_leg0 t_leg1 ms1
  t_leg0="$(ms_now)"
  "$BUILD" vueling search --json --from "$FROM" --to BCN --depart "$DEPART" >"$leg1" 2>>"$err" &
  p1=$!
  "$BUILD" vueling search --json --from BCN --to LGW --depart "$DEPART" >"$leg2" 2>>"$err" &
  p2=$!
  wait "$p1" "$p2"
  t_leg1="$(ms_now)"; ms1=$((t_leg1 - t_leg0))
  printf 'vueling_leg1:%s:%s:1\n' "$leg1" "$ms1" >>"$TMP/specs.txt"
  printf 'vueling_leg2:%s:%s:0\n' "$leg2" "$ms1" >>"$TMP/specs.txt"
}
wall_t0="$(ms_now)"
: >"$TMP/specs.txt"
if [ -n "${DUFFEL_ACCESS_TOKEN:-}" ]; then
  run_timed duffel node "$ROOT/mcp/call-search-flights.mjs" --from "$FROM" --to STN --depart "$DEPART" &
  pid_duffel=$!
else pid_duffel=""; fi
run_timed ryanair "$BUILD" ryanair search --json --from "$FROM" --to STN --depart "$DEPART" &
pid_ryanair=$!
run_vueling & pid_vueling=$!
run_timed travelodge "$BUILD" travelodge search --json London &
pid_travelodge=$!
run_timed hilton "$BUILD" hilton search --json London &
pid_hilton=$!
[ -n "$pid_duffel" ] && wait "$pid_duffel" || true
wait "$pid_ryanair" "$pid_vueling" "$pid_travelodge" "$pid_hilton"
wall_ms=$(($(ms_now) - wall_t0))
META=$(python3 -c "import json; print(json.dumps({'route':'${FROM}->London','ryanair_route':'${FROM}->STN','vueling_target':'${FROM}->LGW','depart':'${DEPART}','check_in':'${CHECK_IN}','check_out':'${CHECK_OUT}'}))")
SPECS=()
while IFS= read -r line; do SPECS+=("$line"); done < "$TMP/specs.txt"
python3 "$MERGE" --out "$OUT" --wall-ms "$wall_ms" --meta "$META" "${SPECS[@]}"
