# Air Transat CLI

Unofficial, agent-friendly CLI for [Air Transat](https://www.airtransat.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airtransat ./cmd/airtransat
```

## Commands

```bash
airtransat search [--json] --from MAD --to BCN --depart 2026-07-01
airtransat read [--json] <id|url>
airtransat brands
```

## Environment

- `AIRTRANSAT_COOKIE` — optional browser cookie when blocked
- `AIRTRANSAT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
