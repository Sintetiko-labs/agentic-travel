# Seaside Collection CLI

Unofficial, agent-friendly CLI for [Seaside Collection](https://www.seaside-collection.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o seaside ./cmd/seaside
```

## Commands

```bash
seaside search [--json] [--limit N] <destination>
seaside read [--json] <id|url>
seaside availability [--json] --check-in DATE --check-out DATE <hotel-id>
seaside brands
```

## Environment

- `SEASIDE_COOKIE` — optional browser cookie when blocked
- `SEASIDE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
