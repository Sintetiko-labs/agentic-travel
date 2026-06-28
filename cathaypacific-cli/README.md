# Cathay Pacific CLI

Unofficial, agent-friendly CLI for [Cathay Pacific](https://www.cathaypacific.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o cathaypacific ./cmd/cathaypacific
```

## Commands

```bash
cathaypacific search [--json] --from MAD --to BCN --depart 2026-07-01
cathaypacific read [--json] <id|url>
cathaypacific brands
```

## Environment

- `CATHAYPACIFIC_COOKIE` — optional browser cookie when blocked
- `CATHAYPACIFIC_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
