# Líbere CLI

Unofficial, agent-friendly CLI for [Líbere](https://www.liberehospitality.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o libere ./cmd/libere
```

## Commands

```bash
libere search [--json] [--limit N] <destination>
libere read [--json] <id|url>
libere availability [--json] --check-in DATE --check-out DATE <hotel-id>
libere brands
```

## Environment

- `LIBERE_COOKIE` — optional browser cookie when blocked
- `LIBERE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
