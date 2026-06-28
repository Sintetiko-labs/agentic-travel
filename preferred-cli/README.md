# Preferred Hotels CLI

Unofficial, agent-friendly CLI for [Preferred Hotels](https://www.preferredhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o preferred ./cmd/preferred
```

## Commands

```bash
preferred search [--json] [--limit N] <destination>
preferred read [--json] <id|url>
preferred availability [--json] --check-in DATE --check-out DATE <hotel-id>
preferred brands
```

## Environment

- `PREFERRED_COOKIE` — optional browser cookie when blocked
- `PREFERRED_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
