# Hainan Airlines CLI

Unofficial, agent-friendly CLI for [Hainan Airlines](https://www.hainanairlines.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hainan ./cmd/hainan
```

## Commands

```bash
hainan search [--json] --from MAD --to BCN --depart 2026-07-01
hainan read [--json] <id|url>
hainan brands
```

## Environment

- `HAINAN_COOKIE` — optional browser cookie when blocked
- `HAINAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
