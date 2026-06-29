#!/usr/bin/env bash
# Shared helpers for parallel wave search (MCP + CLI fan-out).
# Mac residential only — run from Terminal.app, not CI/datacenter.
set -euo pipefail

ROOT="${ROOT:-$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)}"

WAVE_TIMEOUT="${WAVE_TIMEOUT:-30}"
WAVE_TMP=""
WAVE_WALL_START=0
WAVE_MANIFEST_SOURCES="[]"
WAVE_MCP_FALLBACK="[]"
WAVE_JOB_PIDS=()
WAVE_JOB_IDS=()
WAVE_JOB_STARTS=()

wave_now_ms() {
  if command -v python3 >/dev/null 2>&1; then
    python3 -c 'import time; print(int(time.time()*1000))'
  else
    echo $(( $(date +%s) * 1000 ))
  fi
}

wave_init_tmp() {
  WAVE_TMP="$(mktemp -d "${TMPDIR:-/tmp}/agentic-wave.XXXXXX")"
  WAVE_MANIFEST_SOURCES="[]"
  WAVE_MCP_FALLBACK="[]"
  WAVE_JOB_PIDS=()
  WAVE_JOB_IDS=()
  WAVE_JOB_STARTS=()
}

wave_cleanup_tmp() {
  [ -n "$WAVE_TMP" ] && [ -d "$WAVE_TMP" ] && rm -rf "$WAVE_TMP"
}

wave_run_with_timeout() {
  local timeout_sec="$1"
  shift
  if command -v gtimeout >/dev/null 2>&1; then
    gtimeout "$timeout_sec" "$@"
  elif command -v timeout >/dev/null 2>&1; then
    timeout "$timeout_sec" "$@"
  else
    perl -e 'alarm shift; exec @ARGV or die "exec failed: $!\n"' "$timeout_sec" "$@"
  fi
}

wave_append_source_json() {
  local entry="$1"
  WAVE_MANIFEST_SOURCES="$(node -e '
    const arr = JSON.parse(process.argv[1]);
    arr.push(JSON.parse(process.argv[2]));
    console.log(JSON.stringify(arr));
  ' "$WAVE_MANIFEST_SOURCES" "$entry")"
}

wave_append_fallback_json() {
  local entry="$1"
  WAVE_MCP_FALLBACK="$(node -e '
    const arr = JSON.parse(process.argv[1]);
    arr.push(JSON.parse(process.argv[2]));
    console.log(JSON.stringify(arr));
  ' "$WAVE_MCP_FALLBACK" "$entry")"
}

wave_register_skipped() {
  local id="$1" reason="$2"
  wave_append_source_json "$(node -e "console.log(JSON.stringify({id:process.argv[1],status:'skipped',duration_ms:0,error:process.argv[2]}))" "$id" "$reason")"
}

wave_run_bg() {
  local id="$1"
  shift
  local start_ms outfile errfile
  start_ms="$(wave_now_ms)"
  outfile="$WAVE_TMP/${id}.json"
  errfile="$WAVE_TMP/${id}.err"

  (
    if wave_run_with_timeout "$WAVE_TIMEOUT" "$@" >"$outfile" 2>"$errfile"; then
      echo ok >"$WAVE_TMP/${id}.status"
    else
      local ec=$?
      if [ "$ec" -eq 124 ] || [ "$ec" -eq 142 ]; then
        echo timed_out >"$WAVE_TMP/${id}.status"
      else
        echo error >"$WAVE_TMP/${id}.status"
      fi
    fi
  ) &

  WAVE_JOB_PIDS+=("$!")
  WAVE_JOB_IDS+=("$id")
  WAVE_JOB_STARTS+=("$start_ms")
}

wave_wait_all() {
  local i id start_ms end_ms status dur err file entry
  for i in "${!WAVE_JOB_PIDS[@]}"; do
    wait "${WAVE_JOB_PIDS[$i]}" || true
    id="${WAVE_JOB_IDS[$i]}"
    start_ms="${WAVE_JOB_STARTS[$i]}"
    end_ms="$(wave_now_ms)"
    dur=$(( end_ms - start_ms ))
    status="$(cat "$WAVE_TMP/${id}.status" 2>/dev/null || echo error)"
    err=""
    [ -f "$WAVE_TMP/${id}.err" ] && err="$(head -c 200 "$WAVE_TMP/${id}.err" | tr '\n' ' ')"
    file=""
    [ -f "$WAVE_TMP/${id}.json" ] && [ -s "$WAVE_TMP/${id}.json" ] && file="$WAVE_TMP/${id}.json"
    if [ "$status" = "ok" ] && [ -z "$file" ]; then
      status="error"
      err="${err:-empty output}"
    fi
    entry="$(node -e '
      const o = { id: process.argv[1], status: process.argv[2], duration_ms: Number(process.argv[3]) };
      if (process.argv[4]) o.file = process.argv[4];
      if (process.argv[5]) o.error = process.argv[5];
      console.log(JSON.stringify(o));
    ' "$id" "$status" "$dur" "$file" "$err")"
    wave_append_source_json "$entry"
  done
}

wave_build_manifest() {
  local query_json="$1"
  local wall_end wall_ms
  wall_end="$(wave_now_ms)"
  wall_ms=$(( wall_end - WAVE_WALL_START ))
  node -e '
    const manifest = {
      query: JSON.parse(process.argv[1]),
      sources: JSON.parse(process.argv[2]),
      wall_clock_ms: Number(process.argv[4]),
    };
    const fb = JSON.parse(process.argv[3]);
    if (fb.length) manifest.mcp_agent_fallback = fb;
    console.log(JSON.stringify(manifest));
  ' "$query_json" "$WAVE_MANIFEST_SOURCES" "$WAVE_MCP_FALLBACK" "$wall_ms"
}

