#!/usr/bin/env bash
# Build airline + hotel CLIs into /tmp/agentic-travel-bins/ (ad-hoc signed).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
BIN_DIR="${AGENTIC_TRAVEL_BINS:-/tmp/agentic-travel-bins}"
ORCH_DIR="$ROOT/orchestrator"

DEFAULT_SLUGS=(
  ryanair vueling volotea aireuropa binter
  travelodge hilton barcelo marriott melia nh
)

SLUGS=("${@:-${DEFAULT_SLUGS[@]}}")
mkdir -p "$BIN_DIR"

build_slug() {
  local slug="$1"
  local dir="$ROOT/${slug}-cli"
  local out="$BIN_DIR/$slug"
  if [ ! -d "$dir" ]; then
    echo "skip $slug (no ${slug}-cli/)" >&2
    return 0
  fi
  echo "build $slug → $out" >&2
  (cd "$dir" && go mod tidy >/dev/null 2>&1 && go build -o "$out" "./cmd/$slug")
  command -v codesign >/dev/null && codesign -s - -f "$out" 2>/dev/null || true
}

for slug in "${SLUGS[@]}"; do
  build_slug "$slug"
done

echo "build orchestrator binaries" >&2
(cd "$ORCH_DIR" && go mod tidy && go build -o "$BIN_DIR/parallel-flights" ./cmd/parallel-flights)
(cd "$ORCH_DIR" && go build -o "$BIN_DIR/parallel-hotels" ./cmd/parallel-hotels)
command -v codesign >/dev/null && codesign -s - -f "$BIN_DIR/parallel-flights" "$BIN_DIR/parallel-hotels" 2>/dev/null || true

echo "bins ready: $BIN_DIR" >&2
