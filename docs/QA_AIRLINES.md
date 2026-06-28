# QA Smoke Tests — Live Airline CLIs

**Date:** 2026-06-28  
**Repo:** `agentic-travel`  
**Tester:** loop-6 QA automation  
**Environment:** macOS, Go 1.26+, residential network (live API calls)

## Scope

Smoke tests for airline CLIs marked **live** in README (loop 5 priority table) plus Volotea/Binter from iteration 6.

| CLI | README status | Implementation (main / `loop-6/volotea-binter`) |
|-----|---------------|--------------------------------------------------|
| `ryanair` | **live** | Full — farfnd + booking availability |
| `vueling` | **live** | Full — FlightPrice API |
| `volotea` | iteration 6 | Scaffold on `main`; **implemented** on `loop-6/volotea-binter` |
| `binter` | iteration 6 | Scaffold on `main`; **implemented** on `loop-6/volotea-binter` |

## Test matrix (per CLI)

1. `go build -o /tmp/{name} ./cmd/{name}`
2. `--help`, `search --help`, `session doctor --json`
3. Routes: MAD→BCN, MAD→STN, MAD→PMI, BCN→PMI
4. Dates: 2026-07-05, 2026-07-08, 2026-07-15
5. JSON schema: `flights` array (not `null`), `price`/`currency` on hits, no panic
6. Invalid input: `XXX→YYY`, past date `2020-01-01` — clean error, no crash

---

## Summary

| CLI | Build | Help / doctor | Search (live routes) | Invalid input | **Overall** |
|-----|-------|---------------|----------------------|---------------|-------------|
| **ryanair** | PASS | PASS | PASS after fix¹ | WARN² | **PASS** (with fix branch) |
| **vueling** | PASS | PASS | PASS³ | PASS | **PASS** |
| **volotea** | PASS | PASS | PASS⁴ / FAIL⁵ | PASS | **PASS** (scaffold) / **FAIL** (implemented, null JSON) |
| **binter** | PASS | PASS | PASS⁶ / FAIL⁵ | PASS | **PASS** (scaffold) / **FAIL** (implemented, null JSON) |

¹ See fix branch `loop-6/qa-fix-ryanair`  
² Bad airport / past date return exit 0 with empty JSON (no panic) — acceptable but not ideal  
³ Many routes returned HTTP 404 from API (`Not flights were found`); CLI surfaces clean stderr error (exit 1)  
⁴ On `main`: scaffold returns `search not yet implemented` (exit 1)  
⁵ On `loop-6/volotea-binter`: live API works but empty results emit `"flights": null`  
⁶ Binter MAD→PMI etc. hit GraphQL `booking_search`; empty availability returns JSON with `flights: null`

---

## ryanair

**Branch tested:** `loop-6/qa-fix-ryanair` (post-fix)  
**Result:** **PASS**

### Build & CLI surface

| Check | Result |
|-------|--------|
| `go build` | PASS |
| `--help` | PASS |
| `search --help` | PASS |
| `session doctor --json` | PASS (valid JSON, exit 0) |

### Search results (post-fix)

| Route | 2026-07-05 | 2026-07-08 | 2026-07-15 |
|-------|------------|------------|------------|
| MAD→BCN | `flights: []` (0) | `flights: []` (0) | `flights: []` (0) |
| MAD→STN | 1 flight | 1 flight | 1 flight |
| MAD→PMI | 1 flight | 1 flight | 1 flight |
| BCN→PMI | 1 flight | 1 flight | 1 flight |

### Bug found & fixed

**Issue:** Empty search results serialized `"flights": null` instead of `"flights": []` (Go nil slice → JSON null). Broke agent JSON schema validation on MAD→BCN (Ryanair does not fly this route; farfnd returns 0 fares).

**Fix:** `ryanair-cli/internal/client/search.go` (`paginateFlights`) and `availability.go` (`buildFarfndResult`) — initialize empty slice as `[]FlightHit{}`.

**Branch pushed:** `loop-6/qa-fix-ryanair`  
**Commit:** `Fix ryanair search JSON emitting flights:null on empty results.`

### Invalid input

