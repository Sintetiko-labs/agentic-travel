# Tent Hotels CLI

Unofficial, agent-friendly CLI for [Tent Hotels](https://www.tenthotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o tent ./cmd/tent
```

## Commands

```bash
tent search [--json] [--limit N] <destination>
tent read [--json] <id|url>
tent availability [--json] --check-in DATE --check-out DATE <hotel-id>
tent brands
```

## Environment

- `TENT_COOKIE` — optional browser cookie when blocked
- `TENT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
