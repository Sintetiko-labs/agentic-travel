# Qantas CLI

Unofficial, agent-friendly CLI for [Qantas](https://www.qantas.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o qantas ./cmd/qantas
```

## Commands

```bash
qantas search [--json] --from MAD --to BCN --depart 2026-07-01
qantas read [--json] <id|url>
qantas brands
```

## Environment

- `QANTAS_COOKIE` — optional browser cookie when blocked
- `QANTAS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
