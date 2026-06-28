# Zoku CLI

Unofficial, agent-friendly CLI for [Zoku](https://www.livezoku.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o zoku ./cmd/zoku
```

## Commands

```bash
zoku search [--json] [--limit N] <destination>
zoku read [--json] <id|url>
zoku availability [--json] --check-in DATE --check-out DATE <hotel-id>
zoku brands
```

## Environment

- `ZOKU_COOKIE` — optional browser cookie when blocked
- `ZOKU_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
