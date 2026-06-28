# Air Arabia CLI

Unofficial, agent-friendly CLI for [Air Arabia](https://www.airarabia.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airarabia ./cmd/airarabia
```

## Commands

```bash
airarabia search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
airarabia read [--json] [--brand BRAND] <id|url>
airarabia brands
```

## Environment

- `AIRARABIA_COOKIE` — optional browser cookie when blocked
- `AIRARABIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- Air Arabia
- Air Arabia Maroc

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
