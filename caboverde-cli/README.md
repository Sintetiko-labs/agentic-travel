# Cabo Verde Airlines CLI

Unofficial, agent-friendly CLI for [Cabo Verde Airlines](https://www.caboverdeairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o caboverde ./cmd/caboverde
```

## Commands

```bash
caboverde search [--json] --from MAD --to BCN --depart 2026-07-01
caboverde read [--json] <id|url>
caboverde brands
```

## Environment

- `CABOVERDE_COOKIE` — optional browser cookie when blocked
- `CABOVERDE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
