# QA: Partial CLIs (loop-6)

Validated **2026-06-28** on macOS without headed Chrome. Session-dependent live search requires manual `session chrome --wait` steps documented below.

## Doctor results

| CLI | Build | `--help` | Doctor status | Probe HTTP | Next step hint | Search without session |
|-----|-------|----------|---------------|------------|----------------|------------------------|
| **melia** | OK | OK | `blocked` | 404 | `melia session chrome --wait --timeout 3m` | Clean error: `akamai blocked — MELIA_COOKIE required` |
| **nh** | OK | OK | `blocked` | 403 | `nh session chrome --wait --timeout 3m` | Clean error: `akamai blocked — NH_COOKIE required` |
| **iberostar** | OK | OK | `blocked`* | 404 | `iberostar session chrome --wait --timeout 3m` | Clean error: `akamai blocked — IBEROSTAR_COOKIE required` |
| **hotusa** | OK | OK | `missing_session` | 200† | `hotusa session chrome --wait --timeout 3m` | Clean error: `hotusa blocked request — run: hotusa session chrome --wait (HTTP 400)` |
| **easyjet** | OK | OK | `blocked` | 403 | `easyjet session chrome --wait --timeout 3m` | Clean error: `akamai blocked — EASYJET_COOKIE required` |
| **aireuropa** | OK | OK | `blocked` | 403 | `aireuropa session chrome --wait --timeout 3m` | Partial stub: Amadeus booking redirect URL (no fares/times without session) |
| **iberiaexpress** | OK | OK | `blocked` | 200‡ | `iberiaexpress session chrome --wait --timeout 3m` | Clean error: `incapsula blocked — IBERIAEXPRESS_COOKIE required` |
| **marriott** | OK | OK | `blocked` | 403 | `marriott session chrome --wait --timeout 3m` | Clean error: `akamai blocked — MARRIOTT_COOKIE required` |

\* With stale/incomplete cookies on disk, doctor may report `api_error` until cookies are cleared or refreshed.

† Hotusa booking host uses cert `*.booking-channel.com`; doctor probes corporate site `grupohotusa.com` (TLS-safe health check).

‡ HTTP 200 Incapsula JS challenge body — doctor correctly classifies as WAF challenge.

## Validation commands

```bash
# Per CLI (example: melia)
cd melia-cli && go build -o melia ./cmd/melia
./melia --help
./melia session doctor --json
./melia search Madrid          # expect clean error without session
```

Airline search probes:

```bash
easyjet search --from MAD --to PMI --depart 2026-07-15
aireuropa search --from MAD --to BCN --depart 2026-07-15
iberiaexpress search --from MAD --to BCN --depart 2026-07-15
```

Hotel search probes:

```bash
melia search Madrid
nh search Madrid
iberostar search Madrid
hotusa search Madrid
marriott search London
```

## Manual session steps (headed Chrome required for live data)

All partial CLIs persist cookies to `~/.{slug}/cookies.json`.

```bash
# 1. Launch headed Chrome (or use --replace to auto-launch)
{slug} session chrome --wait --timeout 3m

# 2. Browse the brand site until WAF cookies settle (_abck+bm_sz, cf_clearance, or Incapsula pair)

# 3. Verify
{slug} session doctor --json   # expect status: ok (or blocked if cookies stale)

# 4. Live search
{slug} search ...
```

Brand-specific start URLs used by `session chrome`:

| CLI | Start URL |
|-----|-----------|
| melia | `https://www.melia.com/es` |
| nh | `https://www.nh-hotels.com/es` |
| iberostar | `https://www.iberostar.com/es` |
| hotusa | `https://www.hotusa.com/es` (TLS SNI: `booking-channel.com`) |
| easyjet | `https://www.easyjet.com/es` |
| aireuropa | `https://www.aireuropa.com/es/es` |
| iberiaexpress | `https://www.iberiaexpress.com/es` |
| marriott | `https://www.marriott.com/` |

## Code fixes (loop-6/qa-fix-*)

| Branch | Fix |
|--------|-----|
| `loop-6/qa-fix-melia` | Map BFF 404/Next.js shell to session-required; doctor skips stale cookies on probe |
| `loop-6/qa-fix-iberostar` | GraphQL 404 → session hint |
| `loop-6/qa-fix-hotusa` | Doctor probe → `grupohotusa.com` (avoids TLS mismatch on `www.hotusa.com`) |
| `loop-6/qa-fix-travelkit` | `IsAppNotFoundWithoutSession`, doctor `probeCookie` only when WAF-ready |

Shared travelkit changes apply to all CLIs using `session doctor`.

### Client notes

- **melia** — BFF POST `/services/search/hotels/v2/search`; Akamai session required for JSON results.
- **nh** — GET `/nh/es/api/v1/hotels/search`; Akamai on API, clean block message without session.
- **iberostar** — GraphQL POST `/api/graphql`; 404 without session now surfaces as block hint.
- **hotusa** — Custom TLS to `booking-channel.com` via `internal/client/fetch.go`; HTML link scrape; session for live inventory.
- **easyjet** — GET `/ejavailability/api/v5/availability/query`; Akamai 403 without session.
- **aireuropa** — dapi `channel-home` Amadeus redirect + `flightinfo` fallback; booking redirect works without session but fares need chrome.
- **iberiaexpress** — GET `/api/availability/v1/flights`; Incapsula; bootstrap fetch of `/es/` first.
- **marriott** — HTML scrape `/search/findHotels.mi`; Akamai residential/session required.

## Pass criteria (this QA run)

- [x] All eight CLIs build and print `--help`
- [x] `session doctor --json` returns structured status + probe HTTP + next_step
- [x] `search` without session returns a clear error (no panic); aireuropa returns redirect stub only
- [x] Obvious client bugs fixed without headed Chrome (doctor probe, 404/session mapping, hotusa TLS probe)
- [ ] Live search with fresh session — **manual** (headed Chrome)
