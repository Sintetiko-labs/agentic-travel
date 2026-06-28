#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
source "$ROOT/scripts/mac-cli-common.sh"
[ $# -ge 1 ] || { echo usage >&2; exit 2; }
slug="$1"; shift
bin="$(mac_cli_cached_bin "$slug")"
[ -x "$bin" ] || { echo "missing $bin — run mac-build-all.sh" >&2; exit 1; }
exec "$bin" "$@"
