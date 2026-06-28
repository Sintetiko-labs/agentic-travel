# Freebird Airlines CLI

Unofficial, agent-friendly CLI for [Freebird Airlines](https://www.freebirdairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o freebird ./cmd/freebird
```

## Commands

```bash
freebird search [--json] --from MAD --to BCN --depart 2026-07-01
freebird read [--json] <id|url>
freebird brands
```

## Environment

- `FREEBIRD_COOKIE` — optional browser cookie when blocked
- `FREEBIRD_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
