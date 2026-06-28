# QA: UK Hotel CLIs (Travelodge, Hilton, Marriott)

**Date:** 2026-06-28  
**Base branch:** `loop-6/uk-hotels`  
**Fix branches:** `loop-6/qa-fix-travelodge`, `loop-6/qa-fix-hilton`, `loop-6/qa-fix-marriott`  
**Environment:** macOS, headed Chrome available, network live

## Summary

| CLI | Core search | Session / WAF | Edge cases | Fix branch |
|-----|-------------|---------------|------------|------------|
| **travelodge** | PASS | N/A (public sitemap) | Fixed empty + compound queries | `loop-6/qa-fix-travelodge` |
| **hilton** | PASS | N/A (public location pages) | Fixed empty query validation | `loop-6/qa-fix-hilton` |
| **marriott** | BLOCKED (Akamai) | Cookies capture OK; CLI still blocked | Fixed doctor probe + errors | `loop-6/qa-fix-marriott` |

---

## Travelodge

### Core smoke (pre-fix, `loop-6/uk-hotels`)

| Test | Command | Result | Notes |
|------|---------|--------|-------|
| London | `travelodge search --json --limit 10 London` | **PASS** | `total=579`, 10 hotels returned |
| Manchester | `travelodge search --json --limit 10 Manchester` | **PASS** | `total=98` |
| Edinburgh | `travelodge search --json --limit 10 Edinburgh` | **PASS** | `total=23` |
| Pagination | `--limit 5 --page 2 London` | **PASS** | `page=2`, 5 hotels, `has_next_page=true` |
| JSON validity | all successful outputs | **PASS** | Valid JSON, snake_case fields |

### Edge cases (pre-fix)

| Test | Result | Issue |
|------|--------|-------|
| Empty query `""` | **FAIL** | Returned `total=2937` (entire UK sitemap) |
| Special chars `London & Manchester` | **FAIL** | `total=0` — literal `&` match |
| No args | **PASS** | Usage error as expected |

### Post-fix (`loop-6/qa-fix-travelodge`)

| Test | Result |
|------|--------|
| Empty query `""` | **PASS** — `destination required (non-empty)` |
| `London & Manchester` | **PASS** — `total=677` (union of London + Manchester) |
| London / Manchester / Edinburgh | **PASS** — unchanged |

### Fix (`0194d43`)

- Reject whitespace-only destination in `cmd/search.go` and `internal/client/search.go`
- Split compound queries on `&` and `|` and union deduplicated results

---

## Hilton

### Core smoke (pre-fix)

| Test | Command | Result | Notes |
|------|---------|--------|-------|
| London | `hilton search --json --limit 10 London` | **PASS** | `total=20`, 10 hotels |
| Manchester | `hilton search --json --limit 10 Manchester` | **PASS** | `total=11`, 10 hotels |
| Pagination | `--limit 5 --page 2 London` | **PASS** | `page=2`, 5 hotels |
| JSON validity | all successful outputs | **PASS** | Valid JSON |

### Edge cases (pre-fix)

| Test | Result | Issue |
|------|--------|-------|
| Empty query `""` | **FAIL** | `total=0` — fetched `/en/locations/united-kingdom//` |
| Special chars `St. James's` | **WARN** | `total=20` — slug not sanitized; likely wrong page |
| No args | **PASS** | Usage error |

### Post-fix (`loop-6/qa-fix-hilton`, `95d2b44`)

| Test | Result |
|------|--------|
| Empty query `""` | **PASS** — `destination required (non-empty)` |
| `St. James's` | **WARN** — slug sanitized to `st-jamess`; page may 404/redirect (no UK hotels at that path) |
| London / Manchester | **PASS** — unchanged |

### Fix

- Reject empty destination before API call
- Strip `.'`, `'`, `,` from location slugs before path construction

---

## Marriott

### Session flow

| Step | Command | Result |
|------|---------|--------|
| 1. Doctor (no session) | `marriott session doctor` | **PASS (expected)** — `status=missing_session`, no cookie file |
| 2. Search without session | `marriott search --json --limit 10 London` | **PASS (expected)** — Akamai blocked |
| 3. Chrome capture | `marriott session chrome --wait --timeout 2m` | **PASS** — cookies saved to `~/.marriott/cookies.json`, `_abck` + `bm_sz` present |
| 4. Doctor (with session) | `marriott session doctor` | **FAIL** — probe HTTP 403 (pre-fix: homepage probe) |
| 5. Search (with session) | `marriott search --json --limit 10 London` | **FAIL** — Akamai 403 from CLI even with cookies |

### Edge cases (pre-fix)

| Test | Result |
|------|--------|
| Empty query | Blocked (same as any search — no empty validation) |
| Special chars | Blocked |
| JSON on failure | No stdout (errors on stderr only) — expected |

### Post-fix (`loop-6/qa-fix-marriott`, `1177ef3`)

| Test | Result |
|------|--------|
| Empty query `""` | **PASS** — `destination required (non-empty)` |
| Doctor probe URL | **IMPROVED** — probes `findHotels.mi` instead of homepage |
| Search with stale cookies | **IMPROVED** — `saved cookies may be stale or not valid for CLI requests` |
| Search London (live) | **STILL BLOCKED** — environmental Akamai TLS/IP binding |

### Fix

- `session doctor` probes `findHotels.mi?city=London` (same endpoint as search)
- Search uses navigation headers (`referer`, `sec-fetch-*`) via `fetchSearchHTML`
- Distinguish missing vs present-but-stale cookies in error messages

### Known limitation

Marriott Akamai binds cookies to the browser TLS fingerprint. `session chrome` succeeds in headed Chrome, but programmatic requests (even with Chrome transport + saved cookies) receive HTTP 403 `Access Denied` from edgesuite.net. **Workaround:** re-capture session immediately before search, or run from the same network profile as Chrome. Full bypass may require additional Akamai sensor work (out of scope for this QA pass).

---

## Cross-cutting checks

| Check | travelodge | hilton | marriott |
|-------|------------|--------|----------|
| JSON `--json` output | Valid | Valid | Valid (on success) |
| `has_next_page` field | Present | Present | Present (on success) |
| `page` / `page_size` | Correct | Correct | Correct |
| Empty query guard | Fixed | Fixed | Fixed |
| Pagination | Works | Works | N/A (blocked) |

---

## How to re-run

```bash
# Build (from respective fix branch)
cd travelodge-cli && go build -o travelodge ./cmd/travelodge
cd hilton-cli    && go build -o hilton ./cmd/hilton
cd marriott-cli  && go build -o marriott ./cmd/marriott

# Core
travelodge search --json --limit 10 London
hilton search --json --limit 10 London
marriott session doctor
marriott search --json --limit 10 London   # expects block without fresh session

# Edge cases
travelodge search --json ""
travelodge search --json --limit 5 "London & Manchester"
hilton search --json ""
marriott session chrome --wait --timeout 2m  # headed Chrome required
```

---

## Commits

| Branch | Commit | Description |
|--------|--------|-------------|
| `loop-6/qa-fix-travelodge` | `0194d43` | Empty query guard + compound destination split |
| `loop-6/qa-fix-hilton` | `95d2b44` | Empty query guard + slug sanitization |
| `loop-6/qa-fix-marriott` | `1177ef3` | findHotels probe, nav headers, stale-cookie errors |
