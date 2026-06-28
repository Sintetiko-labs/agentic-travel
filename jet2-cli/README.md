# Jet2 CLI

Unofficial, agent-friendly CLI for [Jet2](https://www.jet2.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o jet2 ./cmd/jet2
```

## Commands

```bash
jet2 search [--json] --from MAD --to BCN --depart 2026-07-01
jet2 read [--json] <id|url>
jet2 brands
```

## Environment

- `JET2_COOKIE` — optional browser cookie when blocked
- `JET2_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
