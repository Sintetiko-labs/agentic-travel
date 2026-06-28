# Saudia CLI

Unofficial, agent-friendly CLI for [Saudia](https://www.saudia.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o saudia ./cmd/saudia
```

## Commands

```bash
saudia search [--json] --from MAD --to BCN --depart 2026-07-01
saudia read [--json] <id|url>
saudia brands
```

## Environment

- `SAUDIA_COOKIE` — optional browser cookie when blocked
- `SAUDIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
