# Ryanair CLI

Unofficial, agent-friendly CLI for [Ryanair](https://www.ryanair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ryanair ./cmd/ryanair
```

## Commands

```bash
ryanair search [--json] --from MAD --to BCN --depart 2026-07-01
ryanair read [--json] <id|url>
ryanair brands
```

## Environment

- `RYANAIR_COOKIE` — optional browser cookie when blocked
- `RYANAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
