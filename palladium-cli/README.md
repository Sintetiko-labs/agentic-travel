# Palladium CLI

Unofficial, agent-friendly CLI for [Palladium](https://www.palladiumhotelgroup.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o palladium ./cmd/palladium
```

## Commands

```bash
palladium search [--json] [--limit N] [--brand BRAND] <destination>
palladium read [--json] [--brand BRAND] <id|url>
palladium availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
palladium brands
```

## Environment

- `PALLADIUM_COOKIE` — optional browser cookie when blocked
- `PALLADIUM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Palladium booking API:

- Palladium Hotel Group
- Ushuaïa Ibiza Beach Hotel
- Hard Rock Hotel Ibiza
- TRS Hotels
- Grand Palladium
- Palladium Hotels

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
