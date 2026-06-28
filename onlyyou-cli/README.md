# Only YOU CLI

Unofficial, agent-friendly CLI for [Only YOU](https://www.onlyyouhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o onlyyou ./cmd/onlyyou
```

## Commands

```bash
onlyyou search [--json] [--limit N] <destination>
onlyyou read [--json] <id|url>
onlyyou availability [--json] --check-in DATE --check-out DATE <hotel-id>
onlyyou brands
```

## Environment

- `ONLYYOU_COOKIE` — optional browser cookie when blocked
- `ONLYYOU_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
