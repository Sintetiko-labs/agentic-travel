# Silken CLI

Unofficial, agent-friendly CLI for [Silken](https://www.hoteles-silken.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o silken ./cmd/silken
```

## Commands

```bash
silken search [--json] [--limit N] <destination>
silken read [--json] <id|url>
silken availability [--json] --check-in DATE --check-out DATE <hotel-id>
silken brands
```

## Environment

- `SILKEN_COOKIE` — optional browser cookie when blocked
- `SILKEN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **partial** — listing page data attributes (`html-cards`); no Madrid properties (try Barcelona/Bilbao)
