# Air Canada CLI

Unofficial, agent-friendly CLI for [Air Canada](https://www.aircanada.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aircanada ./cmd/aircanada
```

## Commands

```bash
aircanada search [--json] --from MAD --to BCN --depart 2026-07-01
aircanada read [--json] <id|url>
aircanada brands
```

## Environment

- `AIRCANADA_COOKIE` — optional browser cookie when blocked
- `AIRCANADA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
