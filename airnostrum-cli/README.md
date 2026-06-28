# Air Nostrum CLI

Unofficial, agent-friendly CLI for [Air Nostrum](https://www.airnostrum.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airnostrum ./cmd/airnostrum
```

## Commands

```bash
airnostrum search [--json] --from MAD --to BCN --depart 2026-07-01
airnostrum read [--json] <id|url>
airnostrum brands
```

## Environment

- `AIRNOSTRUM_COOKIE` — optional browser cookie when blocked
- `AIRNOSTRUM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
