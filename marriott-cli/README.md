# Marriott CLI

Unofficial, agent-friendly CLI for [Marriott](https://www.marriott.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o marriott ./cmd/marriott
```

## Commands

```bash
marriott search [--json] [--limit N] [--brand BRAND] <destination>
marriott read [--json] [--brand BRAND] <id|url>
marriott availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
marriott brands
marriott session chrome|sync|doctor
```

## Search (London / UK)

`search` calls `findHotels.mi` for the destination city. **Akamai blocks unauthenticated requests** — capture a headed Chrome session first.

```bash
marriott session chrome --wait --timeout 3m   # browse London search if challenged
marriott session doctor
marriott search --json London
```

Without session, `search` returns `akamai blocked — HILTON_COOKIE required` style error with `marriott session chrome` hint.

## Environment

- `MARRIOTT_COOKIE` — browser cookie from headed Chrome (required for live search)
- `MARRIOTT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Session (required for search)

Marriott returns HTTP 403 (Akamai) without `_abck` + `bm_sz` cookies:

```bash
marriott session chrome --wait --timeout 3m
marriott session sync          # copy cookies from running Chrome
marriott session doctor
```

Chrome starts at a London `findHotels.mi` URL; cookies save to `~/.marriott/cookies.json`.

## Sub-brands

Marriott, JW Marriott, Ritz-Carlton, St. Regis, W Hotels, Westin, Sheraton, Courtyard, and more — use `--brand` to filter.

## Status

Category: **hotel** · Search: **partial** (findHotels + session chrome) · Session: **required**
