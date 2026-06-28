# Sonder CLI

Unofficial, agent-friendly CLI for [Sonder](https://www.sonder.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o sonder ./cmd/sonder
```

## Commands

```bash
sonder search [--json] [--limit N] <destination>
sonder read [--json] <id|url>
sonder availability [--json] --check-in DATE --check-out DATE <hotel-id>
sonder brands
```

## Environment

- `SONDER_COOKIE` — optional browser cookie when blocked
- `SONDER_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
