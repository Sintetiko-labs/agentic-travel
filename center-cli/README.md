# Hoteles Center CLI

Unofficial, agent-friendly CLI for [Hoteles Center](https://www.hotelescenter.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o center ./cmd/center
```

## Commands

```bash
center search [--json] [--limit N] <destination>
center read [--json] <id|url>
center availability [--json] --check-in DATE --check-out DATE <hotel-id>
center brands
```

## Environment

- `CENTER_COOKIE` — optional browser cookie when blocked
- `CENTER_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
