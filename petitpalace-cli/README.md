# Petit Palace CLI

Unofficial, agent-friendly CLI for [Petit Palace](https://www.petitpalace.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o petitpalace ./cmd/petitpalace
```

## Commands

```bash
petitpalace search [--json] [--limit N] <destination>
petitpalace read [--json] <id|url>
petitpalace availability [--json] --check-in DATE --check-out DATE <hotel-id>
petitpalace brands
```

## Environment

- `PETITPALACE_COOKIE` — optional browser cookie when blocked
- `PETITPALACE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
