# Viva Hotels CLI

Unofficial, agent-friendly CLI for [Viva Hotels](https://www.vivahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o viva ./cmd/viva
```

## Commands

```bash
viva search [--json] [--limit N] <destination>
viva read [--json] <id|url>
viva availability [--json] --check-in DATE --check-out DATE <hotel-id>
viva brands
```

## Environment

- `VIVA_COOKIE` — optional browser cookie when blocked
- `VIVA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
