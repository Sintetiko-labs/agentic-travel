# Limehome CLI

Unofficial, agent-friendly CLI for [Limehome](https://www.limehome.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o limehome ./cmd/limehome
```

## Commands

```bash
limehome search [--json] [--limit N] <destination>
limehome read [--json] <id|url>
limehome availability [--json] --check-in DATE --check-out DATE <hotel-id>
limehome brands
```

## Environment

- `LIMEHOME_COOKIE` — optional browser cookie when blocked
- `LIMEHOME_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
