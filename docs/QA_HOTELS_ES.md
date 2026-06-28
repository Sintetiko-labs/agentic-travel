# QA — Spanish Hotel CLIs (live smoke tests)

**Date:** 2026-06-28  
**Runner:** `scripts/qa-hotels-es.py`  
**Cities:** Madrid, Barcelona, Palma, Valencia (Silken: Barcelona if no Madrid)

## Results summary

| CLI | Overall | Notes |
|-----|---------|-------|
| barcelo | WARN | page 2 empty when total ≤ 5 |
| riu | PASS | Bcn/Val n/a; Palma via Mallorca alias |
| catalonia | WARN | no Valencia property |
| h10 | PASS | Valencia n/a |
| palladium | PASS | Palma n/a |
| lopesan | PASS | mainland n/a; Canary-only chain |
| princess | PASS | |
| eurostars | WARN | page 2 warn |
| vincci | PASS | Palma n/a |
| sercotel | PASS | |
| silken | WARN | Madrid n/a; page 2 warn |

## Fixes (branch `loop-6/qa-fix-hotels-es`)

- `travelkit/destination/aliases.go` — city alias expansion
- `travelkit/hotel/readld.go` — shared read via search lookup + JSON-LD
- Per-CLI `read` + search fixes for riu, catalonia, h10, palladium, lopesan, princess, eurostars, vincci, sercotel, silken

## Re-run

```bash
python3 scripts/qa-hotels-es.py
```
