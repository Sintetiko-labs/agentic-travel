# Asiana Airlines CLI

Unofficial, agent-friendly CLI for [Asiana Airlines](https://www.flyasiana.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o asiana ./cmd/asiana
```

## Commands

```bash
asiana search [--json] --from MAD --to BCN --depart 2026-07-01
asiana read [--json] <id|url>
asiana brands
```

## Environment

- `ASIANA_COOKIE` — optional browser cookie when blocked
- `ASIANA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
