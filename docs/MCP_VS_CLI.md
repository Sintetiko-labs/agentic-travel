# MCP vs CLI — Architecture for agentic-travel

**Loop 7** architecture decision: when agents should call **official MCP servers** (Duffel, Amadeus, Booking.com partner APIs) vs the **brand CLIs** in this monorepo, and how to combine both in a hybrid orchestration layer.

> This document is for **agent orchestrators** and human maintainers. It does not replace per-CLI READMEs or [AGENTS.md](../AGENTS.md).

---

## Executive summary

| Layer | Role | Best for |
|-------|------|----------|
| **MCP (official APIs)** | Broad, schema-stable, multi-carrier / multi-chain search | Discovery, price comparison, GDS coverage, booking handoff URLs |
| **CLI (this repo)** | Brand-native, reverse-engineered, Spain-heavy | Direct inventory from chains with no partner API, WAF bypass, regional OTAs |
| **Hybrid router** | MCP first, CLI fallback | Production agent flows that need both coverage and depth |

**Recommended default:** MCP for **aggregate flight + hotel search**; CLI for **named brand** or **Spain regional chain** requests where MCP returns no match or stale data.

---

## When to use official MCP

Use MCP when the user intent is **route- or city-level discovery** without naming a specific hotel group or LCC, or when you need **contractual, versioned APIs**.

### Duffel (flights)

| Use MCP when | Examples |
|--------------|----------|
| Multi-airline comparison on a city pair | "Cheapest MAD→LHR next Friday" |
| IATA-native search with structured offers | Round-trip, cabin, passenger counts |
| Stable JSON schema for agent tools | `offer_id`, segments, `total_amount` |
| Booking / order creation via supported carriers | Post-search checkout handoff |

