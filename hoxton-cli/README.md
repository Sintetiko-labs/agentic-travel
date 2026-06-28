# The Hoxton CLI

Unofficial, agent-friendly CLI for [The Hoxton](https://www.thehoxton.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hoxton ./cmd/hoxton
```

## Commands

```bash
hoxton search [--json] [--limit N] <destination>
hoxton read [--json] <id|url>
hoxton availability [--json] --check-in DATE --check-out DATE <hotel-id>
hoxton brands
```

## Environment

- `HOXTON_COOKIE` — optional browser cookie when blocked
- `HOXTON_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
