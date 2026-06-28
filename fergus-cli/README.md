# Fergus CLI

Unofficial, agent-friendly CLI for [Fergus](https://www.fergushotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o fergus ./cmd/fergus
```

## Commands

```bash
fergus search [--json] [--limit N] <destination>
fergus read [--json] <id|url>
fergus availability [--json] --check-in DATE --check-out DATE <hotel-id>
fergus brands
```

## Environment

- `FERGUS_COOKIE` — optional browser cookie when blocked
- `FERGUS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
