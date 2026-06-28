# Air France-KLM CLI

Unofficial, agent-friendly CLI for [Air France-KLM](https://www.airfrance.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o airfranceklm ./cmd/airfranceklm
```

## Commands

```bash
airfranceklm search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
airfranceklm read [--json] [--brand BRAND] <id|url>
airfranceklm brands
```

## Environment

- `AIRFRANCEKLM_COOKIE` — optional browser cookie when blocked
- `AIRFRANCEKLM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- Air France
- KLM
- Transavia

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
