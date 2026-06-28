# Grupotel CLI

Unofficial, agent-friendly CLI for [Grupotel](https://www.grupotel.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o grupotel ./cmd/grupotel
```

## Commands

```bash
grupotel search [--json] [--limit N] <destination>
grupotel read [--json] <id|url>
grupotel availability [--json] --check-in DATE --check-out DATE <hotel-id>
grupotel brands
```

## Environment

- `GRUPOTEL_COOKIE` — optional browser cookie when blocked
- `GRUPOTEL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
