# Korean Air CLI

Unofficial, agent-friendly CLI for [Korean Air](https://www.koreanair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o koreanair ./cmd/koreanair
```

## Commands

```bash
koreanair search [--json] --from MAD --to BCN --depart 2026-07-01
koreanair read [--json] <id|url>
koreanair brands
```

## Environment

- `KOREANAIR_COOKIE` — optional browser cookie when blocked
- `KOREANAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
