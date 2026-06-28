# MCP Travel Inventory — Loop 7 Research

Research date: **2026-06-28**. Scope: MCP (Model Context Protocol) servers for travel, hotels, and airlines that could **replace or augment** the **194 brand CLIs** in [Sintetiko-labs/agentic-travel](https://github.com/Sintetiko-labs/agentic-travel) (111 hotel + 83 airline CLIs, reverse-engineered from public sites).

Related docs: [MCP_VS_CLI.md](./MCP_VS_CLI.md) · [MCP_RELIABILITY.md](./MCP_RELIABILITY.md) · [MCP_SETUP.md](./MCP_SETUP.md) · [mcp/README.md](../mcp/README.md)

---

## Executive summary

| Finding | Detail |
|---------|--------|
| **Official brand MCPs are rare** | Of the major brands in agentic-travel (Marriott, Hilton, Accor, Ryanair, Vueling, Meliá, …), **none publish a first-party MCP** today. |
| **OTA / metasearch leaders lead** | Kiwi.com, Expedia, trivago, TourRadar, Wingie Enuygun, Turkish Airlines (lab), Sabre (enterprise) have official or semi-official MCP endpoints. |
| **Google Flights / Skyscanner / Kayak** | No official MCP from Google or Skyscanner/Kayak. Community scrapers and aggregators (`trvl`, `google-flights-mcp`, `mcp-skyscanner`, Kayak scraper) fill the gap. |
| **Booking.com** | ChatGPT launch partner (PhocusWire, Oct 2025) but **no public self-serve MCP URL** found; community Playwright/scraper MCPs and Apify/Bright Data proxies exist. |
| **Best augment for agentic-travel** | **Duffel MCP** (already integrated) + **Kiwi.com** + **trivago** + **Gondola** (multi-chain hotels) + **trvl** (broad meta, no keys) + brand CLIs for Spain/LCC depth. |

---

## Methodology

1. Web search: official MCP announcements (Kiwi, Expedia, Sabre, Booking.com, hotel chains, GDS).
2. GitHub: `gh search repos "mcp travel"` and manual curation of flight/hotel MCP repos.
3. Directories: AltexSoft travel MCP landscape (Dec 2025), [mcpservers.org](https://mcpservers.org), [Glama](https://glama.ai/mcp), [Cursor Directory](https://cursor.directory), [PulseMCP](https://www.pulsemcp.com), [Remote MCP List](https://remotemcplist.com).
4. Cross-check against agentic-travel brand list (README priority CLIs + chain groups).

---

## Tier 1 — Official / vendor-hosted (integrate first)

### Duffel MCP *(already in repo)*

| Field | Value |
|-------|-------|
| **Name** | duffel-mcp |
| **URL** | https://github.com/bokangsibolla/duffel-mcp |
| **Auth** | `DUFFEL_ACCESS_TOKEN` (free test token) |
| **Capabilities** | `search_flights`, `get_offer`, `search_stays` — **read-only**, no booking in v0.1 |
| **vs CLI** | **Higher reliability** than reverse-engineered CLIs for multi-carrier IATA search; **no WAF**. Misses many LCCs and brand-direct member rates. |
| **Mac / IP** | **None** — official REST API, any network |

**agentic-travel overlap:** Complements `ryanair`, `vueling`, `iberia`, etc. Does not replace brand-specific fares or Akamai-gated sites.

---

### Kiwi.com Flight Search MCP *(official)*

| Field | Value |
|-------|-------|
| **Name** | kiwi-com-flight-search |
| **URL** | https://mcp.kiwi.com · install: https://github.com/alpic-ai/kiwi-mcp-server-public |
| **Auth** | **None** (open remote MCP) |
| **Capabilities** | `search-flight` — one-way/round-trip, ±3 day flex, passengers, cabin; **booking link** per result. Search only (no multi-city, bags, account yet). |
| **vs CLI** | **More reliable** for meta flight discovery than scraping; **less depth** than `ryanair-cli` / `vueling-cli` for airline-direct pricing and session-gated live fares. |
| **Mac / IP** | **None** — hosted HTTPS endpoint |

**agentic-travel overlap:** Aggregator alternative to running 83 airline CLIs for discovery; still need CLIs for named LCC (Ryanair WAF, Vueling Skysales detail).

---

### Expedia Travel Recommendations MCP *(official)*

| Field | Value |
|-------|-------|
| **Name** | expedia-travel-recommendations-mcp |
| **URL** | https://github.com/ExpediaGroup/expedia-travel-recommendations-mcp |
| **Auth** | `EXPEDIA_API_KEY` |
| **Capabilities** | Hotels, flights, activities, cars — **recommendations** (not full booking lifecycle). stdio + streamable-http (`:9900/mcp`). |
| **vs CLI** | **Stable API** vs scraping OTAs; hotel coverage ≠ Spanish regional chains (Meliá, Barceló, H10). |
| **Mac / IP** | **None** |

**agentic-travel overlap:** City-level hotel/flight discovery; does not replace `melia`, `barcelo`, `h10`, etc.

---

### trivago MCP *(official)*

| Field | Value |
|-------|-------|
| **Name** | trivago accommodation metasearch |
| **URL** | https://mcp.trivago.com/mcp |
| **Auth** | **None** |
| **Capabilities** | `trivago-search-suggestions`, `trivago-accommodation-search`, `trivago-accommodation-radius-search` — metasearch across booking portals |
| **vs CLI** | **Reliable metasearch**; compares OTAs, not hotel-direct member/loyalty rates from `marriott-cli`, `hilton-cli`, `accor-cli`. |
| **Mac / IP** | **None** |

---

### Gondola MCP *(official product, free tier)*

| Field | Value |
|-------|-------|
| **Name** | Gondola hotel search |
| **URL** | https://mcp.gondola.ai/mcp · https://www.gondola.ai/mcp |
| **Auth** | **None** |
| **Capabilities** | `search_hotels` — real-time availability across **Marriott, Hilton, Hyatt, IHG, Accor, Wyndham**; member/AAA rates, points, booking links |
| **vs CLI** | **Potentially replaces discovery** for major chains vs partial `marriott-cli` / `hilton-cli` (Akamai); **verify** live rate parity for Spain properties. Booking via link, not in-MCP. |
| **Mac / IP** | **None** |

**agentic-travel overlap:** Direct substitute candidate for **Marriott, Hilton, Accor** *search* phases; CLIs still needed for WAF bypass, sub-brand filters, and Spain-only chains.

---

### Wingie Enuygun MCP *(official)*

| Field | Value |
|-------|-------|
| **Name** | enuygun / Wingie Enuygun |
| **URL** | https://mcp.enuygun.com/mcp |
| **Auth** | **OAuth 2.0** |
| **Capabilities** | 34 tools — flights, hotels, bus, car rental; search, booking, cancellations, weather, account |
| **vs CLI** | Full-stack OTA MCP; strong for Turkey/MENA; limited overlap with Spain-heavy agentic-travel brands. |
| **Mac / IP** | **None** (remote) |

---

### Turkish Airlines MCP *(Digital Lab — semi-official)*

| Field | Value |
|-------|-------|
| **Name** | turkish-airlines |
| **URL** | https://mcp.turkishtechlab.com/mcp |
| **Auth** | **OAuth 2.0** (Miles&Smiles for profile tools) |
| **Capabilities** | 13–14 tools: flight search/status, booking lookup (PNR), check-in, baggage, promotions, loyalty profile |
| **vs CLI** | **More reliable** than scraping if user flies TK; single-carrier only. |
| **Mac / IP** | **None** |

**agentic-travel overlap:** No `turkish` CLI in priority list; optional add-on for TK routes.

---

### Sabre MCP *(official, enterprise)*

| Field | Value |
|-------|-------|
| **Name** | Sabre agentic APIs / MCP |
| **URL** | Embedded in SabreMosaic — **no public consumer URL** |
| **Auth** | Sabre agency/TMC credentials |
| **Capabilities** | GDS shopping, booking, servicing — flights, hotels, cars; agentic rebooking, disruption handling |
| **vs CLI** | **Gold standard reliability** for licensed agencies; **not accessible** to hobby/agentic-travel consumers without Sabre contract. |
| **Mac / IP** | Enterprise deployment |

---

### TourRadar MCP *(official)*

| Field | Value |
|-------|-------|
| **Name** | TourRadar tours |
| **URL** | TourRadar MCP (ChatGPT / partner integrations — see TourRadar press) |
| **Auth** | Partner / platform-specific |
| **Capabilities** | 50k+ multi-day tours — search, itineraries, brochures, **booking via AI** |
| **vs CLI** | N/A — tours not in agentic-travel CLI scope |
| **Mac / IP** | Platform-dependent |

---

### Apaleo MCP *(official, PMS — alpha)*

| Field | Value |
|-------|-------|
| **Name** | Apaleo hospitality PMS |
| **URL** | Apaleo developer MCP (alpha) |
| **Auth** | Apaleo API credentials |
| **Capabilities** | 237 API endpoints — availability, bookings, payments, guest profiles (hotel ops, not consumer metasearch) |
| **vs CLI** | B2B property management; not a consumer hotel search replacement |
| **Mac / IP** | None |

---

## Tier 2 — Community wrappers over official APIs

### Amadeus MCP ecosystem

Amadeus **does not publish** a first-party MCP. Multiple community servers wrap [Amadeus for Developers](https://developers.amadeus.com):

| Server | URL | Auth | Tools (summary) | vs CLI | Mac/IP |
|--------|-----|------|-----------------|--------|--------|
| **CVamsi27/travel-mcp-server** | https://github.com/CVamsi27/travel-mcp-server | `AMADEUS_CLIENT_ID` + secret | 43 tools: flights, hotels, transfers, activities, analytics | **Licensed GDS-lite**; Amadeus Self-Service **sunsets Jul 2026** | None |
| **donghyun-chae/mcp-amadeus** | https://github.com/donghyun-chae/mcp-amadeus · PyPI `mcp-amadeus` | Amadeus OAuth | Flight offers search | Same | None |
| **lev-corrupted/travel-mcp-server** | https://github.com/lev-corrupted/travel-mcp-server | Amadeus + optional AviationStack | Flights, hotels, airports, cheapest dates, flight tracking | Same + tracking | None |
| **prathush21/travel-amadeus-mcp** | LobeHub / Playbooks listings | Amadeus keys | 50+ tools (per directory listings) | Broadest Amadeus surface | None |

**agentic-travel overlap:** `aireuropa-cli` uses Amadeus `dapi` on the public site — MCP Amadeus may overlap for AE routes but **does not replace** Spain regional hotels or LCC CLIs.

---

### TripAdvisor Content API MCP

| Field | Value |
|-------|-------|
| **Name** | tripadvisor-mcp |
| **URL** | https://github.com/pab1it0/tripadvisor-mcp |
| **Auth** | TripAdvisor Content API key |
| **Capabilities** | Search locations, details, reviews, photos, nearby POI — **no booking** |
| **vs CLI** | Reviews/POI layer; not a hotel booking substitute |
| **Mac / IP** | None |

---

### TravelCode MCP *(corporate travel API)*

| Field | Value |
|-------|-------|
| **Name** | mcp-travelcode |
| **URL** | https://github.com/Travel-Code-Inc/mcp-travelcode · hosted `mcp.travel-code.com` |
| **Auth** | TravelCode API / OAuth (hosted) |
| **Capabilities** | Airports, airlines, flight search (async polling), **booking management**, delay stats |
| **vs CLI** | Commercial TMC API; full book/servicing vs read-only Duffel |
| **Mac / IP** | None |

---

### SerpAPI Google Travels MCP

| Field | Value |
|-------|-------|
| **Name** | mcp-google-travels |
| **URL** | https://github.com/modellers/mcp-google-travels |
| **Auth** | `SERPAPI_API_KEY` (paid) |
| **Capabilities** | Google Flights + Google Hotels + vacation rentals via SerpAPI — browse only |
| **vs CLI** | **Paid, stable** vs free scrapers; still indirect Google data |
| **Mac / IP** | None (SerpAPI cloud) |

---

## Tier 3 — Reverse-engineered / scraper MCPs (mirror CLI risk profile)

These use the same class of techniques as agentic-travel CLIs (undocumented APIs, Playwright, scraping). Reliability and ToS risk are **similar or worse** than maintained CLIs.

| Name | URL | Auth | Capabilities | vs agentic-travel CLI | Mac/residential IP |
|------|-----|------|--------------|----------------------|-------------------|
| **trvl** | https://github.com/MikkoParkkola/trvl | **None** | Unified `travel` router: Google Flights, Kiwi, Skiplagged, 6 hotel sources, trains, ferries, 36 fare hacks | **Broad meta** — could reduce need for many airline CLIs for *discovery*; **not** brand-native booking | Scrapers may need **residential IP** for Google; single Go binary |
| **google-flights-mcp** | https://github.com/andreacappelletti97/google-flights-mcp | None | 12 tools — search, price calendar, layover analysis, booking URLs | vs `fli`/CLI: same undocumented Google endpoint fragility | **Locale/currency from IP**; rate limits |
| **Flyan (Ryanair)** | https://github.com/victorlane/Flyan | None | `find_flights`, `find_anywhere_under`, `cheapest_per_day` | **Direct overlap** with `ryanair-cli` — MCP is thinner; CLI has session/doctor | Ryanair API often **WAF-sensitive** — same as CLI |
| **mcp-skyscanner** | https://github.com/shadyvb/mcp-skyscanner | None | `search_flights`, `search_airports` via unofficial Skyscanner client | Experimental; **no official Skyscanner MCP** | `curl_cffi` TLS fingerprinting; IP may matter |
| **Flight-Search-MCP (Kayak)** | https://github.com/alan5543/Flight-Search-MCP | Smithery remote optional | Kayak scrape — prices, bags, stops | Overlaps meta search; fragile | Scraper — **residential recommended** |
| **stays (Google Hotels)** | https://github.com/him229/stays | None | `search_hotels`, `get_hotel_details` — Google Hotels batchexecute | vs `hilton-cli`/`marriott-cli`: meta not direct chain | Google anti-bot — **IP/browser risk** |
| **JMMonte/travel-skill** | https://github.com/JMMonte/travel-skill | None | Flights (`fli`), hotels (Playwright), Airbnb, GTFS transit | Composite scraper stack | Hotels module uses **headless Chrome** |
| **mcp-booking (Playwright)** | https://github.com/markswendsen-code/mcp-booking | Booking.com account for book/history | 14 tools — search, availability, reservations; **booking = browser link only** | vs hypothetical `booking-cli`: same Playwright fragility | **Mac + browser** for login session |
| **Bright Data Booking MCP** | https://brightdata.com/ai/mcp-server/booking | Bright Data token | Scrape public Booking.com data | Enterprise proxy scrape | Bright Data **residential/datacenter** proxy network |
| **Apify travel-mcp** | https://apify.com/nexgendata/travel-mcp-server | Apify token | Airbnb, Booking.com, TripAdvisor, Craigslist rentals | Paid scrape aggregator | Apify cloud IPs |
| **openbnb / Airbnb MCP** | https://github.com/openbnb-org/mcp-server-airbnb | None | Airbnb search + listing details (scrape) | No Airbnb CLI in repo | Scrape — periodic breakage |
| **Apify Hilton/Marriott scrapers** | Apify marketplace | Apify token | Directory/rates scrapers per chain | Overlap with `hilton-cli`, `marriott-cli` | Cloud scrape |

---

## Tier 4 — Brands requested — MCP status

| Brand | Official MCP? | Best available MCP alternative | agentic-travel CLI |
|-------|---------------|-------------------------------|-------------------|
| **Marriott** | ❌ No | Gondola, Apify scraper, `stays`/Google Hotels | `marriott-cli` (partial, Akamai) |
| **Hilton** | ❌ No | Gondola, Apify scraper | `hilton-cli` (live UK) |
| **Accor** | ❌ No | Gondola | `accor-cli` |
| **Ryanair** | ❌ No | Flyan MCP, trvl, Kiwi meta | `ryanair-cli` (live) |
| **Amadeus** | ❌ No (API only) | Community Amadeus MCPs | Used by `aireuropa-cli` indirectly |
| **Sabre** | ✅ Enterprise only | Sabre Mosaic MCP (TMC) | — |
| **Google Flights** | ❌ No | google-flights-mcp, trvl, SerpAPI | — |
| **Skyscanner** | ❌ No | mcp-skyscanner, trvl | — |
| **Booking.com** | ⚠️ ChatGPT partner; no public MCP URL | mcp-booking, Apify, Bright Data, EmilyThaHuman/booking-mcp-server | — |
| **Expedia** | ✅ Yes | expedia-travel-recommendations-mcp | — |
| **Kayak** | ❌ No | Flight-Search-MCP, trvl | — |
| **TripAdvisor** | ❌ No (Content API) | pab1it0/tripadvisor-mcp, Apify | — |
| **Kiwi.com** | ✅ Yes | https://mcp.kiwi.com | — |
| **Vueling** | ❌ No | Duffel/Kiwi meta only | `vueling-cli` (live) |
| **Meliá / Barceló / H10 / …** | ❌ No | trivago/Booking meta only | Brand CLIs (Spain priority) |
| **Turkish Airlines** | ✅ Lab | mcp.turkishtechlab.com | — |
| **Duffel** | N/A (API vendor) | duffel-mcp *(integrated)* | — |
| **trivago** | ✅ Yes | mcp.trivago.com | — |

---

## Directory & registry findings

| Source | Travel-relevant entries |
|--------|-------------------------|
| **AltexSoft** (Dec 2025) | Kiwi, Turkish Airlines, Sabre, Expedia, TourRadar, Apaleo, Kismet, VariFlight, Flightradar24, Rakuten Travel, Airbnb unofficial, TripAdvisor unofficial, National Rail, NS Dutch Railways |
| **Anthropic MCP docs** | No curated travel list; recommends **official vendor servers only** |
| **OpenAI MCP / ChatGPT Apps** | Booking.com + Expedia as launch partners (Oct 2025); consumer adds via Apps settings |
| **Cursor Directory** | `trvl` listed |
| **Glama / LobeHub / Smithery / PulseMCP** | Amadeus wrappers, trvl, google-flights-mcp, tripadvisor-mcp, duffel variants |
| **GitHub `mcp travel` search** | 30+ repos; most are **planner demos** (mock data), not production travel APIs |

---

## Ranked integratable MCPs (value for agentic-travel)

Scoring: **coverage** of Madrid/London/Spain use cases, **reliability** vs CLI scrape, **setup friction**, **cost**, **complement** to existing 194 CLIs.

| Rank | MCP | Why | Replaces CLI? | Keep CLI for |
|------|-----|-----|---------------|--------------|
| **1** | **Duffel MCP** *(integrated)* | Official multi-carrier + stays; no WAF; already vendored in repo | Partial — multi-airline discovery | Ryanair, Vueling, Volotea, Binter, session partials |
| **2** | **Kiwi.com** | Official, zero-auth remote URL, booking links | Discovery / meta routes | Airline-direct LCC fares, Spain regional |
| **3** | **Gondola** | Free, multi-chain (Marriott/Hilton/Accor/…), member + points | Chain **search** for global brands | Akamai partials, Spain chains, sub-brand filters |
| **4** | **trivago** | Official, zero-auth hotel metasearch | OTA price comparison | Direct chain loyalty rates |
| **5** | **trvl** | No API keys; flights+hotels+ground in one binary; 66 tool aliases | **Broad discovery** across many providers | Named brand requests, booking confirmation, WAF brands |
| **6** | **Expedia MCP** | Official recommendations API | Package/hotel/flight ideas | Spanish regional hotels, LCCs |
| **7** | **google-flights-mcp** | Rich Google Flights tooling (12 tools) | Google-native flight UX | Airline-direct; endpoint fragility |
| **8** | **Flyan (Ryanair)** | Drop-in MCP for Ryanair-only prompts | Ryanair **search-only** shortcuts | Full fare detail, booking flow, WAF session |
| **9** | **Amadeus MCP** (CVamsi27 or lev-corrupted) | Broad GDS self-service | Legacy carriers, hotels by city code | Sunset Jul 2026; LCC gaps; Spain hotels |
| **10** | **TripAdvisor MCP** | Reviews and POI enrichment | Never replaces booking | — |
| **11** | **TravelCode MCP** | End-to-end corporate book/servicing | TMC workflows | Consumer Spain brands |
| **12** | **Wingie Enuygun** | Full OAuth travel stack | MENA/Turkey trips | Spain-focused repo |
| **13** | **mcp-skyscanner / Kayak MCP** | Meta flight scrape | Optional meta | Official API closed; scrape fragility |
| **14** | **mcp-booking / Apify / Bright Data** | Booking.com access without partner contract | OTA hotel search | ToS/cost; not chain-direct |
| **15** | **Turkish Airlines MCP** | Single-carrier official | TK-specific | Rest of airline CLIs |

### Do not prioritize for agentic-travel core

- **Sabre MCP** — enterprise-only, no public endpoint.
- **Mock planner MCPs** (`ismailrz/travel-mcp-server`, classroom demos) — fake data.
- **Apify per-chain scrapers** — redundant with maintained CLIs + Gondola unless CI needs cloud execution.

---

## Hybrid routing recommendation (MCP + CLI)

```
IF user names a brand in agentic-travel (ryanair, melia, barcelo, …)
   → use matching CLI (--json)

ELSE IF intent is multi-carrier or city-pair discovery
   → Duffel MCP or Kiwi.com MCP

ELSE IF intent is major-chain hotel (Marriott, Hilton, Accor, IHG, Hyatt)
   → Gondola MCP first; fallback marriott-cli / hilton-cli / accor-cli on empty or Akamai

ELSE IF intent is OTA / price comparison
   → trivago MCP or Expedia MCP

ELSE IF intent is broad trip planning (trains, ferries, hacks)
   → trvl MCP

ELSE IF WAF / session required (doctor flags, partial CLIs)
   → session chrome + brand CLI (unchanged)
```

---

## Mac / residential IP matrix

| MCP class | Residential IP required? | Notes |
|-----------|-------------------------|-------|
| Official API (Duffel, Amadeus, Expedia, Kiwi remote, trivago, Gondola) | **No** | Preferred for CI and cloud agents |
| OAuth remote (Enuygun, Turkish Airlines) | **No** | Token in client |
| Google Flights / Hotels scrapers | **Often yes** | Currency/locale from IP; anti-bot |
| Ryanair (Flyan) / Skyscanner unofficial | **Often yes** | Same WAF class as `ryanair-cli` |
| Playwright Booking.com MCP | **Yes for login** | Browser session on Mac |
| Apify / Bright Data | **Proxy provided** | Paid; not user's residential IP |

---

## Gaps — no MCP found (CLI remains authoritative)

Spanish/regional hotel groups: **Meliá, Barceló, Riu, Catalonia, H10, Palladium, Lopesan, Hotusa/Crisol, Vincci, Silken, Sercotel, Eurostars, …**

Spanish/LCC airlines: **Vueling, Volotea, Binter, Iberia Express, Air Europa, EasyJet** (partial), **Iberia**

Global chains without MCP: **Marriott, Hilton, Accor** (consumer) — use Gondola for search only

Meta without official MCP: **Google Flights, Skyscanner, Kayak**

---

## Next steps (loop 7)

1. Add **Kiwi** + **trivago** + **Gondola** to `.cursor/mcp.json.example` as optional remote servers.
2. Spike **trvl** binary alongside Duffel for hack-heavy discovery (document IP requirements).
3. Do **not** replace `ryanair-cli` with Flyan MCP until WAF parity tested on Mac residential ([MCP_RELIABILITY.md](./MCP_RELIABILITY.md)).
4. Track Booking.com for public MCP URL (ChatGPT Apps partner).
5. Plan Amadeus MCP migration before **July 2026** sunset — prefer **Duffel** path already started.

---

## References

- AltexSoft: [MCP Servers in Travel](https://www.altexsoft.com/blog/mcp-servers-travel/) (Dec 2025)
- PhocusWire: [How MCP could reshape travel](https://www.phocuswire.com/how-mcp-could-reshape-travel)
- AbrarNotes: [Travel MCP Server Directory](https://abrarnotes.com/travel-industry/travel-mcp-server-directory/2026/04)
- Kiwi.com: [MCP announcement](https://media.kiwi.com/company-news/kiwi-com-releases-mcp-server-prototype/)
- OpenAI: [Building MCP servers for ChatGPT Apps](https://developers.openai.com/api/docs/mcp)
