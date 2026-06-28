# Aegean CLI

Unofficial, agent-friendly CLI for [Aegean](https://www.aegeanair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aegean ./cmd/aegean
```

## Commands

```bash
aegean search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
aegean read [--json] [--brand BRAND] <id|url>
aegean brands
```

## Environment

- `AEGEAN_COOKIE` — optional browser cookie when blocked
- `AEGEAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- Aegean Airlines
- Olympic Air

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
