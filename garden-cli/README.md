# Garden Hotels CLI

Unofficial, agent-friendly CLI for [Garden Hotels](https://www.gardenhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o garden ./cmd/garden
```

## Commands

```bash
garden search [--json] [--limit N] <destination>
garden read [--json] <id|url>
garden availability [--json] --check-in DATE --check-out DATE <hotel-id>
garden brands
```

## Environment

- `GARDEN_COOKIE` — optional browser cookie when blocked
- `GARDEN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
