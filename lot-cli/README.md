# LOT Polish Airlines CLI

Unofficial, agent-friendly CLI for [LOT Polish Airlines](https://www.lot.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o lot ./cmd/lot
```

## Commands

```bash
lot search [--json] --from MAD --to BCN --depart 2026-07-01
lot read [--json] <id|url>
lot brands
```

## Environment

- `LOT_COOKIE` — optional browser cookie when blocked
- `LOT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
