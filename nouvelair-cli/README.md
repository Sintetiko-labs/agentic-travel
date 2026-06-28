# Nouvelair CLI

Unofficial, agent-friendly CLI for [Nouvelair](https://www.nouvelair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o nouvelair ./cmd/nouvelair
```

## Commands

```bash
nouvelair search [--json] --from MAD --to BCN --depart 2026-07-01
nouvelair read [--json] <id|url>
nouvelair brands
```

## Environment

- `NOUVELAIR_COOKIE` — optional browser cookie when blocked
- `NOUVELAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
