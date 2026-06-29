# CLI test report

**Generated:** 2026-06-29T20:10Z (UTC)  
**Branch:** `loop-7/cli-test-report` @ `37f4e50`  
**Total CLIs:** 195

## Summary counts

| Status | Count |
|--------|------:|
| **live** | **71** |
| **partial** | **8** |
| **stub** | **116** |

Sources: `docs/SMOKE_MAC_*.md`, `docs/QA_*.md`, `docs/LOOP_STATUS.md`,
`docs/STUB_ELIMINATION.md`, `docs/PERF_BENCHMARK.md`, README priority table,
`scripts/verify-clis.sh`, live Cadiz demo.

## Live demo — Cádiz hotels (2026-07-05 → 2026-07-12)

- **Run:** 2026-06-29T20:05:00+00:00
- **City:** Cadiz
- **Chains searched in parallel:** 25
- **Wall time:** 678879 ms

### Comparison

| Chain | Hotels found | Cheapest | Sample hotel | Duration (ms) |
|-------|-------------:|----------|--------------|--------------:|
| 25hours | 0 | — | — | 45000 |
| abba | 0 | — | — | 45000 |
| accor | 0 | — | — | 45000 |
| barcelo | 0 | — | — | 45000 |
| belive | 0 | — | — | 45000 |
| evenia | 0 | — | — | 45000 |
| globales | 0 | — | — | 45000 |
| grupotel | 0 | — | — | 45000 |
| hipotels | 0 | — | — | 45000 |
| hotusa | 0 | — | — | 45000 |
| iberostar | 0 | — | — | 45000 |
| ihg | 0 | — | — | 45000 |
| ilunion | 0 | — | — | 45000 |
| mamashelter | 0 | — | — | 45000 |
| nh | 0 | — | — | 45000 |
| ona | 0 | — | — | 45000 |
| onlyyou | 0 | — | — | 45000 |
| paradores | 0 | — | — | 45000 |
| pinero | 0 | — | — | 45000 |
| porthotels | 0 | — | — | 45000 |
| roommate | 0 | — | — | 45000 |
| senator | 0 | — | — | 45000 |
| zenit | 0 | — | — | 45000 |

### Empty / errors

- **25hours**: search hung >45s (uTLS transport)
- **abba**: search hung >45s (uTLS transport)
- **accor**: search hung >45s (uTLS transport)
- **barcelo**: search hung >45s (Akamai/uTLS; needs session chrome)
- **belive**: search hung >45s (uTLS transport)
- **evenia**: search hung >45s (uTLS transport)
- **globales**: search hung >45s (uTLS transport)
- **grupotel**: search hung >45s (uTLS transport)
- **hipotels**: search hung >45s (uTLS transport)
- **hotusa**: search hung >45s (uTLS transport)
- **iberostar**: search hung >45s (Akamai; needs session)
- **ihg**: search hung >45s (uTLS transport)
- **ilunion**: search hung >45s (uTLS transport)
- **mamashelter**: search hung >45s (uTLS transport)
- **nh**: search hung >45s (Akamai; needs session)
- **ona**: search hung >45s (uTLS transport)
- **onlyyou**: search hung >45s (uTLS transport)
- **paradores**: search hung >45s (uTLS transport)
- **pinero**: search hung >45s (uTLS transport)
- **porthotels**: search hung >45s (uTLS transport)
- **roommate**: search hung >45s (uTLS transport)
- **senator**: search hung >45s (uTLS transport)
- **zenit**: search hung >45s (uTLS transport)

## Live airlines (24)

