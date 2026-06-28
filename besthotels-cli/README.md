# Best Hotels CLI

Unofficial, agent-friendly CLI for [Best Hotels](https://www.besthotels.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o besthotels ./cmd/besthotels
```

## Commands

```bash
besthotels search [--json] [--limit N] <destination>
besthotels read [--json] <id|url>
besthotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
besthotels brands
```

## Environment

- `BESTHOTELS_COOKIE` — optional browser cookie when blocked
- `BESTHOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
