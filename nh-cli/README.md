# NH Hotel Group CLI

Unofficial, agent-friendly CLI for [NH Hotel Group](https://www.nh-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o nh ./cmd/nh
```

## Commands

```bash
nh search [--json] [--limit N] [--brand BRAND] <destination>
nh read [--json] [--brand BRAND] <id|url>
nh availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
nh brands
```

## Environment

- `NH_COOKIE` — optional browser cookie when blocked
- `NH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the NH Hotel Group booking API:

- NH Hotel Group
- NH Hotels
- NH Collection
- nhow

Use `--brand` to select a sub-brand when searching.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `/nh/es/api/v1/hotels/search`; needs `NH_COOKIE` |
| `read` / `availability` | implemented |
| Rate limit | `NH_REQUEST_DELAY` (~2s) |
