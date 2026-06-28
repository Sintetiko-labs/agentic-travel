# Rosewood CLI

Unofficial, agent-friendly CLI for [Rosewood](https://www.rosewoodhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o rosewood ./cmd/rosewood
```

## Commands

```bash
rosewood search [--json] [--limit N] <destination>
rosewood read [--json] <id|url>
rosewood availability [--json] --check-in DATE --check-out DATE <hotel-id>
rosewood brands
```

## Environment

- `ROSEWOOD_COOKIE` — optional browser cookie when blocked
- `ROSEWOOD_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
