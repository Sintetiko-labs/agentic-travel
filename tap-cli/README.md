# TAP Air Portugal CLI

Unofficial, agent-friendly CLI for [TAP Air Portugal](https://www.flytap.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o tap ./cmd/tap
```

## Commands

```bash
tap search [--json] --from MAD --to BCN --depart 2026-07-01
tap read [--json] <id|url>
tap brands
```

## Environment

- `TAP_COOKIE` — optional browser cookie when blocked
- `TAP_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
