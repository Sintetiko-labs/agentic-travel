# Singapore Airlines CLI

Unofficial, agent-friendly CLI for [Singapore Airlines](https://www.singaporeair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o singaporeairlines ./cmd/singaporeairlines
```

## Commands

```bash
singaporeairlines search [--json] --from MAD --to BCN --depart 2026-07-01
singaporeairlines read [--json] <id|url>
singaporeairlines brands
```

## Environment

- `SINGAPOREAIRLINES_COOKIE` — optional browser cookie when blocked
- `SINGAPOREAIRLINES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
