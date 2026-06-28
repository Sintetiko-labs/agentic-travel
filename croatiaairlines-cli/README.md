# Croatia Airlines CLI

Unofficial, agent-friendly CLI for [Croatia Airlines](https://www.croatiaairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o croatiaairlines ./cmd/croatiaairlines
```

## Commands

```bash
croatiaairlines search [--json] --from MAD --to BCN --depart 2026-07-01
croatiaairlines read [--json] <id|url>
croatiaairlines brands
```

## Environment

- `CROATIAAIRLINES_COOKIE` — optional browser cookie when blocked
- `CROATIAAIRLINES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
