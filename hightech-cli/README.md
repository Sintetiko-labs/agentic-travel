# High Tech Hotels CLI

Unofficial, agent-friendly CLI for [High Tech Hotels](https://www.hthoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o hightech ./cmd/hightech
```

## Commands

```bash
hightech search [--json] [--limit N] <destination>
hightech read [--json] <id|url>
hightech availability [--json] --check-in DATE --check-out DATE <hotel-id>
hightech brands
```

## Environment

- `HIGHTECH_COOKIE` — optional browser cookie when blocked
- `HIGHTECH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
