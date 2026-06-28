# Radisson CLI

Unofficial, agent-friendly CLI for [Radisson](https://www.radissonhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o radisson ./cmd/radisson
```

## Commands

```bash
radisson search [--json] [--limit N] [--brand BRAND] <destination>
radisson read [--json] [--brand BRAND] <id|url>
radisson availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
radisson brands
```

## Environment

- `RADISSON_COOKIE` — optional browser cookie when blocked
- `RADISSON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Radisson booking API:

- Radisson Hotel Group
- Radisson Blu
- Radisson RED
- Radisson Collection
- Park Inn by Radisson

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
