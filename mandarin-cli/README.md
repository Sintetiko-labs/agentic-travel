# Mandarin Oriental CLI

Unofficial, agent-friendly CLI for [Mandarin Oriental](https://www.mandarinoriental.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o mandarin ./cmd/mandarin
```

## Commands

```bash
mandarin search [--json] [--limit N] <destination>
mandarin read [--json] <id|url>
mandarin availability [--json] --check-in DATE --check-out DATE <hotel-id>
mandarin brands
```

## Environment

- `MANDARIN_COOKIE` — optional browser cookie when blocked
- `MANDARIN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
