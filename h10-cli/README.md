# H10 CLI

Unofficial, agent-friendly CLI for [H10](https://www.h10hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o h10 ./cmd/h10
```

## Commands

```bash
h10 search [--json] [--limit N] [--brand BRAND] <destination>
h10 read [--json] [--brand BRAND] <id|url>
h10 availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
h10 brands
```

## Environment

- `H10_COOKIE` — optional browser cookie when blocked
- `H10_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the H10 booking API:

- H10 Hotels
- H10
- Ocean by H10

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
