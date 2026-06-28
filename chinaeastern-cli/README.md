# China Eastern CLI

Unofficial, agent-friendly CLI for [China Eastern](https://www.ceair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o chinaeastern ./cmd/chinaeastern
```

## Commands

```bash
chinaeastern search [--json] --from MAD --to BCN --depart 2026-07-01
chinaeastern read [--json] <id|url>
chinaeastern brands
```

## Environment

- `CHINAEASTERN_COOKIE` — optional browser cookie when blocked
- `CHINAEASTERN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
