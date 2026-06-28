# Norse Atlantic Airways CLI

Unofficial, agent-friendly CLI for [Norse Atlantic Airways](https://www.flynorse.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o norse ./cmd/norse
```

## Commands

```bash
norse search [--json] --from MAD --to BCN --depart 2026-07-01
norse read [--json] <id|url>
norse brands
```

## Environment

- `NORSE_COOKIE` — optional browser cookie when blocked
- `NORSE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
