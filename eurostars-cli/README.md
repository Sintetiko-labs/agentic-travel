# Eurostars CLI

Unofficial, agent-friendly CLI for [Eurostars](https://www.eurostarshotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o eurostars ./cmd/eurostars
```

## Commands

```bash
eurostars search [--json] [--limit N] [--brand BRAND] <destination>
eurostars read [--json] [--brand BRAND] <id|url>
eurostars availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
eurostars brands
```

## Environment

- `EUROSTARS_COOKIE` — optional browser cookie when blocked
- `EUROSTARS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Eurostars booking API:

- Eurostars Hotel Company
- Eurostars Hotels
- Exe Hotels
- Ikonik Hotels
- Áurea Hotels
- Tandem Suites

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
