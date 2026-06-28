# Mama Shelter CLI

Unofficial, agent-friendly CLI for [Mama Shelter](https://www.mamashelter.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o mamashelter ./cmd/mamashelter
```

## Commands

```bash
mamashelter search [--json] [--limit N] <destination>
mamashelter read [--json] <id|url>
mamashelter availability [--json] --check-in DATE --check-out DATE <hotel-id>
mamashelter brands
```

## Environment

- `MAMASHELTER_COOKIE` — optional browser cookie when blocked
- `MAMASHELTER_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
