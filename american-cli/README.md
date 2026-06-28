# American Airlines CLI

Unofficial, agent-friendly CLI for [American Airlines](https://www.aa.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o american ./cmd/american
```

## Commands

```bash
american search [--json] --from MAD --to BCN --depart 2026-07-01
american read [--json] <id|url>
american brands
```

## Environment

- `AMERICAN_COOKIE` — optional browser cookie when blocked
- `AMERICAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
