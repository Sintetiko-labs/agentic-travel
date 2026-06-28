# Condor CLI

Unofficial, agent-friendly CLI for [Condor](https://www.condor.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o condor ./cmd/condor
```

## Commands

```bash
condor search [--json] --from MAD --to BCN --depart 2026-07-01
condor read [--json] <id|url>
condor brands
```

## Environment

- `CONDOR_COOKIE` — optional browser cookie when blocked
- `CONDOR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
