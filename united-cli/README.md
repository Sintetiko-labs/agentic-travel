# United Airlines CLI

Unofficial, agent-friendly CLI for [United Airlines](https://www.united.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o united ./cmd/united
```

## Commands

```bash
united search [--json] --from MAD --to BCN --depart 2026-07-01
united read [--json] <id|url>
united brands
```

## Environment

- `UNITED_COOKIE` — optional browser cookie when blocked
- `UNITED_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
