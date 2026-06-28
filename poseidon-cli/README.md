# Hoteles Poseidón CLI

Unofficial, agent-friendly CLI for [Hoteles Poseidón](https://www.hoteles-poseidon.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o poseidon ./cmd/poseidon
```

## Commands

```bash
poseidon search [--json] [--limit N] <destination>
poseidon read [--json] <id|url>
poseidon availability [--json] --check-in DATE --check-out DATE <hotel-id>
poseidon brands
```

## Environment

- `POSEIDON_COOKIE` — optional browser cookie when blocked
- `POSEIDON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
