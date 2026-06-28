# Locke CLI

Unofficial, agent-friendly CLI for [Locke](https://www.lockeliving.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o locke ./cmd/locke
```

## Commands

```bash
locke search [--json] [--limit N] <destination>
locke read [--json] <id|url>
locke availability [--json] --check-in DATE --check-out DATE <hotel-id>
locke brands
```

## Environment

- `LOCKE_COOKIE` — optional browser cookie when blocked
- `LOCKE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
