# Pierre & Vacances CLI

Unofficial, agent-friendly CLI for [Pierre & Vacances](https://www.pierreetvacances.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o pierrevacances ./cmd/pierrevacances
```

## Commands

```bash
pierrevacances search [--json] [--limit N] [--brand BRAND] <destination>
pierrevacances read [--json] [--brand BRAND] <id|url>
pierrevacances availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
pierrevacances brands
```

## Environment

- `PIERREVACANCES_COOKIE` — optional browser cookie when blocked
- `PIERREVACANCES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Pierre & Vacances booking API:

- Pierre & Vacances
- Pierre & Vacances España
- Apartamentos Pierre & Vacances

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
