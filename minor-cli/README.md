# Minor Hotels CLI

Unofficial, agent-friendly CLI for [Minor Hotels](https://www.minorhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o minor ./cmd/minor
```

## Commands

```bash
minor search [--json] [--limit N] [--brand BRAND] <destination>
minor read [--json] [--brand BRAND] <id|url>
minor availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
minor brands
```

## Environment

- `MINOR_COOKIE` — optional browser cookie when blocked
- `MINOR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Minor Hotels booking API:

- Avani
- Tivoli
- Minor Hotels

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
