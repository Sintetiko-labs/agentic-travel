# Travelodge CLI

Unofficial, agent-friendly CLI for [Travelodge](https://www.travelodge.co.uk).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o travelodge ./cmd/travelodge
```

## Commands

```bash
travelodge search [--json] [--limit N] <destination>
travelodge read [--json] <id|url>
travelodge availability [--json] --check-in DATE --check-out DATE <hotel-id>
travelodge brands
travelodge session chrome|sync|doctor
```

## Search (London / UK)

`search` calls the public `/api/v2/hotel` JSON API (same backend as travelodge.co.uk search results). Default stay dates: **2026-07-05 → 2026-07-06**. No session required in most cases.

```bash
travelodge search --json London
travelodge search --json "Central London" --limit 5
```

## Read

`read` fetches hotel detail from a page URL (JSON-LD). Use `hotel_url` from search output:

```bash
travelodge read --json /hotels/318/London-Covent-Garden-hotel
```

## Environment

- `TRAVELODGE_COOKIE` — optional browser cookie when blocked
- `TRAVELODGE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **live** — `/api/v2/hotel` JSON API (July 2026 dates by default)
Read: **partial** — JSON-LD on hotel page URLs
Availability: **scaffold**
