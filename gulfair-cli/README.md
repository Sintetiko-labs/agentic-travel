# Gulf Air CLI

Unofficial, agent-friendly CLI for [Gulf Air](https://www.gulfair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o gulfair ./cmd/gulfair
```

## Commands

```bash
gulfair search [--json] --from MAD --to BCN --depart 2026-07-01
gulfair read [--json] <id|url>
gulfair brands
```

## Environment

- `GULFAIR_COOKIE` — optional browser cookie when blocked
- `GULFAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
