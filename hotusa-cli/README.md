# Hotusa CLI

Unofficial, agent-friendly CLI for [Hotusa](https://www.hotusa.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hotusa ./cmd/hotusa
```

## Commands

```bash
hotusa search [--json] [--limit N] [--brand BRAND] <destination>
hotusa read [--json] [--brand BRAND] <id|url>
hotusa availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
hotusa brands
```

## Environment

- `HOTUSA_COOKIE` — optional browser cookie when blocked
- `HOTUSA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Hotusa booking API:

- Hotusa
- Crisol Hotels

Use `--brand` to select a sub-brand when searching.

## Status

Category: **hotel**

Search: **partial** — custom TLS (cert `booking-channel.com`) + HTML links; may need `hotusa session chrome --wait`
