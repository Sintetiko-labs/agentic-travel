# Kuwait Airways CLI

Unofficial, agent-friendly CLI for [Kuwait Airways](https://www.kuwaitairways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o kuwaitairways ./cmd/kuwaitairways
```

## Commands

```bash
kuwaitairways search [--json] --from MAD --to BCN --depart 2026-07-01
kuwaitairways read [--json] <id|url>
kuwaitairways brands
```

## Environment

- `KUWAITAIRWAYS_COOKIE` — optional browser cookie when blocked
- `KUWAITAIRWAYS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
