# Vietnam Airlines CLI

Unofficial, agent-friendly CLI for [Vietnam Airlines](https://www.vietnamairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o vietnamairlines ./cmd/vietnamairlines
```

## Commands

```bash
vietnamairlines search [--json] --from MAD --to BCN --depart 2026-07-01
vietnamairlines read [--json] <id|url>
vietnamairlines brands
```

## Environment

- `VIETNAMAIRLINES_COOKIE` — optional browser cookie when blocked
- `VIETNAMAIRLINES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
