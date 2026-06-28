# Soho Boutique CLI

Unofficial, agent-friendly CLI for [Soho Boutique](https://www.sohohoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o soho ./cmd/soho
```

## Commands

```bash
soho search [--json] [--limit N] <destination>
soho read [--json] <id|url>
soho availability [--json] --check-in DATE --check-out DATE <hotel-id>
soho brands
```

## Environment

- `SOHO_COOKIE` — optional browser cookie when blocked
- `SOHO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
