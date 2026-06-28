# Hoteles Santos CLI

Unofficial, agent-friendly CLI for [Hoteles Santos](https://www.hoteles-santos.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o santos ./cmd/santos
```

## Commands

```bash
santos search [--json] [--limit N] <destination>
santos read [--json] <id|url>
santos availability [--json] --check-in DATE --check-out DATE <hotel-id>
santos brands
```

## Environment

- `SANTOS_COOKIE` — optional browser cookie when blocked
- `SANTOS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
