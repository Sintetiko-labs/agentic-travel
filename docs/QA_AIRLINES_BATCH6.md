# QA Smoke — Major Airlines Batch 6

**Branch:** `loop-7/airlines-major-batch-6`  
**Route:** MAD→LHR one-way, depart 2026-07-15  
**Script:** `scripts/smoke-mac-airlines-major-batch6.sh`

| CLI | API |
|-----|-----|
| qatar | DAPI `/fares` + `/calendar-fares` fallback |
| emirates | `POST /service/search/search-results` |
| etihad | `GET /edge/deeplink/calendar` |
| britishairways | `GET /api/grp/v1/bff/calendar` |
| turkish | `POST /api/v1/availability` |
| norwegian | `GET /resourceipr/api/airportpaircalendarflexibility` |
| jet2 | `POST /api/Availability/SearchFlights` |
| tui | `GET /flight/en/api/search/searchResults` |
| airfranceklm | AF air-bounds / KLM calendar / Transavia offers |
| lufthansagroup | LH lowestfares / Eurowings search.api.json |

Status codes: **LIVE** (flights>0), **EMPTY** (API OK, 0 flights), **WAF** (Akamai/Incapsula), **ERROR** (other).
