# Norwegian CLI

Unofficial, agent-friendly CLI for [Norwegian](https://www.norwegian.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o norwegian ./cmd/norwegian
```

## Commands

```bash
norwegian search [--json] --from MAD --to BCN --depart 2026-07-01
norwegian read [--json] <id|url>
norwegian brands
```

## Environment

- `NORWEGIAN_COOKIE` — optional browser cookie when blocked
- `NORWEGIAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
