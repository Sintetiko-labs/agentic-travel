# Volotea CLI

Unofficial, agent-friendly CLI for [Volotea](https://www.volotea.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o volotea ./cmd/volotea
```

## Commands

```bash
volotea search [--json] --from MAD --to BCN --depart 2026-07-01
volotea read [--json] <id|url>
volotea brands
```

## Environment

- `VOLOTEA_COOKIE` — optional browser cookie when blocked
- `VOLOTEA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **live** — Volotea booking SPA API (`api.volotea.com/api/spa/voe/v1/flights/search`, anonymous login)
