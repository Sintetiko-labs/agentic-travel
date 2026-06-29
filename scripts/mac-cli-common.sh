#!/usr/bin/env bash
_mac_cli_cache_root="${HOME}/.cache/agentic-travel"
AGENTIC_TRAVEL_BIN_CACHE="${AGENTIC_TRAVEL_BIN_CACHE:-${_mac_cli_cache_root}/bin}"
AGENTIC_TRAVEL_LOCK_DIR="${AGENTIC_TRAVEL_LOCK_DIR:-${_mac_cli_cache_root}/locks}"
LIVE_CLI_SLUGS=(ryanair vueling volotea binter barcelo riu catalonia h10 palladium lopesan princess eurostars hotusa vincci silken sercotel travelodge hilton globales grupotel hipotels senator medplaya zenit abba porthotels ona belive evenia ilunion petitpalace paradores roommate onlyyou pinero melia nh iberostar)
mac_cli_cached_bin() { printf '%s/%s\n' "$AGENTIC_TRAVEL_BIN_CACHE" "$1"; }
mac_cli_sign() {
  local bin="$1"
  command -v xattr >/dev/null 2>&1 && xattr -cr "$bin" 2>/dev/null || true
  command -v codesign >/dev/null 2>&1 && codesign -s - -f "$bin" 2>/dev/null || true
}
mac_cli_needs_rebuild() {
  local dir="$1" bin="$2"
  [ ! -x "$bin" ] && return 0
  [ "$dir/go.mod" -nt "$bin" ] && return 0
  find "$dir" -name '*.go' -newer "$bin" -print -quit 2>/dev/null | grep -q . && return 0
  return 1
}
mac_cli_acquire_lock() {
  local slug="$1" lockdir="$AGENTIC_TRAVEL_LOCK_DIR/${slug}.lockdir" attempts=0
  while ! mkdir "$lockdir" 2>/dev/null; do
    sleep 0.2; attempts=$((attempts+1))
    [ "$attempts" -ge 3000 ] && { echo "timeout lock: $slug" >&2; return 1; }
  done
  printf '%s\n' "$lockdir"
}
mac_cli_release_lock() { rmdir "$1" 2>/dev/null || true; }
mac_cli_build_cached() {
  local slug="$1" root="$2" dir="$root/${slug}-cli" bin lockdir
  bin="$(mac_cli_cached_bin "$slug")"
  [ -d "$dir" ] || { echo "unknown slug: $slug" >&2; return 1; }
  mkdir -p "$AGENTIC_TRAVEL_BIN_CACHE" "$AGENTIC_TRAVEL_LOCK_DIR"
  mac_cli_needs_rebuild "$dir" "$bin" || return 0
  lockdir="$(mac_cli_acquire_lock "$slug")" || return 1
  if mac_cli_needs_rebuild "$dir" "$bin"; then
    echo "build $slug → $bin" >&2
    (cd "$dir" && go build -o "$bin" "./cmd/$slug") || { mac_cli_release_lock "$lockdir"; return 1; }
    mac_cli_sign "$bin"
  fi
  mac_cli_release_lock "$lockdir"
}
