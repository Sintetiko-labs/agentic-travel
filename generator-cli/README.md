# Generator Hostels CLI

Unofficial, agent-friendly CLI for [Generator Hostels](https://www.staygenerator.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o generator ./cmd/generator
```

## Commands

```bash
generator search [--json] [--limit N] <destination>
generator read [--json] <id|url>
generator availability [--json] --check-in DATE --check-out DATE <hotel-id>
generator brands
```

## Environment

- `GENERATOR_COOKIE` — optional browser cookie when blocked
- `GENERATOR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
