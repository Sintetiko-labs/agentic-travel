# Hospes CLI

Unofficial, agent-friendly CLI for [Hospes](https://www.hospes.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hospes ./cmd/hospes
```

## Commands

```bash
hospes search [--json] [--limit N] <destination>
hospes read [--json] <id|url>
hospes availability [--json] --check-in DATE --check-out DATE <hotel-id>
hospes brands
```

## Environment

- `HOSPES_COOKIE` — optional browser cookie when blocked
- `HOSPES_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
