# Ona Hotels CLI

Unofficial, agent-friendly CLI for [Ona Hotels](https://www.onahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ona ./cmd/ona
```

## Commands

```bash
ona search [--json] [--limit N] [--brand BRAND] <destination>
ona read [--json] [--brand BRAND] <id|url>
ona availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
ona brands
```

## Environment

- `ONA_COOKIE` — optional browser cookie when blocked
- `ONA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Ona Hotels booking API:

- Ona Hotels
- Ona Hotels & Apartments

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
