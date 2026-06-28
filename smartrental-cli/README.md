# SmartRental CLI

Unofficial, agent-friendly CLI for [SmartRental](https://www.smartrental.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o smartrental ./cmd/smartrental
```

## Commands

```bash
smartrental search [--json] [--limit N] <destination>
smartrental read [--json] <id|url>
smartrental availability [--json] --check-in DATE --check-out DATE <hotel-id>
smartrental brands
```

## Environment

- `SMARTRENTAL_COOKIE` — optional browser cookie when blocked
- `SMARTRENTAL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
