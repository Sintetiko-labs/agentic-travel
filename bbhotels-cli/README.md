# B&B Hotels CLI

Unofficial, agent-friendly CLI for [B&B Hotels](https://www.hotel-bb.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o bbhotels ./cmd/bbhotels
```

## Commands

```bash
bbhotels search [--json] [--limit N] <destination>
bbhotels read [--json] <id|url>
bbhotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
bbhotels brands
```

## Environment

- `BBHOTELS_COOKIE` — optional browser cookie when blocked
- `BBHOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
