# Vincci CLI

Unofficial, agent-friendly CLI for [Vincci](https://www.vinccihoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o vincci ./cmd/vincci
```

## Commands

```bash
vincci search [--json] [--limit N] <destination>
vincci read [--json] <id|url>
vincci availability [--json] --check-in DATE --check-out DATE <hotel-id>
vincci brands
```

## Environment

- `VINCCI_COOKIE` — optional browser cookie when blocked
- `VINCCI_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
