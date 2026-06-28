#!/usr/bin/env bash
# Focused Mac smoke: melia + iberostar only.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
LOG="$ROOT/scripts/smoke-melia-iberostar.log"
: > "$LOG"

log() { echo "[$(date '+%H:%M:%S')] $*" | tee -a "$LOG"; }

run_one() {
  local cli="$1"
  local search_args="$2"
  local dir="$ROOT/${cli}-cli"
  local bin="$dir/$cli"
  log "========== $cli =========="
  (cd "$dir" && go build -o "$cli" "./cmd/$cli") >>"$LOG" 2>&1
  codesign -s - --force "$bin" >/dev/null 2>&1 || true
  xattr -cr "$bin" 2>/dev/null || true

  log "doctor before"
  "$bin" session doctor --json 2>&1 | tee -a "$LOG" || true

  log "session chrome --replace --wait --timeout 3m"
  if "$bin" session chrome --replace --wait --timeout 3m >>"$LOG" 2>&1; then
    log "session chrome OK"
  else
    log "session chrome non-zero (partial cookies may be saved)"
  fi

  log "doctor after"
  "$bin" session doctor --json 2>&1 | tee -a "$LOG" || true

  log "search $search_args"
  set +e
  out=$("$bin" $search_args 2>&1)
  code=$?
  set -e
  echo "$out" | head -c 800 | tee -a "$LOG"
  log "search exit=$code"

  pkill -f "${HOME}/.${cli}/chrome-profile" 2>/dev/null || true
  sleep 2
}

log "smoke melia + iberostar — residential IP, headed Chrome"
run_one melia 'search --json Madrid --limit 3'
run_one iberostar 'search --json Madrid --limit 3'
log "done — log: $LOG"
