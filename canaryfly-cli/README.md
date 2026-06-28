# Canaryfly CLI

Unofficial, agent-friendly CLI for [Canaryfly](https://www.canaryfly.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o canaryfly ./cmd/canaryfly
```

## Commands

```bash
canaryfly search [--json] --from MAD --to BCN --depart 2026-07-01
canaryfly read [--json] <id|url>
canaryfly brands
```

## Environment

- `CANARYFLY_COOKIE` — optional browser cookie when blocked
- `CANARYFLY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
