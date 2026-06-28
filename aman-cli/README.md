# Aman CLI

Unofficial, agent-friendly CLI for [Aman](https://www.aman.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aman ./cmd/aman
```

## Commands

```bash
aman search [--json] [--limit N] <destination>
aman read [--json] <id|url>
aman availability [--json] --check-in DATE --check-out DATE <hotel-id>
aman brands
```

## Environment

- `AMAN_COOKIE` — optional browser cookie when blocked
- `AMAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