wave_mcp_http_or_fallback() {
  local id="$1" url="$2" tool="$3" args_json="$4" fallback_server="$5"
  local start_ms outfile errfile
  start_ms="$(wave_now_ms)"
  outfile="$WAVE_TMP/${id}.json"
  errfile="$WAVE_TMP/${id}.err"

  if [ ! -d "${ROOT:-}/mcp/node_modules/@modelcontextprotocol/sdk" ]; then
    wave_append_fallback_json "$(node -e "console.log(JSON.stringify({server:process.argv[1],tool:process.argv[2],args:JSON.parse(process.argv[3]),note:'MCP SDK missing — run (cd mcp && npm ci); use CallMcpTool from agent'}))" "$fallback_server" "$tool" "$args_json")"
    wave_append_source_json "$(node -e "console.log(JSON.stringify({id:process.argv[1],status:'skipped',duration_ms:0,error:'MCP SDK not installed'}))" "$id")"
    return
  fi

  (
    if wave_run_with_timeout "$WAVE_TIMEOUT" node "$ROOT/mcp/call-mcp-http.mjs" --url "$url" --tool "$tool" --args "$args_json" >"$outfile" 2>"$errfile"; then
      echo ok >"$WAVE_TMP/${id}.status"
    else
      local ec=$?
      if [ "$ec" -eq 124 ] || [ "$ec" -eq 142 ]; then
        echo timed_out >"$WAVE_TMP/${id}.status"
      else
        echo error >"$WAVE_TMP/${id}.status"
      fi
    fi
  ) &

  WAVE_JOB_PIDS+=("$!")
  WAVE_JOB_IDS+=("$id")
  WAVE_JOB_STARTS+=("$start_ms")

  # On error after wait, caller may add fallback — handled in wave_wait_mcp_with_fallback
}

wave_run_duffel_mcp() {
  local from="$1" to="$2" depart="$3"
  local id="duffel-mcp"

  if [ -z "${DUFFEL_ACCESS_TOKEN:-}" ]; then
    wave_register_skipped "$id" "DUFFEL_ACCESS_TOKEN not set"
    return
  fi

  if [ ! -f "$ROOT/mcp/vendor/duffel-mcp/dist/index.js" ]; then
    wave_register_skipped "$id" "Duffel MCP not installed — run ./mcp/install.sh"
    return
  fi

  wave_run_bg "$id" env MCP_FROM="$from" MCP_TO="$to" MCP_DEPART="$depart" \
    "$ROOT/scripts/mcp-travel-search.sh"
}

wave_run_kiwi_mcp() {
  local from="$1" to="$2" depart="$3" adults="${4:-1}"
  local args_json
  args_json="$(node -e "console.log(JSON.stringify({origin:process.argv[1],destination:process.argv[2],departureDate:process.argv[3],adults:Number(process.argv[4]),cabin:'economy'}))" "$from" "$to" "$depart" "$adults")"
  wave_mcp_http_or_fallback "kiwi-mcp" "${KIWI_MCP_URL:-https://mcp.kiwi.com}" "search-flight" "$args_json" "kiwi-com-flight-search"
}

wave_run_gondola_mcp() {
  local city="$1" check_in="$2" check_out="$3" guests="${4:-2}"
  local args_json
  args_json="$(node -e "console.log(JSON.stringify({location:process.argv[1],check_in:process.argv[2],check_out:process.argv[3],guests:Number(process.argv[4])}))" "$city" "$check_in" "$check_out" "$guests")"
  wave_mcp_http_or_fallback "gondola-mcp" "${GONDOLA_MCP_URL:-https://mcp.gondola.ai/mcp}" "search_hotels" "$args_json" "gondola"
}

wave_add_mcp_fallbacks_for_errors() {
  local from="$1" to="$2" depart="$3" city="$4" check_in="$5" check_out="$6"
  local id status
  for id in kiwi-mcp gondola-mcp; do
    status="$(cat "$WAVE_TMP/${id}.status" 2>/dev/null || echo "")"
    if [ "$status" = "error" ] || [ "$status" = "timed_out" ]; then
      case "$id" in
        kiwi-mcp)
          wave_append_fallback_json "$(node -e "console.log(JSON.stringify({server:'kiwi-com-flight-search',tool:'search-flight',args:{origin:process.argv[1],destination:process.argv[2],departureDate:process.argv[3],adults:1,cabin:'economy'},note:'HTTP MCP failed — use CallMcpTool from Cursor agent'}))" "$from" "$to" "$depart")"
          ;;
        gondola-mcp)
          wave_append_fallback_json "$(node -e "console.log(JSON.stringify({server:'gondola',tool:'search_hotels',args:{location:process.argv[1],check_in:process.argv[2],check_out:process.argv[3],guests:2},note:'HTTP MCP failed — use CallMcpTool from Cursor agent'}))" "$city" "$check_in" "$check_out")"
          ;;
      esac
    fi
  done
}

wave_default_checkout() {
  local depart="$1"
  node -e '
    const d = new Date(process.argv[1] + "T12:00:00Z");
    d.setUTCDate(d.getUTCDate() + 3);
    console.log(d.toISOString().slice(0, 10));
  ' "$depart"
}
