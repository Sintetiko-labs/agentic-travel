# PLAY Airlines CLI

Unofficial, agent-friendly CLI for [PLAY Airlines](https://www.flyplay.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o play ./cmd/play
```

## Commands

```bash
play search [--json] --from MAD --to BCN --depart 2026-07-01
play read [--json] <id|url>
play brands
```

## Environment

- `PLAY_COOKIE` — optional browser cookie when blocked
- `PLAY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
