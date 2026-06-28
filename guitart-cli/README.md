# Guitart CLI

Unofficial, agent-friendly CLI for [Guitart](https://www.guitarthotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o guitart ./cmd/guitart
```

## Commands

```bash
guitart search [--json] [--limit N] <destination>
guitart read [--json] <id|url>
guitart availability [--json] --check-in DATE --check-out DATE <hotel-id>
guitart brands
```

## Environment

- `GUITART_COOKIE` — optional browser cookie when blocked
- `GUITART_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
