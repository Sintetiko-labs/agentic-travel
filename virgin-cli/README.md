# Virgin Hotels CLI

Unofficial, agent-friendly CLI for [Virgin Hotels](https://www.virginhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o virgin ./cmd/virgin
```

## Commands

```bash
virgin search [--json] [--limit N] <destination>
virgin read [--json] <id|url>
virgin availability [--json] --check-in DATE --check-out DATE <hotel-id>
virgin brands
```

## Environment

- `VIRGIN_COOKIE` — optional browser cookie when blocked
- `VIRGIN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
