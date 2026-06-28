# Thai Airways CLI

Unofficial, agent-friendly CLI for [Thai Airways](https://www.thaiairways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o thaiairways ./cmd/thaiairways
```

## Commands

```bash
thaiairways search [--json] --from MAD --to BCN --depart 2026-07-01
thaiairways read [--json] <id|url>
thaiairways brands
```

## Environment

- `THAIAIRWAYS_COOKIE` — optional browser cookie when blocked
- `THAIAIRWAYS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
