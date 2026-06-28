# Hipotels CLI

Unofficial, agent-friendly CLI for [Hipotels](https://www.hipotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hipotels ./cmd/hipotels
```

## Commands

```bash
hipotels search [--json] [--limit N] <destination>
hipotels read [--json] <id|url>
hipotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
hipotels brands
```

## Environment

- `HIPOTELS_COOKIE` — optional browser cookie when blocked
- `HIPOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
