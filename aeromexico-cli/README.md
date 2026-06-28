# Aeroméxico CLI

Unofficial, agent-friendly CLI for [Aeroméxico](https://www.aeromexico.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aeromexico ./cmd/aeromexico
```

## Commands

```bash
aeromexico search [--json] --from MAD --to BCN --depart 2026-07-01
aeromexico read [--json] <id|url>
aeromexico brands
```

## Environment

- `AEROMEXICO_COOKIE` — optional browser cookie when blocked
- `AEROMEXICO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
