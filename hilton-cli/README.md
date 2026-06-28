# Hilton CLI

Unofficial, agent-friendly CLI for [Hilton](https://www.hilton.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hilton ./cmd/hilton
```

## Commands

```bash
hilton search [--json] [--limit N] [--brand BRAND] <destination>
hilton read [--json] [--brand BRAND] <id|url>
hilton availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
hilton brands
hilton session chrome|sync|doctor
```

## Search (London / UK)

`search` fetches Hilton destination listing pages (`/en/locations/united-kingdom/{city}/`) and parses hotel cards. Works without session for UK city pages in most residential networks.

```bash
hilton search --json London
hilton search --json London --brand Waldorf --limit 5
```

## Environment

- `HILTON_COOKIE` — optional browser cookie when blocked (GraphQL/search APIs)
- `HILTON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Session (Akamai / GraphQL)

Some endpoints (`/graphql/customer`, availability APIs) return 403 without cookies:

```bash
hilton session chrome --wait --timeout 3m
hilton session doctor
hilton search --json London
```

Chrome opens the London locations page; cookies save to `~/.hilton/cookies.json`.

## Sub-brands

- Hilton, Conrad, Waldorf Astoria, DoubleTree, Canopy, Curio Collection, Hampton

Use `--brand` to filter search results by name.

## Status

Category: **hotel** · Search: **live** (UK locations HTML) · Session: partial (Akamai on GraphQL/availability)
