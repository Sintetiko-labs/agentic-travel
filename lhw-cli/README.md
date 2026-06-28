# Leading Hotels CLI

Unofficial, agent-friendly CLI for [Leading Hotels](https://www.lhw.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o lhw ./cmd/lhw
```

## Commands

```bash
lhw search [--json] [--limit N] <destination>
lhw read [--json] <id|url>
lhw availability [--json] --check-in DATE --check-out DATE <hotel-id>
lhw brands
```

## Environment

- `LHW_COOKIE` — optional browser cookie when blocked
- `LHW_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
