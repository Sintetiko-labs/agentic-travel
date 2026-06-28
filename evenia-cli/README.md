# Evenia CLI

Unofficial, agent-friendly CLI for [Evenia](https://www.eveniahotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o evenia ./cmd/evenia
```

## Commands

```bash
evenia search [--json] [--limit N] <destination>
evenia read [--json] <id|url>
evenia availability [--json] --check-in DATE --check-out DATE <hotel-id>
evenia brands
```

## Environment

- `EVENIA_COOKIE` — optional browser cookie when blocked
- `EVENIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
