# 25hours Hotels CLI

Unofficial, agent-friendly CLI for [25hours Hotels](https://www.25hours-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o 25hours ./cmd/25hours
```

## Commands

```bash
25hours search [--json] [--limit N] <destination>
25hours read [--json] <id|url>
25hours availability [--json] --check-in DATE --check-out DATE <hotel-id>
25hours brands
```

## Environment

- `25HOURS_COOKIE` — optional browser cookie when blocked
- `25HOURS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
