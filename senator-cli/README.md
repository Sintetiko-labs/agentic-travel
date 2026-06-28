# Senator CLI

Unofficial, agent-friendly CLI for [Senator](https://www.senator.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o senator ./cmd/senator
```

## Commands

```bash
senator search [--json] [--limit N] [--brand BRAND] <destination>
senator read [--json] [--brand BRAND] <id|url>
senator availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
senator brands
```

## Environment

- `SENATOR_COOKIE` — optional browser cookie when blocked
- `SENATOR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Senator booking API:

- Senator Hotels & Resorts
- Playa Senator

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
