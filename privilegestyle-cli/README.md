# Privilege Style CLI

Unofficial, agent-friendly CLI for [Privilege Style](https://www.privilegestyle.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o privilegestyle ./cmd/privilegestyle
```

## Commands

```bash
privilegestyle search [--json] --from MAD --to BCN --depart 2026-07-01
privilegestyle read [--json] <id|url>
privilegestyle brands
```

## Environment

- `PRIVILEGESTYLE_COOKIE` — optional browser cookie when blocked
- `PRIVILEGESTYLE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
