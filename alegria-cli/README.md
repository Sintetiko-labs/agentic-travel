# Alegria CLI

Unofficial, agent-friendly CLI for [Alegria](https://www.alegriahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o alegria ./cmd/alegria
```

## Commands

```bash
alegria search [--json] [--limit N] <destination>
alegria read [--json] <id|url>
alegria availability [--json] --check-in DATE --check-out DATE <hotel-id>
alegria brands
```

## Environment

- `ALEGRIA_COOKIE` — optional browser cookie when blocked
- `ALEGRIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
