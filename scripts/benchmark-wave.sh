#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
export AGENTIC_TRAVEL_ROOT="$ROOT"
SHELL_OUT="$ROOT/wave-result-shell.json"
GO_OUT="$ROOT/wave-result-go.json"
export WAVE_OUT="$SHELL_OUT"
echo "== shell wave ==" >&2
"$ROOT/scripts/wave-search-madrid-london.sh" >&2
export WAVE_OUT="$GO_OUT"
echo "== go wave ==" >&2
( cd "$ROOT/orchestrator" && AGENTIC_TRAVEL_ROOT="$ROOT" go run . )
python3 -c "import json; from pathlib import Path; s=json.loads(Path('${SHELL_OUT}').read_text()); g=json.loads(Path('${GO_OUT}').read_text()); print(json.dumps({'shell_wall_ms':s.get('wall_ms'),'go_wall_ms':g.get('wall_ms'),'shell_flights_total':s.get('flights_total'),'go_flights_total':g.get('flights_total'),'shell_hotels_total':s.get('hotels_total'),'go_hotels_total':g.get('hotels_total')}, indent=2))"
