# Air Algérie CLI

Unofficial, agent-friendly CLI for [Air Algérie](https://www.airalgerie.dz).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airalgerie ./cmd/airalgerie
```

## Commands

```bash
airalgerie search [--json] --from MAD --to BCN --depart 2026-07-01
airalgerie read [--json] <id|url>
airalgerie brands
```

## Environment

- `AIRALGERIE_COOKIE` — optional browser cookie when blocked
- `AIRALGERIE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
