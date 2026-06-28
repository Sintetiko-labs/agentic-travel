# Hyatt CLI

Unofficial, agent-friendly CLI for [Hyatt](https://www.hyatt.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hyatt ./cmd/hyatt
```

## Commands

```bash
hyatt search [--json] [--limit N] [--brand BRAND] <destination>
hyatt read [--json] [--brand BRAND] <id|url>
hyatt availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
hyatt brands
```

## Environment

- `HYATT_COOKIE` — optional browser cookie when blocked
- `HYATT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Hyatt booking API:

- Hyatt
- Grand Hyatt
- Hyatt Regency
- Hyatt Centric
- Thompson Hotels
- Andaz
- Alua Hotels
- Dreams Resorts
- Secrets Resorts
- Zoëtry
- Inclusive Collection

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