| Input | Result |
|-------|--------|
| `XXX→YYY` | exit 0, empty JSON (WARN — no validation error) |
| `MAD→BCN 2020-01-01` | exit 0, empty JSON (WARN) |

No panics observed.

---

## vueling

**Branch tested:** `main` / `loop-6/qa-fix-ryanair` (vueling unchanged)  
**Result:** **PASS**

### Build & CLI surface

| Check | Result |
|-------|--------|
| `go build` | PASS |
| `--help` | PASS |
| `search --help` | PASS |
| `session doctor --json` | PASS |

### Search results

API variability observed across runs:

- **Run A:** MAD→PMI, BCN→PMI returned valid JSON with priced flights on some dates.
- **Run B:** All 12 route×date combos returned HTTP 404 `{"Message":"Not flights were found"}` — CLI exits 1 with clean stderr message.

Both behaviors are acceptable (no panic, no malformed JSON on error path).

### Invalid input

| Input | Result |
|-------|--------|
| `XXX→YYY` | exit 1, clean error |
| Past date | exit 1, clean error |

**Fixes:** None required.

---

## volotea

### On `main` (scaffold)

**Result:** **PASS** (scaffold expectations)

- All 12 searches return `search not yet implemented for Volotea` (exit 1)
- `session doctor --json` valid
- Invalid input: clean errors (exit 1)

### On `loop-6/volotea-binter` (implemented)

**Result:** **FAIL** (JSON schema)

| Route | Typical result |
|-------|----------------|
| MAD→BCN | exit 0, `"flights": null`, `source: flights/search` |
| MAD→PMI | exit 0 or 1 with partial JSON on error paths |
| BCN→PMI | Similar null-array issue when empty |

**Bug:** Same nil-slice JSON null pattern as ryanair pre-fix.  
**Fix prepared:** Ensure `Flights` is `[]FlightHit{}` when `flights[start:end]` is empty/nil.  
**Branch:** `loop-6/qa-fix-volotea` (local; push pending git lock during QA run)

---

## binter

### On `main` (scaffold)

**Result:** **PASS** (scaffold expectations)

- All searches return `search not yet implemented for Binter` (exit 1)
- Invalid input: clean errors

### On `loop-6/volotea-binter` (implemented)

**Result:** **FAIL** (JSON schema on empty results)

Example:

```json
{
  "query": "MAD-PMI 2026-07-05",
  "total": 0,
  "flights": null,
  "source": "booking_search"
}
```

MAD→PMI is not a Binter route (Canarias carrier); empty availability should return `"flights": []`.

**Fix needed:** `binter-cli/internal/client/search.go` — same empty-slice pattern.  
**Branch:** `loop-6/qa-fix-binter` (not pushed — apply same fix as volotea/ryanair)

---

## Branches pushed

| Branch | CLI | Status |
|--------|-----|--------|
| `loop-6/qa-fix-ryanair` | ryanair | **Pushed** — fixes `flights: null` |
| `loop-6/qa-fix-volotea` | volotea | Fix committed locally; push may need retry |
| `loop-6/qa-fix-binter` | binter | Fix not yet committed |

## Recommendations

1. **Merge `loop-6/qa-fix-ryanair`** into main — confirmed fix, no regressions on STN/PMI routes.
2. **Apply empty-slice fix** to volotea + binter on `loop-6/volotea-binter` before marking live in README.
3. **Ryanair input validation:** consider rejecting unknown IATA codes and past dates with exit 1 instead of empty success JSON.
4. **Vueling 404 handling:** optional improvement — map API 404 to empty `flights: []` JSON instead of stderr-only error for agent ergonomics.
5. **Shared helper:** add `travelkit` helper to guarantee `[]T{}` serialization for empty result slices across all airline CLIs.

## Reproduce

```bash
# Build
for cli in ryanair vueling volotea binter; do
  (cd ${cli}-cli && go build -o /tmp/$cli ./cmd/$cli)
done

# Single search
/tmp/ryanair search --json --from MAD --to STN --depart 2026-07-05

# Doctor
/tmp/vueling session doctor --json | jq .

# Schema check
/tmp/ryanair search --json --from MAD --to BCN --depart 2026-07-05 | \
  python3 -c "import json,sys; d=json.load(sys.stdin); assert isinstance(d['flights'], list)"
```
