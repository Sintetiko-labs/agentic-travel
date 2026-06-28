# Barceló CLI

Unofficial, agent-friendly CLI for [Barceló](https://www.barcelo.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o barcelo ./cmd/barcelo
```

## Commands

```bash
barcelo search [--json] [--limit N] [--brand BRAND] <destination>
barcelo read [--json] [--brand BRAND] <id|url>
barcelo availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
barcelo brands
```

## Environment

- `BARCELO_COOKIE` — optional browser cookie when blocked
- `BARCELO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Barceló booking API:

- Barceló Hotel Group
- Barceló Hotels & Resorts
- Royal Hideaway
- Occidental Hotels & Resorts
- Allegro Hotels

Use `--brand` to select a sub-brand when searching.

## Status

| Feature | Status |
|---------|--------|
| `search` | **live** — JSON-LD from `/es/hoteles` |
| `read` / `availability` | implemented |
| Rate limit | `BARCELO_REQUEST_DELAY` (~2s) |

```bash
barcelo search --json Madrid
```
