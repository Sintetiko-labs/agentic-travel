# Hoteles Elba CLI

Unofficial, agent-friendly CLI for [Hoteles Elba](https://www.hoteleselba.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o elba ./cmd/elba
```

## Commands

```bash
elba search [--json] [--limit N] <destination>
elba read [--json] <id|url>
elba availability [--json] --check-in DATE --check-out DATE <hotel-id>
elba brands
```

## Environment

- `ELBA_COOKIE` — optional browser cookie when blocked
- `ELBA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
