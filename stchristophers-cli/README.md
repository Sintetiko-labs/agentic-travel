# St Christopher's CLI

Unofficial, agent-friendly CLI for [St Christopher's](https://www.st-christophers.co.uk).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o stchristophers ./cmd/stchristophers
```

## Commands

```bash
stchristophers search [--json] [--limit N] <destination>
stchristophers read [--json] <id|url>
stchristophers availability [--json] --check-in DATE --check-out DATE <hotel-id>
stchristophers brands
```

## Environment

- `STCHRISTOPHERS_COOKIE` — optional browser cookie when blocked
- `STCHRISTOPHERS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
