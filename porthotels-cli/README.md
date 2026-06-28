# Port Hotels CLI

Unofficial, agent-friendly CLI for [Port Hotels](https://www.porthotels.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o porthotels ./cmd/porthotels
```

## Commands

```bash
porthotels search [--json] [--limit N] <destination>
porthotels read [--json] <id|url>
porthotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
porthotels brands
```

## Environment

- `PORTHOTELS_COOKIE` — optional browser cookie when blocked
- `PORTHOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
