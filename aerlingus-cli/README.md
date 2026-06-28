# Aer Lingus CLI

Unofficial, agent-friendly CLI for [Aer Lingus](https://www.aerlingus.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aerlingus ./cmd/aerlingus
```

## Commands

```bash
aerlingus search [--json] --from MAD --to BCN --depart 2026-07-01
aerlingus read [--json] <id|url>
aerlingus brands
```

## Environment

- `AERLINGUS_COOKIE` — optional browser cookie when blocked
- `AERLINGUS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
