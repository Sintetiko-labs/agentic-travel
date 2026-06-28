# BLESS Collection CLI

Unofficial, agent-friendly CLI for [BLESS Collection](https://www.blesscollectionhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o bless ./cmd/bless
```

## Commands

```bash
bless search [--json] [--limit N] <destination>
bless read [--json] <id|url>
bless availability [--json] --check-in DATE --check-out DATE <hotel-id>
bless brands
```

## Environment

- `BLESS_COOKIE` — optional browser cookie when blocked
- `BLESS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