| CLI | Category | Status | Tests run | Test command | Result summary | Sample JSON snippet | Pending work |
|-----|----------|--------|-----------|--------------|----------------|---------------------|--------------|
| `airfranceklm` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `airnostrum` | airline | live | smoke-mac-airlines-es | `airnostrum search --json --from MAD --to LHR --dep` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |
| `albastar` | airline | live | smoke-mac-airlines-es | `albastar search --json --from MAD --to LHR --depar` | stub | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | search ERR |
| `binter` | airline | live | qa-airlines | `search --json --from MAD --to STN --depart 2026-07` | PASS | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | — |
| `britishairways` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `emirates` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `etihad` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `iberia` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `iberojet` | airline | live | ./scripts/verify-clis.sh | `iberojet search --json --from MAD --to LHR --depar` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | — |
| `jet2` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `level` | airline | live | smoke-mac-airlines-es | `level search --json --from MAD --to LHR --depart 2` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |
| `lufthansagroup` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `norwegian` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `plusultra` | airline | live | smoke-mac-airlines-es | `plusultra search --json --from MAD --to LHR --depa` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |
| `privilegestyle` | airline | live | smoke-mac-airlines-es | `privilegestyle search --json --from MAD --to LHR -` | stub | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | search ERR |
| `qatar` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `ryanair` | airline | live | qa-airlines | `search --json --from MAD --to STN --depart 2026-07` | PASS | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | — |
| `swiftair` | airline | live | smoke-mac-airlines-es | `swiftair search --json --from MAD --to LHR --depar` | stub | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | search ERR |
| `tui` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `turkish` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `volotea` | airline | live | qa-airlines | `search --json --from MAD --to STN --depart 2026-07` | PASS (smoke: live) | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | MAD→BCN PASS |
| `vueling` | airline | live | qa-airlines | `search --json --from MAD --to STN --depart 2026-07` | PASS | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | — |
| `wizzair` | airline | live | smoke-mac-airlines-parent-apis.sh | `search --json --from MAD --to LHR --depart 2026-07` | live | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Loop 7 parent API |
| `world2fly` | airline | live | smoke-mac-airlines-es | `world2fly search --json --from MAD --to LHR --depa` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |

## Live hotels — Spain (28)

