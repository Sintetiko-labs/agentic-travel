# easyHotel CLI

Unofficial, agent-friendly CLI for [easyHotel](https://www.easyhotel.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o easyhotel ./cmd/easyhotel
```

## Commands

```bash
easyhotel search [--json] [--limit N] <destination>
easyhotel read [--json] <id|url>
easyhotel availability [--json] --check-in DATE --check-out DATE <hotel-id>
easyhotel brands
```

## Environment

- `EASYHOTEL_COOKIE` — optional browser cookie when blocked
- `EASYHOTEL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
