# El Al CLI

Unofficial, agent-friendly CLI for [El Al](https://www.elal.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o elal ./cmd/elal
```

## Commands

```bash
elal search [--json] --from MAD --to BCN --depart 2026-07-01
elal read [--json] <id|url>
elal brands
```

## Environment

- `ELAL_COOKIE` — optional browser cookie when blocked
- `ELAL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
