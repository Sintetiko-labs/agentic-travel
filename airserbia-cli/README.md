# Air Serbia CLI

Unofficial, agent-friendly CLI for [Air Serbia](https://www.airserbia.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airserbia ./cmd/airserbia
```

## Commands

```bash
airserbia search [--json] --from MAD --to BCN --depart 2026-07-01
airserbia read [--json] <id|url>
airserbia brands
```

## Environment

- `AIRSERBIA_COOKIE` — optional browser cookie when blocked
- `AIRSERBIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
