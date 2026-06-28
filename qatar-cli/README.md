# Qatar Airways CLI

Unofficial, agent-friendly CLI for [Qatar Airways](https://www.qatarairways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o qatar ./cmd/qatar
```

## Commands

```bash
qatar search [--json] --from MAD --to BCN --depart 2026-07-01
qatar read [--json] <id|url>
qatar brands
```

## Environment

- `QATAR_COOKIE` — optional browser cookie when blocked
- `QATAR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
