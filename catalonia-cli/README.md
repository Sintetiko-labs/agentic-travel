# Catalonia Hotels CLI

Unofficial, agent-friendly CLI for [Catalonia Hotels](https://www.cataloniahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o catalonia ./cmd/catalonia
```

## Commands

```bash
catalonia search [--json] [--limit N] <destination>
catalonia read [--json] <id|url>
catalonia availability [--json] --check-in DATE --check-out DATE <hotel-id>
catalonia brands
```

## Environment

- `CATALONIA_COOKIE` — optional browser cookie when blocked
- `CATALONIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
