# World2Fly CLI

Unofficial, agent-friendly CLI for [World2Fly](https://www.world2fly.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o world2fly ./cmd/world2fly
```

## Commands

```bash
world2fly search [--json] --from MAD --to BCN --depart 2026-07-01
world2fly read [--json] <id|url>
world2fly brands
```

## Environment

- `WORLD2FLY_COOKIE` — optional browser cookie when blocked
- `WORLD2FLY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
