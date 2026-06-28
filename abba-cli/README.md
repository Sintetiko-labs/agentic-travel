# Abba Hoteles CLI

Unofficial, agent-friendly CLI for [Abba Hoteles](https://www.abbahoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o abba ./cmd/abba
```

## Commands

```bash
abba search [--json] [--limit N] <destination>
abba read [--json] <id|url>
abba availability [--json] --check-in DATE --check-out DATE <hotel-id>
abba brands
```

## Environment

- `ABBA_COOKIE` — optional browser cookie when blocked
- `ABBA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
