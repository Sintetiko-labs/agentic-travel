#!/usr/bin/env bash
# Parallel Madrid → London wave: Duffel MCP + Kiwi + Gondola + airline/hotel CLIs.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-$HOME/.cache/agentic-travel/bin}"
BUILD="$ROOT/scripts/mac-build-cli.sh"
MERGE="$ROOT/scripts/wave-merge.py"
HTTP_MCP="$ROOT/mcp/call-http-mcp-tool.mjs"

FROM="${WAVE_FROM:-MAD}"
TO="${WAVE_TO:-STN}"
DEPART="${WAVE_DEPART:-2026-07-15}"
CHECK_IN="${WAVE_CHECK_IN:-2026-07-15}"
CHECK_OUT="${WAVE_CHECK_OUT:-2026-07-18}"
CITY="${WAVE_CITY:-London}"
GUESTS="${WAVE_GUESTS:-2}"
PER_SOURCE_TIMEOUT="${WAVE_SOURCE_TIMEOUT:-14}"

OUT="${WAVE_OUT:-$ROOT/wave-result.json}"
TMP="${WAVE_TMP:-$(mktemp -d "${TMPDIR:-/tmp}/wave-search.XXXXXX")}"
trap 'rm -rf "$TMP"' EXIT
mkdir -p "$BIN_DIR"
export AGENTIC_TRAVEL_BINS="$BIN_DIR"

cli() {
  local slug="$1"
  local cached="$BIN_DIR/$slug"
  if [[ -x "$cached" ]]; then echo "$cached"; else echo "__BUILD__$slug"; fi
}

now_ms() { python3 -c 'import time; print(int(time.time()*1000))'; }

run_job() {
  local id="$1"; shift
  local start end ms rc=0
  start="$(now_ms)"
  if command -v gtimeout >/dev/null 2>&1; then
    gtimeout "$PER_SOURCE_TIMEOUT" "$@" >"$TMP/$id.json" 2>"$TMP/$id.err" || rc=$?
  elif command -v timeout >/dev/null 2>&1; then
    timeout "$PER_SOURCE_TIMEOUT" "$@" >"$TMP/$id.json" 2>"$TMP/$id.err" || rc=$?
  else
    "$@" >"$TMP/$id.json" 2>"$TMP/$id.err" || rc=$?
  fi
  end="$(now_ms)"; ms=$((end - start))
  local ok=false; [[ $rc -eq 0 ]] && ok=true
  python3 -c "import json; print(json.dumps({'id':'$id','ok':$ok,'ms':$ms,'exit_code':$rc,'skipped':False}))"
}

mark_skip() {
  local id="$1" reason="$2"
  python3 -c "import json; o={'skipped':True,'reason':$reason}; print(json.dumps(o))" >"$TMP/$id.json"
  python3 -c "import json; print(json.dumps({'id':'$id','ok':False,'ms':0,'skipped':True,'error':$reason}))"
}

run_cli_job() {
  local id="$1" slug="$2"; shift 2
  local spec; spec="$(cli "$slug")"
  if [[ "$spec" == __BUILD__* ]]; then
    slug="${spec#__BUILD__}"
    run_job "$id" "$BUILD" "$slug" "$@"
  else
    run_job "$id" "$spec" "$@"
  fi
}

wall_start="$(now_ms)"
pids=()

if [[ -n "${DUFFEL_ACCESS_TOKEN:-}" ]]; then
  ( meta=$(run_job duffel node "$ROOT/mcp/call-search-flights.mjs" --from "$FROM" --to "$TO" --depart "$DEPART"); echo "$meta" >"$TMP/duffel.meta.json" ) & pids+=($!)
else
  echo "$(mark_skip duffel "$(python3 -c 'import json;print(json.dumps("DUFFEL_ACCESS_TOKEN unset"))')")" >"$TMP/duffel.meta.json"
fi

if [[ -d "$ROOT/mcp/node_modules/@modelcontextprotocol/sdk" ]]; then
  kiwi_args=$(DEPART="$DEPART" python3 -c "import json,os; d=os.environ['DEPART']; dd=f'{d[8:10]}/{d[5:7]}/{d[0:4]}'; print(json.dumps({'flyFrom':'$FROM','flyTo':'$TO','departureDate':dd,'passengers':{'adults':1},'cabinClass':'M'}))")
  ( meta=$(run_job kiwi node "$HTTP_MCP" --url https://mcp.kiwi.com --tool search-flight --args "$kiwi_args"); echo "$meta" >"$TMP/kiwi.meta.json" ) & pids+=($!)
  gondola_args=$(python3 -c "import json;print(json.dumps({'location':'$CITY','checkin':'$CHECK_IN','checkout':'$CHECK_OUT','num_adults':int('$GUESTS')}))")
  ( meta=$(run_job gondola node "$HTTP_MCP" --url https://mcp.gondola.ai/mcp --tool search_hotels --args "$gondola_args"); echo "$meta" >"$TMP/gondola.meta.json" ) & pids+=($!)
else
  for id in kiwi gondola; do
    echo "$(mark_skip "$id" "$(python3 -c 'import json;print(json.dumps("(cd mcp && npm ci)"))')")" >"$TMP/${id}.meta.json"
  done
fi

( meta=$(run_cli_job ryanair ryanair search --json --from "$FROM" --to "$TO" --depart "$DEPART"); echo "$meta" >"$TMP/ryanair.meta.json" ) & pids+=($!)
( meta=$(run_cli_job vueling vueling search --json --from "$FROM" --to "$TO" --depart "$DEPART"); echo "$meta" >"$TMP/vueling.meta.json" ) & pids+=($!)
( meta=$(run_cli_job travelodge travelodge search --json "$CITY"); echo "$meta" >"$TMP/travelodge.meta.json" ) & pids+=($!)
( meta=$(run_cli_job hilton hilton search --json "$CITY"); echo "$meta" >"$TMP/hilton.meta.json" ) & pids+=($!)

for pid in "${pids[@]}"; do wait "$pid" || true; done
wall_ms=$(( $(now_ms) - wall_start ))
query=$(python3 -c "import json;print(json.dumps({'from':'$FROM','to':'$TO','depart':'$DEPART','city':'$CITY','check_in':'$CHECK_IN','check_out':'$CHECK_OUT'}))")
"$MERGE" --meta-dir "$TMP" --out "$OUT" --wall-ms "$wall_ms" --query "$query" >&2
cat "$OUT"
