# HTop CLI

Unofficial, agent-friendly CLI for [HTop](https://www.htophotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o htop ./cmd/htop
```

## Commands

```bash
htop search [--json] [--limit N] <destination>
htop read [--json] <id|url>
htop availability [--json] --check-in DATE --check-out DATE <hotel-id>
htop brands
```

## Environment

- `HTOP_COOKIE` — optional browser cookie when blocked
- `HTOP_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
