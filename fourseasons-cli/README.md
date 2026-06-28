# Four Seasons CLI

Unofficial, agent-friendly CLI for [Four Seasons](https://www.fourseasons.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o fourseasons ./cmd/fourseasons
```

## Commands

```bash
fourseasons search [--json] [--limit N] <destination>
fourseasons read [--json] <id|url>
fourseasons availability [--json] --check-in DATE --check-out DATE <hotel-id>
fourseasons brands
```

## Environment

- `FOURSEASONS_COOKIE` — optional browser cookie when blocked
- `FOURSEASONS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
