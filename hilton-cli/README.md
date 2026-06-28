# Hilton CLI

Unofficial, agent-friendly CLI for [Hilton](https://www.hilton.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hilton ./cmd/hilton
```

## Commands

```bash
hilton search [--json] [--limit N] [--brand BRAND] <destination>
hilton read [--json] [--brand BRAND] <id|url>
hilton availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
hilton brands
```

## Environment

- `HILTON_COOKIE` — optional browser cookie when blocked
- `HILTON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Hilton booking API:

- Hilton
- Hilton Hotels & Resorts
- Conrad
- Waldorf Astoria
- DoubleTree by Hilton
- Canopy by Hilton
- Curio Collection
- Hampton by Hilton

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
