# Delta Air Lines CLI

Unofficial, agent-friendly CLI for [Delta Air Lines](https://www.delta.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o delta ./cmd/delta
```

## Commands

```bash
delta search [--json] --from MAD --to BCN --depart 2026-07-01
delta read [--json] <id|url>
delta brands
```

## Environment

- `DELTA_COOKIE` — optional browser cookie when blocked
- `DELTA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
