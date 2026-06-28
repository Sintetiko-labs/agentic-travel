# Paradores CLI

Unofficial, agent-friendly CLI for [Paradores](https://www.parador.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o paradores ./cmd/paradores
```

## Commands

```bash
paradores search [--json] [--limit N] <destination>
paradores read [--json] <id|url>
paradores availability [--json] --check-in DATE --check-out DATE <hotel-id>
paradores brands
```

## Environment

- `PARADORES_COOKIE` — optional browser cookie when blocked
- `PARADORES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
