# Iberia Express CLI

Unofficial, agent-friendly CLI for [Iberia Express](https://www.iberiaexpress.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberiaexpress ./cmd/iberiaexpress
```

## Commands

```bash
iberiaexpress search [--json] --from MAD --to BCN --depart 2026-07-01
iberiaexpress read [--json] <id|url>
iberiaexpress brands
```

## Environment

- `IBERIAEXPRESS_COOKIE` — optional browser cookie when blocked
- `IBERIAEXPRESS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
