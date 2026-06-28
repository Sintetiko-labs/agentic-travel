# Etihad Airways CLI

Unofficial, agent-friendly CLI for [Etihad Airways](https://www.etihad.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o etihad ./cmd/etihad
```

## Commands

```bash
etihad search [--json] --from MAD --to BCN --depart 2026-07-01
etihad read [--json] <id|url>
etihad brands
```

## Environment

- `ETIHAD_COOKIE` — optional browser cookie when blocked
- `ETIHAD_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
