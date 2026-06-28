# Globales CLI

Unofficial, agent-friendly CLI for [Globales](https://www.globales.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o globales ./cmd/globales
```

## Commands

```bash
globales search [--json] [--limit N] [--brand BRAND] <destination>
globales read [--json] [--brand BRAND] <id|url>
globales availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
globales brands
```

## Environment

- `GLOBALES_COOKIE` — optional browser cookie when blocked
- `GLOBALES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Globales booking API:

- Globales Hotels
- Hoteles Globales

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
