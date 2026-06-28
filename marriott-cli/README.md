# Marriott CLI

Unofficial, agent-friendly CLI for [Marriott](https://www.marriott.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o marriott ./cmd/marriott
```

## Commands

```bash
marriott search [--json] [--limit N] [--brand BRAND] <destination>
marriott read [--json] [--brand BRAND] <id|url>
marriott availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
marriott brands
```

## Environment

- `MARRIOTT_COOKIE` — optional browser cookie when blocked
- `MARRIOTT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Marriott booking API:

- Marriott
- Marriott Hotels
- JW Marriott
- The Ritz-Carlton
- St. Regis
- W Hotels
- Edition
- Luxury Collection
- Westin
- Sheraton
- Le Méridien
- Renaissance Hotels
- Autograph Collection
- Tribute Portfolio
- AC Hotels
- AC Hotels by Marriott
- Aloft
- Moxy
- Courtyard by Marriott
- Residence Inn

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
