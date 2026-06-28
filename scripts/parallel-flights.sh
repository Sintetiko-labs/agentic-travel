#!/usr/bin/env bash
# Agent entrypoint → parallel-search orchestrator (Go fan-out, 30s/source).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
exec "$ROOT/scripts/parallel-search/parallel-flights.sh" "$@"
