# Axel Hotels CLI

Unofficial, agent-friendly CLI for [Axel Hotels](https://www.axelhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o axel ./cmd/axel
```

## Commands

```bash
axel search [--json] [--limit N] <destination>
axel read [--json] <id|url>
axel availability [--json] --check-in DATE --check-out DATE <hotel-id>
axel brands
```

## Environment

- `AXEL_COOKIE` — optional browser cookie when blocked
- `AXEL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
