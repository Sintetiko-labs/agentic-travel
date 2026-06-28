# Iberojet CLI

Unofficial, agent-friendly CLI for [Iberojet](https://www.iberojet.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberojet ./cmd/iberojet
```

## Commands

```bash
iberojet search [--json] --from MAD --to BCN --depart 2026-07-01
iberojet read [--json] <id|url>
iberojet brands
```

## Environment

- `IBEROJET_COOKIE` — optional browser cookie when blocked
- `IBEROJET_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
