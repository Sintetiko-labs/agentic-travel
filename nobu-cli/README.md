# Nobu Hotels CLI

Unofficial, agent-friendly CLI for [Nobu Hotels](https://www.nobuhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o nobu ./cmd/nobu
```

## Commands

```bash
nobu search [--json] [--limit N] <destination>
nobu read [--json] <id|url>
nobu availability [--json] --check-in DATE --check-out DATE <hotel-id>
nobu brands
```

## Environment

- `NOBU_COOKIE` — optional browser cookie when blocked
- `NOBU_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
