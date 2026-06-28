# Small Luxury Hotels CLI

Unofficial, agent-friendly CLI for [Small Luxury Hotels](https://www.slh.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o slh ./cmd/slh
```

## Commands

```bash
slh search [--json] [--limit N] <destination>
slh read [--json] <id|url>
slh availability [--json] --check-in DATE --check-out DATE <hotel-id>
slh brands
```

## Environment

- `SLH_COOKIE` — optional browser cookie when blocked
- `SLH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
