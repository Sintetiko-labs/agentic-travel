# CoolRooms CLI

Unofficial, agent-friendly CLI for [CoolRooms](https://www.coolrooms.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o coolrooms ./cmd/coolrooms
```

## Commands

```bash
coolrooms search [--json] [--limit N] <destination>
coolrooms read [--json] <id|url>
coolrooms availability [--json] --check-in DATE --check-out DATE <hotel-id>
coolrooms brands
```

## Environment

- `COOLROOMS_COOKIE` — optional browser cookie when blocked
- `COOLROOMS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
