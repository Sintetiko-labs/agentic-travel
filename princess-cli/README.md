# Princess Hotels CLI

Unofficial, agent-friendly CLI for [Princess Hotels](https://www.princess-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o princess ./cmd/princess
```

## Commands

```bash
princess search [--json] [--limit N] <destination>
princess read [--json] <id|url>
princess availability [--json] --check-in DATE --check-out DATE <hotel-id>
princess brands
```

## Environment

- `PRINCESS_COOKIE` — optional browser cookie when blocked
- `PRINCESS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
