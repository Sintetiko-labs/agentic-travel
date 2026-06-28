# Leonardo Hotels CLI

Unofficial, agent-friendly CLI for [Leonardo Hotels](https://www.leonardo-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o leonardo ./cmd/leonardo
```

## Commands

```bash
leonardo search [--json] [--limit N] [--brand BRAND] <destination>
leonardo read [--json] [--brand BRAND] <id|url>
leonardo availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
leonardo brands
```

## Environment

- `LEONARDO_COOKIE` — optional browser cookie when blocked
- `LEONARDO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Leonardo Hotels booking API:

- Leonardo Hotels
- NYX Hotels
- Leonardo Royal
- Leonardo Boutique

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
