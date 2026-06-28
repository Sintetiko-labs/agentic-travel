# Level CLI

Unofficial, agent-friendly CLI for [Level](https://www.flylevel.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o level ./cmd/level
```

## Commands

```bash
level search [--json] --from MAD --to BCN --depart 2026-07-01
level read [--json] <id|url>
level brands
```

## Environment

- `LEVEL_COOKIE` — optional browser cookie when blocked
- `LEVEL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
