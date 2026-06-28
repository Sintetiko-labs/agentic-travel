# easyJet CLI

Unofficial, agent-friendly CLI for [easyJet](https://www.easyjet.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o easyjet ./cmd/easyjet
```

## Commands

```bash
easyjet search [--json] --from MAD --to BCN --depart 2026-07-01
easyjet read [--json] <id|url>
easyjet brands
```

## Environment

- `EASYJET_COOKIE` — optional browser cookie when blocked
- `EASYJET_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
