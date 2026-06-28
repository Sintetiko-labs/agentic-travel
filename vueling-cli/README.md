# Vueling CLI

Unofficial, agent-friendly CLI for [Vueling](https://www.vueling.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o vueling ./cmd/vueling
```

## Commands

```bash
vueling search [--json] --from MAD --to BCN --depart 2026-07-01
vueling read [--json] <id|url>
vueling brands
```

## Environment

- `VUELING_COOKIE` — optional browser cookie when blocked
- `VUELING_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `/bit/v2/flights/search`; needs `VUELING_COOKIE` |
| `read` | implemented |
| Rate limit | `VUELING_REQUEST_DELAY` (~2s) |
