# Egyptair CLI

Unofficial, agent-friendly CLI for [Egyptair](https://www.egyptair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o egyptair ./cmd/egyptair
```

## Commands

```bash
egyptair search [--json] --from MAD --to BCN --depart 2026-07-01
egyptair read [--json] <id|url>
egyptair brands
```

## Environment

- `EGYPTAIR_COOKIE` — optional browser cookie when blocked
- `EGYPTAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
