# Servigroup CLI

Unofficial, agent-friendly CLI for [Servigroup](https://www.servigroup.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o servigroup ./cmd/servigroup
```

## Commands

```bash
servigroup search [--json] [--limit N] <destination>
servigroup read [--json] <id|url>
servigroup availability [--json] --check-in DATE --check-out DATE <hotel-id>
servigroup brands
```

## Environment

- `SERVIGROUP_COOKIE` — optional browser cookie when blocked
- `SERVIGROUP_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
