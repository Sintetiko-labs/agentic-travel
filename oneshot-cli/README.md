# One Shot CLI

Unofficial, agent-friendly CLI for [One Shot](https://www.oneshothotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o oneshot ./cmd/oneshot
```

## Commands

```bash
oneshot search [--json] [--limit N] <destination>
oneshot read [--json] <id|url>
oneshot availability [--json] --check-in DATE --check-out DATE <hotel-id>
oneshot brands
```

## Environment

- `ONESHOT_COOKIE` — optional browser cookie when blocked
- `ONESHOT_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
