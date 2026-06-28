# Lopesan CLI

Unofficial, agent-friendly CLI for [Lopesan](https://www.lopesan.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o lopesan ./cmd/lopesan
```

## Commands

```bash
lopesan search [--json] [--limit N] [--brand BRAND] <destination>
lopesan read [--json] [--brand BRAND] <id|url>
lopesan availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
lopesan brands
```

## Environment

- `LOPESAN_COOKIE` — optional browser cookie when blocked
- `LOPESAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Lopesan booking API:

- Lopesan Hotel Group
- Abora by Lopesan
- Lopesan Hotels
- Lopesan Collection

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
