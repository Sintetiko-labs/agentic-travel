# Wyndham CLI

Unofficial, agent-friendly CLI for [Wyndham](https://www.wyndhamhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o wyndham ./cmd/wyndham
```

## Commands

```bash
wyndham search [--json] [--limit N] [--brand BRAND] <destination>
wyndham read [--json] [--brand BRAND] <id|url>
wyndham availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
wyndham brands
```

## Environment

- `WYNDHAM_COOKIE` — optional browser cookie when blocked
- `WYNDHAM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Wyndham booking API:

- Wyndham Hotels & Resorts
- Ramada
- Wyndham
- Tryp
- Dolce by Wyndham

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
