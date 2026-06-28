# Best Western CLI

Unofficial, agent-friendly CLI for [Best Western](https://www.bestwestern.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o bestwestern ./cmd/bestwestern
```

## Commands

```bash
bestwestern search [--json] [--limit N] [--brand BRAND] <destination>
bestwestern read [--json] [--brand BRAND] <id|url>
bestwestern availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
bestwestern brands
```

## Environment

- `BESTWESTERN_COOKIE` — optional browser cookie when blocked
- `BESTWESTERN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Best Western booking API:

- Best Western
- Best Western Plus
- Best Western Premier
- BWH Hotel Group

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
