# IHG CLI

Unofficial, agent-friendly CLI for [IHG](https://www.ihg.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ihg ./cmd/ihg
```

## Commands

```bash
ihg search [--json] [--limit N] [--brand BRAND] <destination>
ihg read [--json] [--brand BRAND] <id|url>
ihg availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
ihg brands
```

## Environment

- `IHG_COOKIE` — optional browser cookie when blocked
- `IHG_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the IHG booking API:

- IHG Hotels & Resorts
- InterContinental
- Kimpton
- Crowne Plaza
- Holiday Inn
- Holiday Inn Express
- Hotel Indigo
- Six Senses
- Vignette Collection

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
