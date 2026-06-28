# Icelandair CLI

Unofficial, agent-friendly CLI for [Icelandair](https://www.icelandair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o icelandair ./cmd/icelandair
```

## Commands

```bash
icelandair search [--json] --from MAD --to BCN --depart 2026-07-01
icelandair read [--json] <id|url>
icelandair brands
```

## Environment

- `ICELANDAIR_COOKIE` — optional browser cookie when blocked
- `ICELANDAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
