# Design Hotels CLI

Unofficial, agent-friendly CLI for [Design Hotels](https://www.designhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o designhotels ./cmd/designhotels
```

## Commands

```bash
designhotels search [--json] [--limit N] <destination>
designhotels read [--json] <id|url>
designhotels availability [--json] --check-in DATE --check-out DATE <hotel-id>
designhotels brands
```

## Environment

- `DESIGNHOTELS_COOKIE` — optional browser cookie when blocked
- `DESIGNHOTELS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
