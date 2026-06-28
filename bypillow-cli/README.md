# ByPillow CLI

Unofficial, agent-friendly CLI for [ByPillow](https://www.bypillow.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o bypillow ./cmd/bypillow
```

## Commands

```bash
bypillow search [--json] [--limit N] <destination>
bypillow read [--json] <id|url>
bypillow availability [--json] --check-in DATE --check-out DATE <hotel-id>
bypillow brands
```

## Environment

- `BYPILLOW_COOKIE` — optional browser cookie when blocked
- `BYPILLOW_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
