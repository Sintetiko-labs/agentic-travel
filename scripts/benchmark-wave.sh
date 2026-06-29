#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
/usr/bin/time -p "$ROOT/scripts/wave-search-madrid-london.sh" >/dev/null
