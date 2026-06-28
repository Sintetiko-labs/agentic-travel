# Travelodge CLI

Unofficial, agent-friendly CLI for [Travelodge](https://www.travelodge.co.uk).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o travelodge ./cmd/travelodge
```

## Commands

```bash
travelodge search [--json] [--limit N] <destination>
travelodge read [--json] <id|url>
travelodge availability [--json] --check-in DATE --check-out DATE <hotel-id>
travelodge brands
```

## Environment

- `TRAVELODGE_COOKIE` — optional browser cookie when blocked
- `TRAVELODGE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
