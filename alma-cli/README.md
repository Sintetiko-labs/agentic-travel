# Alma Hotels CLI

Unofficial, agent-friendly CLI for [Alma Hotels](https://www.almahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o alma ./cmd/alma
```

## Commands

```bash
alma search [--json] [--limit N] <destination>
alma read [--json] <id|url>
alma availability [--json] --check-in DATE --check-out DATE <hotel-id>
alma brands
```

## Environment

- `ALMA_COOKIE` — optional browser cookie when blocked
- `ALMA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
