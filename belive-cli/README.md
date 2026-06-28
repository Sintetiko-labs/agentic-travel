# Be Live CLI

Unofficial, agent-friendly CLI for [Be Live](https://www.belivehotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o belive ./cmd/belive
```

## Commands

```bash
belive search [--json] [--limit N] <destination>
belive read [--json] <id|url>
belive availability [--json] --check-in DATE --check-out DATE <hotel-id>
belive brands
```

## Environment

- `BELIVE_COOKIE` — optional browser cookie when blocked
- `BELIVE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
