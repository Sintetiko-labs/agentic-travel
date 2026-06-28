# Hoteles RH CLI

Unofficial, agent-friendly CLI for [Hoteles RH](https://www.rhhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o rh ./cmd/rh
```

## Commands

```bash
rh search [--json] [--limit N] <destination>
rh read [--json] <id|url>
rh availability [--json] --check-in DATE --check-out DATE <hotel-id>
rh brands
```

## Environment

- `RH_COOKIE` — optional browser cookie when blocked
- `RH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
