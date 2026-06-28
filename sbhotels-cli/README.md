# SB Hotels CLI

Unofficial, agent-friendly CLI for [SB Hotels](https://www.sbhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o sbhotels ./cmd/sbhotels
```

## Commands

```bash
sbhotels search [--json] [--limit N] <destination>
sbhotels read [--json] <id|url>
sbhotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
sbhotels brands
```

## Environment

- `SBHOTELS_COOKIE` — optional browser cookie when blocked
- `SBHOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
