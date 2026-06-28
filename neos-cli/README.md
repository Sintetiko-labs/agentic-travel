# Neos CLI

Unofficial, agent-friendly CLI for [Neos](https://www.neosair.it).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o neos ./cmd/neos
```

## Commands

```bash
neos search [--json] --from MAD --to BCN --depart 2026-07-01
neos read [--json] <id|url>
neos brands
```

## Environment

- `NEOS_COOKIE` — optional browser cookie when blocked
- `NEOS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
