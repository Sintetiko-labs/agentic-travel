# Numa CLI

Unofficial, agent-friendly CLI for [Numa](https://www.numastays.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o numa ./cmd/numa
```

## Commands

```bash
numa search [--json] [--limit N] <destination>
numa read [--json] <id|url>
numa availability [--json] --check-in DATE --check-out DATE <hotel-id>
numa brands
```

## Environment

- `NUMA_COOKIE` — optional browser cookie when blocked
- `NUMA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
