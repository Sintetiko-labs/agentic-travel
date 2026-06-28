# Zenit CLI

Unofficial, agent-friendly CLI for [Zenit](https://www.zenithoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o zenit ./cmd/zenit
```

## Commands

```bash
zenit search [--json] [--limit N] <destination>
zenit read [--json] <id|url>
zenit availability [--json] --check-in DATE --check-out DATE <hotel-id>
zenit brands
```

## Environment

- `ZENIT_COOKIE` — optional browser cookie when blocked
- `ZENIT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
