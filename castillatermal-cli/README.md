# Castilla Termal CLI

Unofficial, agent-friendly CLI for [Castilla Termal](https://www.castillatermal.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o castillatermal ./cmd/castillatermal
```

## Commands

```bash
castillatermal search [--json] [--limit N] <destination>
castillatermal read [--json] <id|url>
castillatermal availability [--json] --check-in DATE --check-out DATE <hotel-id>
castillatermal brands
```

## Environment

- `CASTILLATERMAL_COOKIE` — optional browser cookie when blocked
- `CASTILLATERMAL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
