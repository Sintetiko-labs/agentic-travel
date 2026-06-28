# Avianca CLI

Unofficial, agent-friendly CLI for [Avianca](https://www.avianca.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o avianca ./cmd/avianca
```

## Commands

```bash
avianca search [--json] --from MAD --to BCN --depart 2026-07-01
avianca read [--json] <id|url>
avianca brands
```

## Environment

- `AVIANCA_COOKIE` — optional browser cookie when blocked
- `AVIANCA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
