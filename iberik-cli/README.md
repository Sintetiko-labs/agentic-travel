# Iberik CLI

Unofficial, agent-friendly CLI for [Iberik](https://www.iberikhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberik ./cmd/iberik
```

## Commands

```bash
iberik search [--json] [--limit N] <destination>
iberik read [--json] <id|url>
iberik availability [--json] --check-in DATE --check-out DATE <hotel-id>
iberik brands
```

## Environment

- `IBERIK_COOKIE` — optional browser cookie when blocked
- `IBERIK_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
