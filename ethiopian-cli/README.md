# Ethiopian Airlines CLI

Unofficial, agent-friendly CLI for [Ethiopian Airlines](https://www.ethiopianairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ethiopian ./cmd/ethiopian
```

## Commands

```bash
ethiopian search [--json] --from MAD --to BCN --depart 2026-07-01
ethiopian read [--json] <id|url>
ethiopian brands
```

## Environment

- `ETHIOPIAN_COOKIE` — optional browser cookie when blocked
- `ETHIOPIAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
