# Binter CLI

Unofficial, agent-friendly CLI for [Binter](https://www.bintercanarias.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o binter ./cmd/binter
```

## Commands

```bash
binter search [--json] --from MAD --to BCN --depart 2026-07-01
binter read [--json] <id|url>
binter brands
```

## Environment

- `BINTER_COOKIE` — optional browser cookie when blocked
- `BINTER_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
