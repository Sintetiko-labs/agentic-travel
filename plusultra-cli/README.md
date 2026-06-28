# Plus Ultra CLI

Unofficial, agent-friendly CLI for [Plus Ultra](https://www.plusultra.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o plusultra ./cmd/plusultra
```

## Commands

```bash
plusultra search [--json] --from MAD --to BCN --depart 2026-07-01
plusultra read [--json] <id|url>
plusultra brands
```

## Environment

- `PLUSULTRA_COOKIE` — optional browser cookie when blocked
- `PLUSULTRA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
