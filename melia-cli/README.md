# Meliá CLI

Unofficial, agent-friendly CLI for [Meliá](https://www.melia.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o melia ./cmd/melia
```

## Commands

```bash
melia search [--json] [--limit N] [--brand BRAND] <destination>
melia read [--json] [--brand BRAND] <id|url>
melia availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
melia brands
```

## Environment

- `MELIA_COOKIE` — optional browser cookie when blocked
- `MELIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Meliá booking API:

- Meliá Hotels International
- Meliá
- Gran Meliá
- ME by Meliá
- The Meliá Collection
- Paradisus
- INNSiDE by Meliá
- Sol by Meliá
- ZEL

Use `--brand` to select a sub-brand when searching.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — BFF `/services/search/hotels/v2/search`; needs `MELIA_COOKIE` (Akamai) |
| `read` / `availability` | implemented (cookie for live) |
| Rate limit | `MELIA_REQUEST_DELAY` (~2s) |
