#!/usr/bin/env bash
# Cadiz hotel parallel search — all live hotel CLIs, July week 2026-07-05 → 2026-07-12.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
# shellcheck source=mac-cli-common.sh
source "$ROOT/scripts/mac-cli-common.sh"
# shellcheck source=wave-search-common.sh
source "$ROOT/scripts/wave-search-common.sh"

CITY="${CADIZ_CITY:-Cadiz}"
CHECK_IN="${CADIZ_CHECKIN:-2026-07-05}"
CHECK_OUT="${CADIZ_CHECKOUT:-2026-07-12}"
LIMIT="${CADIZ_LIMIT:-10}"
OUT="${CADIZ_OUT:-$ROOT/cadiz-hotels-july.json}"
WAVE_TIMEOUT="${WAVE_TIMEOUT:-45}"

# All implemented hotel slugs (non-stub search.go)
HOTEL_SLUGS=(
  barcelo riu catalonia h10 palladium lopesan princess eurostars hotusa vincci silken sercotel
  travelodge hilton globales grupotel hipotels senator medplaya zenit abba porthotels ona belive
  evenia ilunion petitpalace paradores roommate onlyyou pinero melia nh iberostar
  accor ihg mamashelter 25hours citizenm hoxton wyndham hyatt bestwestern radisson
  easyhotel bbhotels numa sonder limehome marriott
  designhotels leonardo
)

echo "=== Phase 1: sequential builds ===" >&2
build_ok=0 build_fail=0
for slug in "${HOTEL_SLUGS[@]}"; do
  if mac_cli_build_cached "$slug" "$ROOT" 2>/dev/null; then
    build_ok=$((build_ok + 1))
  else
    build_fail=$((build_fail + 1))
    echo "build skip: $slug" >&2
  fi
done
echo "builds: $build_ok ok, $build_fail failed" >&2

echo "=== Phase 2: parallel search $CITY ($CHECK_IN → $CHECK_OUT) ===" >&2
wave_init_tmp
WAVE_WALL_START="$(wave_now_ms)"
trap wave_cleanup_tmp EXIT

chains_parallel=0
for slug in "${HOTEL_SLUGS[@]}"; do
  bin="$(mac_cli_cached_bin "$slug")"
  if [ ! -x "$bin" ]; then
    wave_register_skipped "$slug" "binary missing"
    continue
  fi
  chains_parallel=$((chains_parallel + 1))
  # search uses city; availability subcommand needs hotel-id (dates noted in meta)
  wave_run_bg "$slug" "$bin" search --json --limit "$LIMIT" "$CITY"
done

wave_wait_all
wall_ms=$(( $(wave_now_ms) - WAVE_WALL_START ))

manifest="$(wave_build_manifest "$(
  node -e 'console.log(JSON.stringify({city:process.argv[1],check_in:process.argv[2],check_out:process.argv[3],limit:+process.argv[4]}))' \
    "$CITY" "$CHECK_IN" "$CHECK_OUT" "$LIMIT"
)")"

merged="$(node "$ROOT/mcp/merge-wave-result.mjs" <<<"$manifest")"

# Enrich into cadiz-hotels-july.json
node -e '
const fs = require("fs");
const merged = JSON.parse(process.argv[1]);
const meta = {
  city: process.argv[2],
  check_in: process.argv[3],
  check_out: process.argv[4],
  timestamp: new Date().toISOString(),
  chains_parallel: +process.argv[5],
  wall_ms: +process.argv[6],
  note: "search subcommand returns directory listings; use availability --check-in for priced rooms",
};
const chains = [];
const empty_or_errors = [];
for (const src of merged.sources || []) {
  const slug = src.id;
  const dur = src.duration_ms || 0;
  if (src.status !== "ok" || !src.data) {
    empty_or_errors.push({ slug, error: src.error || src.status || "no data" });
    chains.push({ slug, hotels_found: 0, cheapest: null, sample_hotel: null, duration_ms: dur, status: src.status });
    continue;
  }
  const hotels = src.data.hotels || [];
  let cheapest = null;
  let cheapestHotel = null;
  for (const h of hotels) {
    const p = h.price ?? h.min_price ?? h.rate;
    const cur = h.currency || h.price_currency || "EUR";
    if (p != null && (cheapest == null || p < cheapest)) {
      cheapest = p;
      cheapestHotel = h;
    }
  }
  const sample = hotels[0];
  const sampleName = sample ? (sample.name || sample.hotel_name || "—") : "—";
  let cheapestStr = "—";
  if (cheapest != null && cheapestHotel) {
    const cur = cheapestHotel.currency || cheapestHotel.price_currency || "EUR";
    const nights = 7;
  const ppn = cheapestHotel.price_per_night ?? (cheapest / nights);
    cheapestStr = `${cheapest} ${cur}` + (ppn ? ` (~${ppn.toFixed ? ppn.toFixed(2) : ppn}/night)` : "");
  }
  if (hotels.length === 0) {
    empty_or_errors.push({ slug, error: "0 hotels" });
  }
  chains.push({
    slug,
    hotels_found: hotels.length,
    cheapest: cheapestStr,
    sample_hotel: sampleName,
    duration_ms: dur,
    status: "ok",
    sample: sample ? {
      name: sample.name,
      price: sample.price ?? sample.min_price,
      currency: sample.currency,
      booking_url: sample.booking_url || sample.url || sample.hotel_url,
    } : null,
  });
}
const out = { meta, chains, empty_or_errors, raw: merged };
fs.writeFileSync(process.argv[7], JSON.stringify(out, null, 2));
console.log(JSON.stringify({chains_parallel: meta.chains_parallel, with_hotels: chains.filter(c=>c.hotels_found>0).length, wall_ms: meta.wall_ms}));
' "$merged" "$CITY" "$CHECK_IN" "$CHECK_OUT" "$chains_parallel" "$wall_ms" "$OUT"

echo "wrote $OUT" >&2
cat "$OUT" | node -e "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); console.log('parallel:', d.meta.chains_parallel, 'with_hotels:', d.chains.filter(c=>c.hotels_found>0).length)"
