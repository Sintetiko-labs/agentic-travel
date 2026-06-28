# WestJet CLI

Unofficial, agent-friendly CLI for [WestJet](https://www.westjet.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o westjet ./cmd/westjet
```

## Commands

```bash
westjet search [--json] --from MAD --to BCN --depart 2026-07-01
westjet read [--json] <id|url>
westjet brands
```

## Environment

- `WESTJET_COOKIE` — optional browser cookie when blocked
- `WESTJET_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
