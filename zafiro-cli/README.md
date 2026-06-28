# Zafiro CLI

Unofficial, agent-friendly CLI for [Zafiro](https://www.zafirohotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o zafiro ./cmd/zafiro
```

## Commands

```bash
zafiro search [--json] [--limit N] <destination>
zafiro read [--json] <id|url>
zafiro availability [--json] --check-in DATE --check-out DATE <hotel-id>
zafiro brands
```

## Environment

- `ZAFIRO_COOKIE` — optional browser cookie when blocked
- `ZAFIRO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
