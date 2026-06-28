# LATAM CLI

Unofficial, agent-friendly CLI for [LATAM](https://www.latam.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o latam ./cmd/latam
```

## Commands

```bash
latam search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
latam read [--json] [--brand BRAND] <id|url>
latam brands
```

## Environment

- `LATAM_COOKIE` — optional browser cookie when blocked
- `LATAM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- LATAM Airlines
- LATAM Brasil
- LATAM Chile

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
