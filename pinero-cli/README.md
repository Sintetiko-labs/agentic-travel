# Grupo Piñero CLI

Unofficial, agent-friendly CLI for [Grupo Piñero](https://www.bahia-principe.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o pinero ./cmd/pinero
```

## Commands

```bash
pinero search [--json] [--limit N] [--brand BRAND] <destination>
pinero read [--json] [--brand BRAND] <id|url>
pinero availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
pinero brands
```

## Environment

- `PINERO_COOKIE` — optional browser cookie when blocked
- `PINERO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Grupo Piñero booking API:

- Fiesta Hotels & Resorts
- Grupo Piñero
- Bahia Principe

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
