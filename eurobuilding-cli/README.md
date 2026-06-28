# Eurobuilding CLI

Unofficial, agent-friendly CLI for [Eurobuilding](https://www.eurobuilding.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o eurobuilding ./cmd/eurobuilding
```

## Commands

```bash
eurobuilding search [--json] [--limit N] <destination>
eurobuilding read [--json] <id|url>
eurobuilding availability [--json] --check-in DATE --check-out DATE <hotel-id>
eurobuilding brands
```

## Environment

- `EUROBUILDING_COOKIE` — optional browser cookie when blocked
- `EUROBUILDING_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
