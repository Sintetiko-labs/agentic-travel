# Accor CLI

Unofficial, agent-friendly CLI for [Accor](https://all.accor.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o accor ./cmd/accor
```

## Commands

```bash
accor search [--json] [--limit N] [--brand BRAND] <destination>
accor read [--json] [--brand BRAND] <id|url>
accor availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
accor brands
```

## Environment

- `ACCOR_COOKIE` — optional browser cookie when blocked
- `ACCOR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Accor booking API:

- Ibis
- Ibis Budget
- Ibis Styles
- Novotel
- Mercure
- Pullman
- Sofitel
- MGallery
- Fairmont
- Raffles
- Accor

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
