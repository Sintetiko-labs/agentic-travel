# Loop status (parallel workstreams)

Last updated: 2026-06-29 (loop-7 parent APIs: accor, ihg, lufthansagroup)

## Main branch

| Field | Value |
|-------|-------|
| **SHA** | `c9c29da` (+ parent API hardening commits) |
| **Tip** | loop-7: live parent search for Accor (11 brands), IHG (9 brands), Lufthansa Group (7 brands) |
| **Remote** | `Sintetiko-labs/agentic-travel` |

## Loop 7 stub elimination (parent APIs)

| Metric | Value |
|--------|------:|
| **Scaffold stubs before** | 117 |
| **Scaffold stubs after wire** | 117 |
| **Parents implemented** | `accor`, `ihg`, `lufthansagroup` |
| **Brands covered (parents)** | 27 (Accor 11 + IHG 9 + LH Group 7) |
| **Child stubs wired** | 0 — `mamashelter` / `25hours` keep bespoke homepage parsers per `STUB_ELIMINATION.md` |

Parent APIs:

| Parent | API | Child delegates |
|--------|-----|-----------------|
| `accor` | JSON-LD + destination HTML + SSR fallback (`AccorDestinationPath` / `AccorSearchFallbackPath`) | — (mamashelter, 25hours: own search) |
| `ihg` | JSON-LD + IHG property JSON (`IHGSearchPath`) | — |
| `lufthansagroup` | LH `api-shop/lowestfares` + Eurowings `search.api.json` | — |

## Merge status

| Branch | → main | Notes |
|--------|--------|-------|
| `loop-5-iteration5` | **done** (ancestor of main) | Vueling FlightPrice, hotel batch loop 5, session doctor |
| `loop-6/uk-hotels` | **done** | travelodge, hilton, marriott (partial) |
| `loop-6/hotel-batch-es` | **done** | eurostars, hotusa, vincci, silken, sercotel live search |
| `loop-6/hotels-akamai` | **done** | melia, nh, iberostar unblock (partial / Akamai) |
| `loop-6/airlines-partial` | **done** | easyjet, aireuropa, iberiaexpress partial unblock |
| `loop-6/volotea-binter` | **done** | live APIs on main; merge record `5cfd2e0` |
| `loop-6/qa-*` | **done** (integrated) | hotusa, travelkit, melia, iberostar, volotea, binter, ryanair, travelodge, marriott, hilton, hotels-es — on `dead273` / `05d1485` ancestry |

## Loop 6 worker lanes (prep — not merged)

| Lane | Branch | Focus |
|------|--------|-------|
| A | `loop-6/hotels-akamai` | Akamai hotel partials |
| B | `loop-6/hotel-batch-es` | Spanish hotel HTML parsers |
| C | `loop-6/airlines-partial` | Akamai / Incapsula airline partials |
| D | `loop-6/uk-hotels` | UK Madrid→London scenario |
| E | `loop-6/volotea-binter` | Volotea + Binter live APIs |
| F | `loop-6/qa-integration` | Merge hygiene, verify, docs (this file) |
| QA | `loop-6/qa-*` | Smoke / fix passes before merge |

## Priority CLI implementation counts

Counts reflect **documented priority** CLIs in root README (not all 194 scaffolds).

| Status | Count | Slugs |
|--------|------:|-------|
| **live** | 35 | ryanair, vueling, volotea, binter, barcelo, riu, catalonia, h10, palladium, lopesan, princess, eurostars, hotusa, vincci, silken, sercotel, travelodge, hilton, globales, grupotel, hipotels, senator, medplaya, zenit, abba, porthotels, ona, belive, evenia, ilunion, petitpalace, paradores, roommate, onlyyou, pinero, melia, nh, iberostar |
| **partial** | 4 | marriott, easyjet, aireuropa, iberiaexpress |
| **stub** (priority-adjacent) | 0 | — |

**All CLIs (scaffold):** 194 hotel + airline modules — most non-priority slugs remain **stub** (`search not yet implemented`).

## Monorepo verify (`./scripts/verify-clis.sh`)

Build + `{slug} help` for every `*-cli` directory. Run on clean `main` after merges.

| Run | SHA | Pass | Fail | Notes |
|-----|-----|-----:|-----:|-------|
| Loop 6 QA subset (2026-06-28) | `05d1485` | 194 | 0 | hilton, marriott, ryanair, travelodge + docs |
| Loop 6 catch-up (2026-06-28) | `396e6bd` | 194 | 0 | volotea `flights:[]` + merge records (re-run `./scripts/verify-clis.sh` to confirm) |

Update this table after each integration pass.

## Next integration actions

1. ~~Merge `loop-6/volotea-binter`~~ — **done** on `396e6bd`.
2. ~~Merge `loop-6/qa-*` on main~~ — **done** (`05d1485` + merge commits).
3. Keep `./scripts/verify-clis.sh` at **194/194** after each integration pass.
4. Next: Akamai partials (melia, nh, iberostar, easyjet, …) toward more **live** slugs.
