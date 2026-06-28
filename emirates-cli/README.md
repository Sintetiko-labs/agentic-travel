# Emirates CLI

Unofficial, agent-friendly CLI for [Emirates](https://www.emirates.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o emirates ./cmd/emirates
```

## Commands

```bash
emirates search [--json] --from MAD --to BCN --depart 2026-07-01
emirates read [--json] <id|url>
emirates brands
```

## Environment

- `EMIRATES_COOKIE` — optional browser cookie when blocked
- `EMIRATES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
