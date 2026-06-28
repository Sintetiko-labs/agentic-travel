# Room Mate CLI

Unofficial, agent-friendly CLI for [Room Mate](https://www.room-matehotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o roommate ./cmd/roommate
```

## Commands

```bash
roommate search [--json] [--limit N] <destination>
roommate read [--json] <id|url>
roommate availability [--json] --check-in DATE --check-out DATE <hotel-id>
roommate brands
```

## Environment

- `ROOMMATE_COOKIE` — optional browser cookie when blocked
- `ROOMMATE_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
