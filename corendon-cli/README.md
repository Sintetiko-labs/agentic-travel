# Corendon Airlines CLI

Unofficial, agent-friendly CLI for [Corendon Airlines](https://www.corendonairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o corendon ./cmd/corendon
```

## Commands

```bash
corendon search [--json] --from MAD --to BCN --depart 2026-07-01
corendon read [--json] <id|url>
corendon brands
```

## Environment

- `CORENDON_COOKIE` — optional browser cookie when blocked
- `CORENDON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
