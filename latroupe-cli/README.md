# Latroupe CLI

Unofficial, agent-friendly CLI for [Latroupe](https://www.latroupe.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o latroupe ./cmd/latroupe
```

## Commands

```bash
latroupe search [--json] [--limit N] <destination>
latroupe read [--json] <id|url>
latroupe availability [--json] --check-in DATE --check-out DATE <hotel-id>
latroupe brands
```

## Environment

- `LATROUPE_COOKIE` — optional browser cookie when blocked
- `LATROUPE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
