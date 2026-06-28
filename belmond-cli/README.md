# Belmond CLI

Unofficial, agent-friendly CLI for [Belmond](https://www.belmond.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o belmond ./cmd/belmond
```

## Commands

```bash
belmond search [--json] [--limit N] <destination>
belmond read [--json] <id|url>
belmond availability [--json] --check-in DATE --check-out DATE <hotel-id>
belmond brands
```

## Environment

- `BELMOND_COOKIE` — optional browser cookie when blocked
- `BELMOND_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
