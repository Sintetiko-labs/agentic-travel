# Smartwings CLI

Unofficial, agent-friendly CLI for [Smartwings](https://www.smartwings.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o smartwings ./cmd/smartwings
```

## Commands

```bash
smartwings search [--json] --from MAD --to BCN --depart 2026-07-01
smartwings read [--json] <id|url>
smartwings brands
```

## Environment

- `SMARTWINGS_COOKIE` — optional browser cookie when blocked
- `SMARTWINGS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
