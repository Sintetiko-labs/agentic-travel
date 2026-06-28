# Casual Hoteles CLI

Unofficial, agent-friendly CLI for [Casual Hoteles](https://www.casualhoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o casual ./cmd/casual
```

## Commands

```bash
casual search [--json] [--limit N] <destination>
casual read [--json] <id|url>
casual availability [--json] --check-in DATE --check-out DATE <hotel-id>
casual brands
```

## Environment

- `CASUAL_COOKIE` — optional browser cookie when blocked
- `CASUAL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
