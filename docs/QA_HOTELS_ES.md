# QA — Spanish Hotel CLIs

Live smoke tests for 11 hotel CLIs on `main`. Runner: `scripts/qa-hotels-es.py`.

## Branches

| Branch | Scope |
|--------|--------|
| `loop-6/qa-fix-travelkit` | `travelkit/destination`, `travelkit/hotel/readld` |
| `loop-6/qa-fix-{cli}` | Per-chain read/search fixes |

## Last QA snapshot ([build agent](ccfbdb49-806c-4155-8cfb-1bc5eccacbce))

- **PASS:** h10
- **WARN:** barcelo, palladium
- **FAIL:** riu (limit/page/read), catalonia (Valencia), lopesan (build), princess/eurostars/vincci/sercotel/silken (read)

Geographic **n/a** (not fail): RIU Barcelona/Valencia; Lopesan mainland; Palladium/Vincci Palma; H10 Valencia; Silken Madrid.

## Re-run

```bash
python3 scripts/qa-hotels-es.py
```

Results → `scripts/qa-hotels-es-results.json`.
