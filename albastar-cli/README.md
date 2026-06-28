# Albastar CLI

Unofficial, agent-friendly CLI for [Albastar](https://www.albastar.es).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o albastar ./cmd/albastar
```

## Commands

```bash
albastar search [--json] --from MAD --to BCN --depart 2026-07-01
albastar read [--json] <id|url>
albastar brands
```

## Environment

- `ALBASTAR_COOKIE` — optional browser cookie when blocked
- `ALBASTAR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
