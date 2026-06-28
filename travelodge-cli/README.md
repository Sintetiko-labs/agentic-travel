# Travelodge CLI

Unofficial, agent-friendly CLI for [Travelodge](https://www.travelodge.co.uk).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o travelodge ./cmd/travelodge
```

## Commands

```bash
travelodge search [--json] [--limit N] <destination>
travelodge read [--json] <id|url>
travelodge availability [--json] --check-in DATE --check-out DATE <hotel-id>
travelodge brands
travelodge session chrome|sync|doctor
```

## Search (London / UK)

`search` reads `sitemap-fusion.xml` and filters hotel URLs by destination (e.g. `London` matches `london` and `greater-london` paths). No session required in most cases.

```bash
travelodge search --json London
travelodge search --json "greater london" --limit 5
```

## Environment

- `TRAVELODGE_COOKIE` — optional browser cookie when sitemap is WAF-blocked
- `TRAVELODGE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Session (WAF fallback)

If the sitemap returns 403:

```bash
travelodge session chrome --wait --timeout 3m
travelodge session doctor
travelodge search --json London
```

Chrome opens `travelodge.co.uk/uk/london/`; cookies save to `~/.travelodge/cookies.json`.

## Status

Category: **hotel** · Search: **live** (sitemap) · Session: optional (WAF fallback)
