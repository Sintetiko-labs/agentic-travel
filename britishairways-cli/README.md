# British Airways CLI

Unofficial, agent-friendly CLI for [British Airways](https://www.britishairways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o britishairways ./cmd/britishairways
```

## Commands

```bash
britishairways search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
britishairways read [--json] [--brand BRAND] <id|url>
britishairways brands
```

## Environment

- `BRITISHAIRWAYS_COOKIE` — optional browser cookie when blocked
- `BRITISHAIRWAYS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- British Airways
- BA CityFlyer

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
