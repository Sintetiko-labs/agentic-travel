# Único Hotels CLI

Unofficial, agent-friendly CLI for [Único Hotels](https://www.unicohotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o unico ./cmd/unico
```

## Commands

```bash
unico search [--json] [--limit N] <destination>
unico read [--json] <id|url>
unico availability [--json] --check-in DATE --check-out DATE <hotel-id>
unico brands
```

## Environment

- `UNICO_COOKIE` — optional browser cookie when blocked
- `UNICO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