**Do not** use Duffel MCP alone when the user explicitly asks for **Ryanair, Vueling, Volotea, Binter**, etc. — many LCCs are absent or incomplete in GDS/NDC aggregators. Fall back to the matching CLI (see [Hybrid routing](#hybrid-routing)).

### Amadeus (flights + hotels)

| Use MCP when | Examples |
|--------------|----------|
| GDS-backed carriers and interline | Long-haul, legacy flag carriers |
| Hotel search by geo / chain code where Amadeus has content | City + dates + star rating |
| Airport / city lookup, seat maps, branded fares (per licensed product) | Reference data before CLI deep-dive |

**Note:** Air Europa's public site already uses Amadeus `dapi` redirects (`aireuropa-cli`); MCP Amadeus may overlap but is **more reliable** than scrape when licensed. For **fare detail on LCCs**, still prefer brand CLI.

### Booking.com partner APIs (hotels)

| Use MCP when | Examples |
|--------------|----------|
| City / landmark hotel discovery | "Hotels near Sagrada Família" |
| Review scores, photos, property metadata | Comparison tables |
| Affiliate / partner booking links | Handoff when user does not care about chain |

**Do not** rely on Booking.com MCP for **Spanish regional chains** that list only on their own engine (Meliá direct, Barceló, H10, Hotusa/Crisol, Palladium, etc.) — inventory and member rates differ.

### MCP selection checklist

```
IF intent is discovery OR multi-brand OR booking handoff
   AND provider has official MCP + API key
   AND brand is NOT a named Spain regional chain / LCC
THEN use MCP
ELSE consider CLI (below)
```

---

## When CLI is still required

The monorepo exists because **321 brands** share **194 CLIs** with **no public partner API**. Agents must shell out to `{slug} search --json` for these cases.

### Brand-specific requests

User names a chain or sub-brand → map via `scripts/groups.json` or `{slug} brands`:

```bash
melia search --json --brand "Gran Meliá" Madrid
ryanair search --json --from MAD --to STN --depart 2026-07-15
```

| Signal in user message | CLI slug(s) |
|------------------------|-------------|
| Meliá, Paradisus, INNSiDE | `melia` |
| Barceló, Royal Hideaway | `barcelo` |
| Ryanair, Vueling, Volotea, Binter | `ryanair`, `vueling`, `volotea`, `binter` |
| NH, nhow | `nh` |
| Travelodge UK scenario | `travelodge`, `hilton`, `marriott` |

### No public API

| Category | Why CLI | Current status (loop 6) |
|----------|---------|-------------------------|
| Spain hotel groups | Proprietary BFF / GraphQL behind Akamai | `melia`, `nh`, `iberostar` **partial**; `barcelo`, `h10`, `hotusa` **live** |
| European LCCs | Private mobile/booking APIs, bot protection | `ryanair`, `vueling` **live**; `easyjet` **partial** |
| Canary / regional | Small carriers, no GDS | `binter`, `volotea` **live** |
| UK midscale | JSON or HTML listing APIs | `travelodge` **live**, `hilton` **live** |

### Spain regional chains (MCP gap)

These are **high priority for CLI**, not MCP substitution:

- **Balearic / Canarias:** Palladium, RIU, Iberostar, H10, Princess, Lopesan, Barceló
- **Peninsula independents:** Vincci, Silken, Sercotel, Eurostars, Hotusa/Crisol, Catalonia, Senator, Servigroup, MedPlaya, Magic Costa Blanca, …
- **Public sector / unique:** Paradores
- **Apartments / hybrid:** Ona, Pierre & Vacances España, Numa, Líbere

MCP hotel aggregators may surface **some** properties but often miss **member rates**, **package rules**, and **all-room categories**.

### Anti-bot and session chrome

CLIs support headed Chrome session harvest when MCP is not an option:

```bash
melia session chrome --wait --timeout 3m
melia session doctor --json
```

**MCP does not replace** Akamai/Incapsula bypass — that remains a CLI concern (`session chrome`, `travelkit` uTLS, `{PREFIX}_COOKIE`).

### CLI selection checklist

```
IF user names brand OR sub-brand in groups.json
   OR MCP returned empty / stale for that brand
   OR need member/direct rate / WAF-protected BFF
THEN use matching CLI (--json, respect REQUEST_DELAY)
```

---

## Hybrid architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Agent orchestrator                       │
│              (Cursor, Claude Code, custom loop)              │
└──────────────────────────┬──────────────────────────────────┘
                           │
              ┌────────────┴────────────┐
              ▼                         ▼
   ┌──────────────────────┐  ┌──────────────────────┐
   │   MCP router (first)   │  │  CLI router (fallback)│
   │  Duffel / Amadeus /    │  │  agentic-travel bins  │
   │  Booking.com partner   │  │  melia, ryanair, …    │
   └──────────┬─────────────┘  └──────────┬───────────┘
              │                           │
              └────────────┬──────────────┘
                           ▼
              ┌────────────────────────────┐
              │   Normalized result layer   │
              │   travelkit/types shapes    │
              │   hotels[], flights[]       │
              └────────────────────────────┘
```

### Hybrid routing

| Step | Action |
|------|--------|
| 1 | Parse intent: city pair / dates / named brand? |
| 2 | If **no brand** → MCP flight search (Duffel) and/or hotel search (Amadeus or Booking.com) |
| 3 | If **brand named** or MCP **miss** → resolve slug from `groups.json` → CLI `search --json` |
| 4 | Merge results; dedupe by property name + geo or flight number + date |
| 5 | For **availability / read** → MCP offer detail if `offer_id` exists; else CLI `read` / `availability` |
| 6 | Surface `booking_url` from whichever layer succeeded |

### Normalization contract

All layers should map into [`travelkit/types`](../travelkit/types/types.go):

- `HotelSearchResult` / `FlightSearchResult` with **non-null** `hotels` / `flights` arrays (use `[]` not `null`)
- `source` field: `duffel`, `amadeus`, `booking.com`, `melia-bff`, `ryanair-booking`, etc.
- Agents compare `price`, `currency`, `booking_url` across sources

### Example: Madrid → London weekend

| Need | Tool |
|------|------|
| Any flight MAD→LON | Duffel MCP |
| Ryanair STN | `ryanair search --json` (MCP may omit) |
| Hotel near Westminster | Booking.com MCP |
| Travelodge / Hilton UK | `travelodge`, `hilton` CLI |
| Meliá Gran Vía member rate | `melia` CLI only |

---

## Cost model

Rough economics for agent loops at moderate volume (~1k searches/day). Adjust for your Amadeus tier and Duffel markup.

| Approach | Setup cost | Marginal cost | Hidden cost |
|----------|------------|---------------|-------------|
| **Duffel MCP** | API key, test mode free | **Per offer / per order** (~1–3% + carrier fees) | LCC gaps → CLI fallback anyway |
| **Amadeus MCP** | Developer key → commercial tier | **Per call / per transaction** (Self-Service vs Enterprise) | Hotel content gaps in Spain chains |
| **Booking.com partner** | Partner approval | **Commission on booked stay** (CPA) | Not all chains; rate parity rules |
| **CLI scrape** | Engineering time (this repo) | **$0 API fees** | Residential IP, Chrome sessions, maintenance when sites change |
| **Session chrome** | Local Chrome | Operator time | Headed browser on agent host |

### Cost-aware routing rules

1. **Exploration / ranking** — prefer MCP if key already provisioned (predictable cents per call).
2. **High-volume LCC polling** — CLI + `{PREFIX}_REQUEST_DELAY` cheaper than repeated GDS queries with no hits.
3. **Named brand confirmation** — one CLI call beats multi-hop MCP + wrong property match.
4. **Production booking** — MCP or official API for **orders**; CLI for **search-only** unless ToS allows otherwise.

### API keys vs free scrape

| | MCP / official API | CLI scrape |
|--|-------------------|------------|
| **Billing** | Metered, invoice | Infra + maintenance |
| **Legal** | Contractual | Reverse-engineered; **local / personal use** per README |
| **Breakage** | Versioned deprecation notices | Silent HTML/API changes |
| **Best ROI** | Multi-brand discovery | Spain portfolio depth |

---

## Reliability matrix

Legend: **H** high, **M** medium, **L** low, **—** not applicable.

| Dimension | Duffel MCP | Amadeus MCP | Booking.com MCP | Brand CLI (live) | Brand CLI (partial) |
|-----------|:----------:|:-----------:|:---------------:|:----------------:|:-------------------:|
| **Schema stability** | H | H | H | M (travelkit JSON) | M |
| **Multi-brand coverage** | M (GDS/NDC) | H | H | L (one slug) | L |
| **Spain regional hotels** | L | M | M | **H** | M |
| **European LCC flights** | L | L | — | **H** | M |
| **Rate limits** | H (documented) | H | H | M (`REQUEST_DELAY`) | L (Akamai) |
| **Auth / session** | API key | API key | Partner token | Optional cookie / chrome | **chrome required** |
| **Empty-result JSON** | H | H | H | H (if `[]` not `null`) | M |
| **Booking completion** | H (where supported) | H (licensed) | H (affiliate) | URL handoff only | URL handoff |
| **Operational ToS risk** | L | L | L | **M–H** | **H** |

### Per-priority slug (loop 6)

| Slug | Tier | Reliability | Primary path |
|------|------|-------------|--------------|
| `ryanair`, `vueling`, `volotea`, `binter` | live | H for search | **CLI**; MCP supplement for non-LCC legs |
| `travelodge`, `hilton` | live | H | **CLI** UK; MCP for generic city search |
| `barcelo`, `h10`, `hotusa`, `eurostars`, … | live | M–H | **CLI** |
| `melia`, `nh`, `iberostar`, `marriott` | partial | L–M (Akamai) | **CLI** after `session chrome`; MCP for discovery only |
| `easyjet`, `aireuropa`, `iberiaexpress` | partial | L–M | **CLI**; Amadeus MCP overlap for Air Europa long-haul |
| Unimplemented scaffolds | stub | — | Neither; do not call |

### Failure modes and fallback

| Symptom | First fallback | Second fallback |
|---------|----------------|-------------------|
| MCP 401 / quota | Check env key | CLI if brand known |
| MCP empty offers | CLI for named LCC / chain | Widen dates / airports |
| CLI 403 / Akamai | `session chrome --wait` | MCP discovery + manual link |
| CLI `flights: null` | Treat as bug; expect `[]` | Retry or alternate slug |
| Both empty | Report transparently; do not invent prices |

---

## Implementation roadmap (loop 7+)

1. **MCP gateway package** — thin adapters: MCP JSON → `travelkit/types` (new `travelkit/mcp/` or separate repo).
2. **Brand registry** — extend `scripts/groups.json` with `mcp_eligible: false` and `preferred_tool: cli|mcp|hybrid`.
3. **Orchestrator skill** — document env vars: `DUFFEL_API_KEY`, `AMADEUS_CLIENT_ID`, `BOOKING_PARTNER_ID`.
4. **Do not delete CLIs** — MCP narrows the **discovery** surface; CLI remains source of truth for **Spanish chains and LCCs**.

---

## Related docs

- [AGENTS.md](../AGENTS.md) — MCP-first agent guidance
- [README.md](../README.md) — priority CLI status table
- [LOOP_STATUS.md](LOOP_STATUS.md) — live vs partial counts
- [QA_AIRLINES.md](QA_AIRLINES.md) / [QA_HOTELS_ES.md](QA_HOTELS_ES.md) — smoke reliability evidence
