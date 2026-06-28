#!/usr/bin/env bash
# Clone and build bokangsibolla/duffel-mcp into mcp/vendor/duffel-mcp
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
VENDOR="$ROOT/mcp/vendor/duffel-mcp"
REPO="${DUFFEL_MCP_REPO:-https://github.com/bokangsibolla/duffel-mcp.git}"
REF="${DUFFEL_MCP_REF:-main}"

if ! command -v node >/dev/null 2>&1; then
  echo "Node.js 20+ required (https://nodejs.org)" >&2
  exit 1
fi

node_major="$(node -p "process.versions.node.split('.')[0]")"
if [ "$node_major" -lt 20 ]; then
  echo "Node.js 20+ required (found $(node -v))" >&2
  exit 1
fi

if [ -d "$VENDOR/.git" ]; then
  echo "Updating $VENDOR ..."
  git -C "$VENDOR" fetch origin "$REF"
  git -C "$VENDOR" checkout "$REF"
  git -C "$VENDOR" pull --ff-only origin "$REF" 2>/dev/null || true
else
  echo "Cloning $REPO → $VENDOR ..."
  mkdir -p "$(dirname "$VENDOR")"
  git clone --depth 1 --branch "$REF" "$REPO" "$VENDOR"
fi

echo "Building duffel-mcp ..."
(cd "$VENDOR" && npm ci && npm run build)

echo "Installing MCP client deps ..."
(cd "$ROOT/mcp" && npm ci)

echo "Done. Server entry: $VENDOR/dist/index.js"
echo "Set DUFFEL_ACCESS_TOKEN (duffel_test_… from https://duffel.com) then run:"
echo "  ./scripts/mcp-travel-search.sh"
