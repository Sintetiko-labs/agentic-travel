# Wizz Air CLI

Unofficial, agent-friendly CLI for [Wizz Air](https://wizzair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o wizzair ./cmd/wizzair
```

## Commands

```bash
wizzair search [--json] --from MAD --to BCN --depart 2026-07-01
wizzair read [--json] <id|url>
wizzair brands
```

## Environment

- `WIZZAIR_COOKIE` — optional browser cookie when blocked
- `WIZZAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
