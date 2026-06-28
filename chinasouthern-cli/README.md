# China Southern CLI

Unofficial, agent-friendly CLI for [China Southern](https://www.csair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o chinasouthern ./cmd/chinasouthern
```

## Commands

```bash
chinasouthern search [--json] --from MAD --to BCN --depart 2026-07-01
chinasouthern read [--json] <id|url>
chinasouthern brands
```

## Environment

- `CHINASOUTHERN_COOKIE` — optional browser cookie when blocked
- `CHINASOUTHERN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
