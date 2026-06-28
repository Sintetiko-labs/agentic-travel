# Magic Costa Blanca CLI

Unofficial, agent-friendly CLI for [Magic Costa Blanca](https://www.magiccostablanca.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o magic ./cmd/magic
```

## Commands

```bash
magic search [--json] [--limit N] <destination>
magic read [--json] <id|url>
magic availability [--json] --check-in DATE --check-out DATE <hotel-id>
magic brands
```

## Environment

- `MAGIC_COOKIE` — optional browser cookie when blocked
- `MAGIC_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
