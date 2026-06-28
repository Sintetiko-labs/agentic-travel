# Iberostar CLI

Unofficial, agent-friendly CLI for [Iberostar](https://www.iberostar.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberostar ./cmd/iberostar
```

## Commands

```bash
iberostar search [--json] [--limit N] [--brand BRAND] <destination>
iberostar read [--json] [--brand BRAND] <id|url>
iberostar availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
iberostar brands
```

## Environment

- `IBEROSTAR_COOKIE` — optional browser cookie when blocked
- `IBEROSTAR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Iberostar booking API:

- Iberostar
- Iberostar Selection
- Iberostar Grand

Use `--brand` to select a sub-brand when searching.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — GraphQL `/api/graphql`; needs `IBEROSTAR_COOKIE` |
| `read` / `availability` | implemented |
| Rate limit | `IBEROSTAR_REQUEST_DELAY` (~2s) |
