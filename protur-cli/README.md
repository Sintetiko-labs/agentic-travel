# Protur CLI

Unofficial, agent-friendly CLI for [Protur](https://www.protur-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o protur ./cmd/protur
```

## Commands

```bash
protur search [--json] [--limit N] <destination>
protur read [--json] <id|url>
protur availability [--json] --check-in DATE --check-out DATE <hotel-id>
protur brands
```

## Environment

- `PROTUR_COOKIE` — optional browser cookie when blocked
- `PROTUR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
