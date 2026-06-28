# Royal Air Maroc CLI

Unofficial, agent-friendly CLI for [Royal Air Maroc](https://www.royalairmaroc.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o royalairmaroc ./cmd/royalairmaroc
```

## Commands

```bash
royalairmaroc search [--json] --from MAD --to BCN --depart 2026-07-01
royalairmaroc read [--json] <id|url>
royalairmaroc brands
```

## Environment

- `ROYALAIRMAROC_COOKIE` — optional browser cookie when blocked
- `ROYALAIRMAROC_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
