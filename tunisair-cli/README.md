# Tunisair CLI

Unofficial, agent-friendly CLI for [Tunisair](https://www.tunisair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o tunisair ./cmd/tunisair
```

## Commands

```bash
tunisair search [--json] --from MAD --to BCN --depart 2026-07-01
tunisair read [--json] <id|url>
tunisair brands
```

## Environment

- `TUNISAIR_COOKIE` — optional browser cookie when blocked
- `TUNISAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
