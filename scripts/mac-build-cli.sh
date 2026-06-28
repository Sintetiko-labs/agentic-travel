#!/usr/bin/env bash
# Build a single CLI and run a subcommand (used by docs and agents).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-/tmp/agentic-travel-bins}"
if [ $# -lt 1 ]; then echo "usage: mac-build-cli.sh <slug> [args…]" >&2; exit 2; fi
slug="$1"; shift
dir="$ROOT/${slug}-cli"; bin="$BIN_DIR/$slug"
if [ ! -d "$dir" ]; then echo "unknown slug: $slug" >&2; exit 1; fi
mkdir -p "$BIN_DIR"
if [ ! -x "$bin" ] || [ "$dir/go.mod" -nt "$bin" ]; then
  echo "build $slug → $bin" >&2
  (cd "$dir" && go build -o "$bin" "./cmd/$slug")
  command -v codesign >/dev/null && codesign -s - -f "$bin" 2>/dev/null || true
fi
exec "$bin" "$@"
