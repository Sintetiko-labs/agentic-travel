# Czech Airlines CLI

Unofficial, agent-friendly CLI for [Czech Airlines](https://www.csa.cz).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o czechairlines ./cmd/czechairlines
```

## Commands

```bash
czechairlines search [--json] --from MAD --to BCN --depart 2026-07-01
czechairlines read [--json] <id|url>
czechairlines brands
```

## Environment

- `CZECHAIRLINES_COOKIE` — optional browser cookie when blocked
- `CZECHAIRLINES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
