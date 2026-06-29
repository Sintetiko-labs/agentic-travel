# Iberia CLI

Unofficial, agent-friendly CLI for [Iberia](https://www.iberia.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberia ./cmd/iberia
```

## Commands

```bash
iberia search [--json] --from MAD --to BCN --depart 2026-07-01
iberia read [--json] <id|url>
iberia brands
```

## Environment

- `AIRNOSTRUM_COOKIE` — optional browser cookie when blocked
- `AIRNOSTRUM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
