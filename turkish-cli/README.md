# Turkish Airlines Group CLI

Unofficial, agent-friendly CLI for [Turkish Airlines Group](https://www.turkishairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o turkish ./cmd/turkish
```

## Commands

```bash
turkish search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
turkish read [--json] [--brand BRAND] <id|url>
turkish brands
```

## Environment

- `TURKISH_COOKIE` — optional browser cookie when blocked
- `TURKISH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- Turkish Airlines
- Pegasus Airlines
- SunExpress
- AnadoluJet / AJet

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
