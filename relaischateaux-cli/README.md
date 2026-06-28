# Relais & Châteaux CLI

Unofficial, agent-friendly CLI for [Relais & Châteaux](https://www.relaischateaux.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o relaischateaux ./cmd/relaischateaux
```

## Commands

```bash
relaischateaux search [--json] [--limit N] <destination>
relaischateaux read [--json] <id|url>
relaischateaux availability [--json] --check-in DATE --check-out DATE <hotel-id>
relaischateaux brands
```

## Environment

- `RELAISCHATEAUX_COOKIE` — optional browser cookie when blocked
- `RELAISCHATEAUX_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
