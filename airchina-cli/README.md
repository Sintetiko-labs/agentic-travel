# Air China CLI

Unofficial, agent-friendly CLI for [Air China](https://www.airchina.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airchina ./cmd/airchina
```

## Commands

```bash
airchina search [--json] --from MAD --to BCN --depart 2026-07-01
airchina read [--json] <id|url>
airchina brands
```

## Environment

- `AIRCHINA_COOKIE` — optional browser cookie when blocked
- `AIRCHINA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
