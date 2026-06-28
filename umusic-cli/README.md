# UMusic Hotels CLI

Unofficial, agent-friendly CLI for [UMusic Hotels](https://www.umusichotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o umusic ./cmd/umusic
```

## Commands

```bash
umusic search [--json] [--limit N] <destination>
umusic read [--json] <id|url>
umusic availability [--json] --check-in DATE --check-out DATE <hotel-id>
umusic brands
```

## Environment

- `UMUSIC_COOKIE` — optional browser cookie when blocked
- `UMUSIC_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
