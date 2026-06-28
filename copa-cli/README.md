# Copa Airlines CLI

Unofficial, agent-friendly CLI for [Copa Airlines](https://www.copaair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o copa ./cmd/copa
```

## Commands

```bash
copa search [--json] --from MAD --to BCN --depart 2026-07-01
copa read [--json] <id|url>
copa brands
```

## Environment

- `COPA_COOKIE` — optional browser cookie when blocked
- `COPA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
