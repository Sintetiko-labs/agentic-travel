# Mac live smoke — Spanish hotel CLIs

**Date:** 2026-06-28  
**Host:** macOS (darwin 25.3.0), `/Users/fbelchi/github/agentic-travel`  
**Branch:** `loop-7/smoke-integration`  
**Runner:** `scripts/smoke-mac-hotels-es.py` (also `scripts/qa-hotels-es.py` for exhaustive QA)

## Procedure

Per CLI (`barcelo`, `riu`, `catalonia`, `h10`, `palladium`, `lopesan`, `princess`, `eurostars`, `vincci`, `sercotel`, `silken`):

1. `go build -o /tmp/{name} ./cmd/{name}` from `{name}-cli/`
2. `search --json Madrid --limit 3` (`silken`: `Barcelona`)
3. `read --json` on first hotel `id` or `hotel_url` when search returns hotels
4. `session doctor --json`

**PASS:** search JSON with `hotels[].name` (not `search not yet implemented`).  
**BLOCKED:** command timeout or host overload (no definitive pass/fail).  
**FAIL:** build error, stub search, empty results, or doctor failure.

## Results summary

| CLI | Build | Search | Read | Doctor | Status | Notes |
|-----|-------|--------|------|--------|--------|-------|
| barcelo | PASS | — | — | — | **FAIL** | Akamai 403 on `/es/hoteles` without session; needs `barcelo session chrome` |
| riu | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s (utls transport); curl shows ng-state OK |
| catalonia | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl 200 + hotel links OK |
| h10 | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl `/es/hoteles/madrid` 200 OK |
| palladium | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl `data-hotel-name` cards OK |
| lopesan | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl listing OK (Canary-focused) |
| princess | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl Tenerife/Madrid headings OK |
| eurostars | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl embedded JSON OK |
| vincci | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl `/es/hoteles/` links OK |
| sercotel | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl RSC hotel JSON OK |
| silken | PASS | — | — | — | **BLOCKED** | Live CLI hung >120s; curl `data-hotel` cards OK (use Barcelona) |

### Counts

| Metric | Count |
|--------|------:|
| **PASS** | 0 |
| **FAIL** | 1 (barcelo — Akamai) |
| **BLOCKED** | 10 (live CLI timeout during run) |

> **Note:** Earlier same-day exhaustive QA (`docs/QA_HOTELS_ES.md`, `scripts/qa-hotels-es.py`) reported **8 PASS / 3 WARN** when transport was healthy. This Mac run hit **utls TLS hangs** plus concurrent `verify-clis.sh` load; see fixes below.

## Root cause

1. **Transport:** `travelkit/transport` Chrome/uTLS dial could hang past `http.Client.Timeout` on macOS; no stdlib fallback.
2. **Host load:** Multiple `verify-clis.sh` / `go build` workers saturated the machine during smoke.
3. **Barceló:** Site returns HTTP 403 (Akamai) without headed Chrome session — expected, not a stub.

## Fixes (branches)

| Branch | Change |
|--------|--------|
| `loop-7/qa-fix-travelkit` | uTLS ctx deadline, `tcp4` dial, stdlib HTTP fallback in `FetchHTML`/`GetRaw`, `{PREFIX}_STD_HTTP=1` escape hatch |
| `loop-7/qa-fix-barcelo` | Document/session path only — Akamai requires `barcelo session chrome` on fresh Mac |

## Re-run

```bash
# Kill competing verify workers first
pkill -f verify-clis.sh || true

# Quick smoke (search/read/doctor)
python3 scripts/smoke-mac-hotels-es.py /tmp/smoke-hotels-es-results.json

# Exhaustive multi-city QA
python3 scripts/qa-hotels-es.py
```

After merging `loop-7/qa-fix-travelkit`, rebuild all CLIs before re-testing.