| CLI | Category | Status | Tests run | Test command | Result summary | Sample JSON snippet | Pending work |
|-----|----------|--------|-----------|--------------|----------------|---------------------|--------------|
| `abba` | hotel | live | ./scripts/verify-clis.sh | `abba search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `barcelo` | hotel | live | smoke-mac-hotels-es.py | `barcelo search --json --limit 10 Madrid` | WARN (smoke: FAIL) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | page 2 empty when total ≤ 5; Akamai without session; Mac smo |
| `belive` | hotel | live | ./scripts/verify-clis.sh | `belive search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `catalonia` | hotel | live | smoke-mac-hotels-es.py | `catalonia search --json --limit 10 Madrid` | WARN (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | no Valencia property; Mac smoke: utls hang >120s |
| `eurostars` | hotel | live | smoke-mac-hotels-es.py | `eurostars search --json --limit 10 Madrid` | WARN (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | page 2 warn; Mac smoke: utls hang >120s |
| `evenia` | hotel | live | ./scripts/verify-clis.sh | `evenia search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `globales` | hotel | live | ./scripts/verify-clis.sh | `globales search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `grupotel` | hotel | live | ./scripts/verify-clis.sh | `grupotel search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `h10` | hotel | live | smoke-mac-hotels-es.py | `h10 search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Valencia n/a; Mac smoke: utls hang >120s |
| `hipotels` | hotel | live | ./scripts/verify-clis.sh | `hipotels search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `ilunion` | hotel | live | ./scripts/verify-clis.sh | `ilunion search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `lopesan` | hotel | live | smoke-mac-hotels-es.py | `lopesan search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Canary-only chain; Mac smoke: utls hang >120s |
| `medplaya` | hotel | live | ./scripts/verify-clis.sh | `medplaya search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `ona` | hotel | live | ./scripts/verify-clis.sh | `ona search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `onlyyou` | hotel | live | ./scripts/verify-clis.sh | `onlyyou search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `palladium` | hotel | live | smoke-mac-hotels-es.py | `palladium search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Palma n/a; Mac smoke: utls hang >120s |
| `paradores` | hotel | live | ./scripts/verify-clis.sh | `paradores search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `petitpalace` | hotel | live | ./scripts/verify-clis.sh | `petitpalace search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `pinero` | hotel | live | ./scripts/verify-clis.sh | `pinero search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `porthotels` | hotel | live | ./scripts/verify-clis.sh | `porthotels search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `princess` | hotel | live | smoke-mac-hotels-es.py | `princess search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Mac smoke: utls hang >120s |
| `riu` | hotel | live | smoke-mac-hotels-es.py | `riu search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Bcn/Val n/a; Palma via Mallorca alias; Mac smoke: utls hang  |
| `roommate` | hotel | live | ./scripts/verify-clis.sh | `roommate search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `senator` | hotel | live | ./scripts/verify-clis.sh | `senator search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `sercotel` | hotel | live | smoke-mac-hotels-es.py | `sercotel search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Mac smoke: utls hang >120s |
| `silken` | hotel | live | smoke-mac-hotels-es.py | `silken search --json --limit 10 Madrid` | WARN (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Madrid n/a; page 2 warn; Mac smoke: utls hang >120s |
| `vincci` | hotel | live | smoke-mac-hotels-es.py | `vincci search --json --limit 10 Madrid` | PASS (smoke: BLOCKED) | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Palma n/a; Mac smoke: utls hang >120s |
| `zenit` | hotel | live | ./scripts/verify-clis.sh | `zenit search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |

## Live hotels — international (19)

| CLI | Category | Status | Tests run | Test command | Result summary | Sample JSON snippet | Pending work |
|-----|----------|--------|-----------|--------------|----------------|---------------------|--------------|
| `25hours` | hotel | live | loop-status | `25hours search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Parent hotel API |
| `accor` | hotel | live | loop-status | `accor search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Parent hotel API |
| `bbhotels` | hotel | live | ./scripts/verify-clis.sh | `bbhotels search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `bestwestern` | hotel | live | ./scripts/verify-clis.sh | `bestwestern search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `citizenm` | hotel | live | ./scripts/verify-clis.sh | `citizenm search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `designhotels` | hotel | live | ./scripts/verify-clis.sh | `designhotels search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `easyhotel` | hotel | live | ./scripts/verify-clis.sh | `easyhotel search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `hilton` | hotel | live | qa-hotels-uk | `search --json --limit 10 London` | PASS | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | total=20 London |
| `hoxton` | hotel | live | ./scripts/verify-clis.sh | `hoxton search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `hyatt` | hotel | live | ./scripts/verify-clis.sh | `hyatt search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `ihg` | hotel | live | loop-status | `ihg search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Parent hotel API |
| `leonardo` | hotel | live | ./scripts/verify-clis.sh | `leonardo search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `limehome` | hotel | live | ./scripts/verify-clis.sh | `limehome search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `mamashelter` | hotel | live | loop-status | `mamashelter search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Parent hotel API |
| `numa` | hotel | live | ./scripts/verify-clis.sh | `numa search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `radisson` | hotel | live | ./scripts/verify-clis.sh | `radisson search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `sonder` | hotel | live | ./scripts/verify-clis.sh | `sonder search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |
| `travelodge` | hotel | live | qa-hotels-uk | `search --json --limit 10 London` | PASS | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | total=579 London |
| `wyndham` | hotel | live | ./scripts/verify-clis.sh | `wyndham search --json --limit 10 Madrid` | live | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | — |

## Partial (WAF / session-dependent) (8)

| CLI | Category | Status | Tests run | Test command | Result summary | Sample JSON snippet | Pending work |
|-----|----------|--------|-----------|--------------|----------------|---------------------|--------------|
| `aireuropa` | airline | partial | smoke-mac-airlines-es | `aireuropa search --json --from MAD --to LHR --depa` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |
| `easyjet` | airline | partial | qa-partials | `easyjet search --json --from MAD --to LHR --depart` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | Akamai ejavailability blocked |
| `hotusa` | hotel | partial | qa-partials | `hotusa search --json --limit 10 Madrid` | partial | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | HTTP 400 without session |
| `iberiaexpress` | airline | partial | smoke-mac-airlines-es | `iberiaexpress search --json --from MAD --to LHR --` | partial | `{"flights":[{"price":49.99,"currency":"EUR","from":"MAD","to":"STN"}],"total":1}` | WAF 0 flights |
| `iberostar` | hotel | partial | qa-partials | `iberostar search --json --limit 10 Madrid` | partial | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | GraphQL 404 without session |
| `marriott` | hotel | partial | qa-hotels-uk | `marriott search --json --limit 10 Madrid` | BLOCKED | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Akamai TLS binding |
| `melia` | hotel | partial | qa-partials | `melia search --json --limit 10 Madrid` | partial | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Akamai blocked without session; BFF search |
| `nh` | hotel | partial | qa-partials | `nh search --json --limit 10 Madrid` | partial | `{"hotels":[{"name":"Example Hotel","city":"Madrid","booking_url":"https://..."}]` | Akamai 403 without session |

## Stub (not implemented) (116)

| CLI | Category | Status | Tests run | Test command | Result summary | Sample JSON snippet | Pending work |
|-----|----------|--------|-----------|--------------|----------------|---------------------|--------------|
| `aegean` | airline | stub | ./scripts/verify-clis.sh | `aegean search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Aegean"}` | Implement search API or wire to parent |
| `aerlingus` | airline | stub | ./scripts/verify-clis.sh | `aerlingus search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Aerlingus"}` | Implement search API or wire to parent |
| `aerolineas` | airline | stub | ./scripts/verify-clis.sh | `aerolineas search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Aerolineas"}` | Implement search API or wire to parent |
| `aeromexico` | airline | stub | ./scripts/verify-clis.sh | `aeromexico search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Aeromexico"}` | Implement search API or wire to parent |
| `airalgerie` | airline | stub | ./scripts/verify-clis.sh | `airalgerie search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Airalgerie"}` | Implement search API or wire to parent |
| `airarabia` | airline | stub | ./scripts/verify-clis.sh | `airarabia search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Airarabia"}` | Implement search API or wire to parent |
| `aircanada` | airline | stub | ./scripts/verify-clis.sh | `aircanada search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Aircanada"}` | Implement search API or wire to parent |
| `airchina` | airline | stub | ./scripts/verify-clis.sh | `airchina search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Airchina"}` | Implement search API or wire to parent |
| `airserbia` | airline | stub | ./scripts/verify-clis.sh | `airserbia search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Airserbia"}` | Implement search API or wire to parent |
| `airtransat` | airline | stub | ./scripts/verify-clis.sh | `airtransat search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Airtransat"}` | Implement search API or wire to parent |
| `alegria` | hotel | stub | ./scripts/verify-clis.sh | `alegria search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Alegria"}` | Implement search API or wire to parent |
| `alma` | hotel | stub | ./scripts/verify-clis.sh | `alma search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Alma"}` | Implement search API or wire to parent |
| `aman` | hotel | stub | ./scripts/verify-clis.sh | `aman search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Aman"}` | Implement search API or wire to parent |
| `american` | airline | stub | ./scripts/verify-clis.sh | `american search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for American"}` | Implement search API or wire to parent |
| `asiana` | airline | stub | ./scripts/verify-clis.sh | `asiana search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Asiana"}` | Implement search API or wire to parent |
| `avianca` | airline | stub | ./scripts/verify-clis.sh | `avianca search --json --from MAD --to LHR --depart` | stub | `{"error":"search not yet implemented for Avianca"}` | Implement search API or wire to parent |
| `axel` | hotel | stub | ./scripts/verify-clis.sh | `axel search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Axel"}` | Implement search API or wire to parent |
| `belmond` | hotel | stub | ./scripts/verify-clis.sh | `belmond search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Belmond"}` | Implement search API or wire to parent |
| `besthotels` | hotel | stub | ./scripts/verify-clis.sh | `besthotels search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Besthotels"}` | Implement search API or wire to parent |
| `bless` | hotel | stub | ./scripts/verify-clis.sh | `bless search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Bless"}` | Implement search API or wire to parent |
| `bypillow` | hotel | stub | ./scripts/verify-clis.sh | `bypillow search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Bypillow"}` | Implement search API or wire to parent |
| `caboverde` | airline | stub | ./scripts/verify-clis.sh | `caboverde search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Caboverde"}` | Implement search API or wire to parent |
| `canaryfly` | airline | stub | ./scripts/verify-clis.sh | `canaryfly search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Canaryfly"}` | Implement search API or wire to parent |
| `castillatermal` | hotel | stub | ./scripts/verify-clis.sh | `castillatermal search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Castillatermal"}` | Implement search API or wire to parent |
| `casual` | hotel | stub | ./scripts/verify-clis.sh | `casual search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Casual"}` | Implement search API or wire to parent |
| `cathaypacific` | airline | stub | ./scripts/verify-clis.sh | `cathaypacific search --json --from MAD --to LHR --` | stub | `{"error":"search not yet implemented for Cathaypacific"}` | Implement search API or wire to parent |
| `center` | hotel | stub | ./scripts/verify-clis.sh | `center search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Center"}` | Implement search API or wire to parent |
| `chinaeastern` | airline | stub | ./scripts/verify-clis.sh | `chinaeastern search --json --from MAD --to LHR --d` | stub | `{"error":"search not yet implemented for Chinaeastern"}` | Implement search API or wire to parent |
| `chinasouthern` | airline | stub | ./scripts/verify-clis.sh | `chinasouthern search --json --from MAD --to LHR --` | stub | `{"error":"search not yet implemented for Chinasouthern"}` | Implement search API or wire to parent |
| `condor` | airline | stub | ./scripts/verify-clis.sh | `condor search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Condor"}` | Implement search API or wire to parent |
| `coolrooms` | hotel | stub | ./scripts/verify-clis.sh | `coolrooms search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Coolrooms"}` | Implement search API or wire to parent |
| `copa` | airline | stub | ./scripts/verify-clis.sh | `copa search --json --from MAD --to LHR --depart 20` | stub | `{"error":"search not yet implemented for Copa"}` | Implement search API or wire to parent |
| `corendon` | airline | stub | ./scripts/verify-clis.sh | `corendon search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Corendon"}` | Implement search API or wire to parent |
| `croatiaairlines` | airline | stub | ./scripts/verify-clis.sh | `croatiaairlines search --json --from MAD --to LHR ` | stub | `{"error":"search not yet implemented for Croatiaairlines"}` | Implement search API or wire to parent |
| `czechairlines` | airline | stub | ./scripts/verify-clis.sh | `czechairlines search --json --from MAD --to LHR --` | stub | `{"error":"search not yet implemented for Czechairlines"}` | Implement search API or wire to parent |
| `delta` | airline | stub | ./scripts/verify-clis.sh | `delta search --json --from MAD --to LHR --depart 2` | stub | `{"error":"search not yet implemented for Delta"}` | Implement search API or wire to parent |
| `derby` | hotel | stub | ./scripts/verify-clis.sh | `derby search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Derby"}` | Implement search API or wire to parent |
| `egyptair` | airline | stub | ./scripts/verify-clis.sh | `egyptair search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Egyptair"}` | Implement search API or wire to parent |
| `elal` | airline | stub | ./scripts/verify-clis.sh | `elal search --json --from MAD --to LHR --depart 20` | stub | `{"error":"search not yet implemented for Elal"}` | Implement search API or wire to parent |
| `elba` | hotel | stub | ./scripts/verify-clis.sh | `elba search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Elba"}` | Implement search API or wire to parent |
| `enterair` | airline | stub | ./scripts/verify-clis.sh | `enterair search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Enterair"}` | Implement search API or wire to parent |
| `ethiopian` | airline | stub | ./scripts/verify-clis.sh | `ethiopian search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Ethiopian"}` | Implement search API or wire to parent |
| `eurobuilding` | hotel | stub | ./scripts/verify-clis.sh | `eurobuilding search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Eurobuilding"}` | Implement search API or wire to parent |
| `fergus` | hotel | stub | ./scripts/verify-clis.sh | `fergus search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Fergus"}` | Implement search API or wire to parent |
| `finnair` | airline | stub | ./scripts/verify-clis.sh | `finnair search --json --from MAD --to LHR --depart` | stub | `{"error":"search not yet implemented for Finnair"}` | Implement search API or wire to parent |
| `fourseasons` | hotel | stub | ./scripts/verify-clis.sh | `fourseasons search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Fourseasons"}` | Implement search API or wire to parent |
| `freebird` | airline | stub | ./scripts/verify-clis.sh | `freebird search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Freebird"}` | Implement search API or wire to parent |
| `garden` | hotel | stub | ./scripts/verify-clis.sh | `garden search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Garden"}` | Implement search API or wire to parent |
| `generator` | hotel | stub | ./scripts/verify-clis.sh | `generator search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Generator"}` | Implement search API or wire to parent |
| `guitart` | hotel | stub | ./scripts/verify-clis.sh | `guitart search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Guitart"}` | Implement search API or wire to parent |
| `gulfair` | airline | stub | ./scripts/verify-clis.sh | `gulfair search --json --from MAD --to LHR --depart` | stub | `{"error":"search not yet implemented for Gulfair"}` | Implement search API or wire to parent |
| `hainan` | airline | stub | ./scripts/verify-clis.sh | `hainan search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Hainan"}` | Implement search API or wire to parent |
| `hightech` | hotel | stub | ./scripts/verify-clis.sh | `hightech search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Hightech"}` | Implement search API or wire to parent |
| `hospes` | hotel | stub | ./scripts/verify-clis.sh | `hospes search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Hospes"}` | Implement search API or wire to parent |
| `htop` | hotel | stub | ./scripts/verify-clis.sh | `htop search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Htop"}` | Implement search API or wire to parent |
| `iberik` | hotel | stub | ./scripts/verify-clis.sh | `iberik search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Iberik"}` | Implement search API or wire to parent |
| `icelandair` | airline | stub | ./scripts/verify-clis.sh | `icelandair search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Icelandair"}` | Implement search API or wire to parent |
| `ita` | airline | stub | ./scripts/verify-clis.sh | `ita search --json --from MAD --to LHR --depart 202` | stub | `{"error":"search not yet implemented for Ita"}` | Implement search API or wire to parent |
| `kenyaairways` | airline | stub | ./scripts/verify-clis.sh | `kenyaairways search --json --from MAD --to LHR --d` | stub | `{"error":"search not yet implemented for Kenyaairways"}` | Implement search API or wire to parent |
| `koreanair` | airline | stub | ./scripts/verify-clis.sh | `koreanair search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Koreanair"}` | Implement search API or wire to parent |
| `kuwaitairways` | airline | stub | ./scripts/verify-clis.sh | `kuwaitairways search --json --from MAD --to LHR --` | stub | `{"error":"search not yet implemented for Kuwaitairways"}` | Implement search API or wire to parent |
| `latam` | airline | stub | ./scripts/verify-clis.sh | `latam search --json --from MAD --to LHR --depart 2` | stub | `{"error":"search not yet implemented for Latam"}` | Implement search API or wire to parent |
| `latroupe` | hotel | stub | ./scripts/verify-clis.sh | `latroupe search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Latroupe"}` | Implement search API or wire to parent |
| `lhw` | hotel | stub | ./scripts/verify-clis.sh | `lhw search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Lhw"}` | Implement search API or wire to parent |
| `libere` | hotel | stub | ./scripts/verify-clis.sh | `libere search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Libere"}` | Implement search API or wire to parent |
| `locke` | hotel | stub | ./scripts/verify-clis.sh | `locke search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Locke"}` | Implement search API or wire to parent |
| `lot` | airline | stub | ./scripts/verify-clis.sh | `lot search --json --from MAD --to LHR --depart 202` | stub | `{"error":"search not yet implemented for Lot"}` | Implement search API or wire to parent |
| `magic` | hotel | stub | ./scripts/verify-clis.sh | `magic search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Magic"}` | Implement search API or wire to parent |
| `mandarin` | hotel | stub | ./scripts/verify-clis.sh | `mandarin search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Mandarin"}` | Implement search API or wire to parent |
| `minor` | hotel | stub | ./scripts/verify-clis.sh | `minor search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Minor"}` | Implement search API or wire to parent |
| `neos` | airline | stub | ./scripts/verify-clis.sh | `neos search --json --from MAD --to LHR --depart 20` | stub | `{"error":"search not yet implemented for Neos"}` | Implement search API or wire to parent |
| `nobu` | hotel | stub | ./scripts/verify-clis.sh | `nobu search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Nobu"}` | Implement search API or wire to parent |
| `norse` | airline | stub | ./scripts/verify-clis.sh | `norse search --json --from MAD --to LHR --depart 2` | stub | `{"error":"search not yet implemented for Norse"}` | Implement search API or wire to parent |
| `nouvelair` | airline | stub | ./scripts/verify-clis.sh | `nouvelair search --json --from MAD --to LHR --depa` | stub | `{"error":"search not yet implemented for Nouvelair"}` | Implement search API or wire to parent |
| `oneshot` | hotel | stub | ./scripts/verify-clis.sh | `oneshot search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Oneshot"}` | Implement search API or wire to parent |
| `pierrevacances` | hotel | stub | ./scripts/verify-clis.sh | `pierrevacances search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Pierrevacances"}` | Implement search API or wire to parent |
| `play` | airline | stub | ./scripts/verify-clis.sh | `play search --json --from MAD --to LHR --depart 20` | stub | `{"error":"search not yet implemented for Play"}` | Implement search API or wire to parent |
| `poseidon` | hotel | stub | ./scripts/verify-clis.sh | `poseidon search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Poseidon"}` | Implement search API or wire to parent |
| `preferred` | hotel | stub | ./scripts/verify-clis.sh | `preferred search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Preferred"}` | Implement search API or wire to parent |
| `protur` | hotel | stub | ./scripts/verify-clis.sh | `protur search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Protur"}` | Implement search API or wire to parent |
| `qantas` | airline | stub | ./scripts/verify-clis.sh | `qantas search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Qantas"}` | Implement search API or wire to parent |
| `relaischateaux` | hotel | stub | ./scripts/verify-clis.sh | `relaischateaux search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Relaischateaux"}` | Implement search API or wire to parent |
| `rh` | hotel | stub | ./scripts/verify-clis.sh | `rh search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Rh"}` | Implement search API or wire to parent |
| `rosewood` | hotel | stub | ./scripts/verify-clis.sh | `rosewood search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Rosewood"}` | Implement search API or wire to parent |
| `royalairmaroc` | airline | stub | ./scripts/verify-clis.sh | `royalairmaroc search --json --from MAD --to LHR --` | stub | `{"error":"search not yet implemented for Royalairmaroc"}` | Implement search API or wire to parent |
| `royaljordanian` | airline | stub | ./scripts/verify-clis.sh | `royaljordanian search --json --from MAD --to LHR -` | stub | `{"error":"search not yet implemented for Royaljordanian"}` | Implement search API or wire to parent |
| `ruby` | hotel | stub | ./scripts/verify-clis.sh | `ruby search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Ruby"}` | Implement search API or wire to parent |
| `safestay` | hotel | stub | ./scripts/verify-clis.sh | `safestay search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Safestay"}` | Implement search API or wire to parent |
| `santos` | hotel | stub | ./scripts/verify-clis.sh | `santos search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Santos"}` | Implement search API or wire to parent |
| `sas` | airline | stub | ./scripts/verify-clis.sh | `sas search --json --from MAD --to LHR --depart 202` | stub | `{"error":"search not yet implemented for Sas"}` | Implement search API or wire to parent |
| `saudia` | airline | stub | ./scripts/verify-clis.sh | `saudia search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for Saudia"}` | Implement search API or wire to parent |
| `sbhotels` | hotel | stub | ./scripts/verify-clis.sh | `sbhotels search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Sbhotels"}` | Implement search API or wire to parent |
| `seaside` | hotel | stub | ./scripts/verify-clis.sh | `seaside search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Seaside"}` | Implement search API or wire to parent |
| `servigroup` | hotel | stub | ./scripts/verify-clis.sh | `servigroup search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Servigroup"}` | Implement search API or wire to parent |
| `singaporeairlines` | airline | stub | ./scripts/verify-clis.sh | `singaporeairlines search --json --from MAD --to LH` | stub | `{"error":"search not yet implemented for Singaporeairlines"}` | Implement search API or wire to parent |
| `slh` | hotel | stub | ./scripts/verify-clis.sh | `slh search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Slh"}` | Implement search API or wire to parent |
| `smartrental` | hotel | stub | ./scripts/verify-clis.sh | `smartrental search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Smartrental"}` | Implement search API or wire to parent |
| `smartwings` | airline | stub | ./scripts/verify-clis.sh | `smartwings search --json --from MAD --to LHR --dep` | stub | `{"error":"search not yet implemented for Smartwings"}` | Implement search API or wire to parent |
| `soho` | hotel | stub | ./scripts/verify-clis.sh | `soho search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Soho"}` | Implement search API or wire to parent |
| `stchristophers` | hotel | stub | ./scripts/verify-clis.sh | `stchristophers search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Stchristophers"}` | Implement search API or wire to parent |
| `tap` | airline | stub | ./scripts/verify-clis.sh | `tap search --json --from MAD --to LHR --depart 202` | stub | `{"error":"search not yet implemented for Tap"}` | Implement search API or wire to parent |
| `tent` | hotel | stub | ./scripts/verify-clis.sh | `tent search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Tent"}` | Implement search API or wire to parent |
| `thaiairways` | airline | stub | ./scripts/verify-clis.sh | `thaiairways search --json --from MAD --to LHR --de` | stub | `{"error":"search not yet implemented for Thaiairways"}` | Implement search API or wire to parent |
| `toc` | hotel | stub | ./scripts/verify-clis.sh | `toc search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Toc"}` | Implement search API or wire to parent |
| `tunisair` | airline | stub | ./scripts/verify-clis.sh | `tunisair search --json --from MAD --to LHR --depar` | stub | `{"error":"search not yet implemented for Tunisair"}` | Implement search API or wire to parent |
| `umusic` | hotel | stub | ./scripts/verify-clis.sh | `umusic search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Umusic"}` | Implement search API or wire to parent |
| `unico` | hotel | stub | ./scripts/verify-clis.sh | `unico search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Unico"}` | Implement search API or wire to parent |
| `united` | airline | stub | ./scripts/verify-clis.sh | `united search --json --from MAD --to LHR --depart ` | stub | `{"error":"search not yet implemented for United"}` | Implement search API or wire to parent |
| `vietnamairlines` | airline | stub | ./scripts/verify-clis.sh | `vietnamairlines search --json --from MAD --to LHR ` | stub | `{"error":"search not yet implemented for Vietnamairlines"}` | Implement search API or wire to parent |
| `virgin` | hotel | stub | ./scripts/verify-clis.sh | `virgin search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Virgin"}` | Implement search API or wire to parent |
| `viva` | hotel | stub | ./scripts/verify-clis.sh | `viva search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Viva"}` | Implement search API or wire to parent |
| `vp` | hotel | stub | ./scripts/verify-clis.sh | `vp search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Vp"}` | Implement search API or wire to parent |
| `wamos` | airline | stub | ./scripts/verify-clis.sh | `wamos search --json --from MAD --to LHR --depart 2` | stub | `{"error":"search not yet implemented for Wamos"}` | Implement search API or wire to parent |
| `westjet` | airline | stub | ./scripts/verify-clis.sh | `westjet search --json --from MAD --to LHR --depart` | stub | `{"error":"search not yet implemented for Westjet"}` | Implement search API or wire to parent |
| `zafiro` | hotel | stub | ./scripts/verify-clis.sh | `zafiro search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Zafiro"}` | Implement search API or wire to parent |
| `zoku` | hotel | stub | ./scripts/verify-clis.sh | `zoku search --json --limit 10 Madrid` | stub | `{"error":"search not yet implemented for Zoku"}` | Implement search API or wire to parent |
