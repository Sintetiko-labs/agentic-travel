# Derby Hotels CLI

Unofficial, agent-friendly CLI for [Derby Hotels](https://www.derbyhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o derby ./cmd/derby
```

## Commands

```bash
derby search [--json] [--limit N] <destination>
derby read [--json] <id|url>
derby availability [--json] --check-in DATE --check-out DATE <hotel-id>
derby brands
```

## Environment

- `DERBY_COOKIE` — optional browser cookie when blocked
- `DERBY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
